package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	//log "github.com/Sirupsen/logrus"
	"github.com/xtracdev/automated-perf-test/UI"
	"github.com/xtracdev/automated-perf-test/perfTestUtils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var configurationSettings *perfTestUtils.Config
var checkTestReadyness bool
var globals map[string]string

const (
	TRAINING_MODE = 1
	TESTING_MODE  = 2
)

func init() {

	//Command line ags
	var gbs bool
	var reBaseMemory bool
	var reBaseAll bool
	var configFilePath string

	//Process command line arugments.
	flag.BoolVar(&gbs, "gbs", false, "Genertate Base Performance Staticists for this server")
	flag.BoolVar(&reBaseMemory, "reBaseMemory", false, "Generate new base peak memory for this server")
	flag.BoolVar(&reBaseAll, "reBaseAll", false, "Generate new base for memory and service resposne times for this server")
	flag.BoolVar(&checkTestReadyness, "checkTestReadyness", false, "Simple check to see if system requires training.")
	flag.Parse()

	//Read and paser config file if present.
	configurationSettings = new(perfTestUtils.Config)
	if configFilePath != "" {
		fileContent, fileErr := ioutil.ReadFile(configFilePath)
		if fileErr != nil {
			//log.Info("No config file found. ")
			fmt.Println("No config file found at path: ", configFilePath)
			os.Exit(1)
		} else {
			xmlError := xml.Unmarshal(fileContent, &configurationSettings)
			if xmlError != nil {
				//log.Info("Failed to parse config file ", configFilePath, ". Error:", xmlError)
				fmt.Println("Failed to parse config file ", configFilePath, ". Error:", xmlError)
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("No config file specified. Using default values.")
		configurationSettings.SetDefaults()
	}

	//Get Hostname for this machine.
	host, err := os.Hostname()
	if err != nil {
		//log.Error("Failed to resolve host name. Error:", err)
		fmt.Println("Failed to resolve host name. Error:", err)
		os.Exit(1)
	}
	configurationSettings.ExecutionHost = host
	configurationSettings.GBS = gbs
	configurationSettings.ReBaseMemory = reBaseMemory
	configurationSettings.ReBaseAll = reBaseAll

	//Initilize globals map
	globals = make(map[string]string)

}

//Main Test Method
func main() {

	if checkTestReadyness {
		readyForTest, _ := isReadyForTest(configurationSettings.ExecutionHost)
		if !readyForTest {
			fmt.Println("System is not ready for testing.")
			os.Exit(1)
		} else {
			fmt.Println("System is ready for testing.")
			os.Exit(0)
		}
	}

	//Validate config()
	configurationSettings.PrintAndValidateConfig()

	//Determine testing mode.
	if configurationSettings.GBS || configurationSettings.ReBaseAll {
		if configurationSettings.ReBaseAll {
			runInTrainingMode(configurationSettings.ExecutionHost, true)
		} else {
			readyForTest, _ := isReadyForTest(configurationSettings.ExecutionHost)
			if !readyForTest {
				runInTrainingMode(configurationSettings.ExecutionHost, false)
			} else {
				fmt.Println("System is ready for testing. Training is not required....")
			}
		}
	} else {
		readyForTest, basePerfStats := isReadyForTest(configurationSettings.ExecutionHost)
		if readyForTest {
			runInTestingMode(basePerfStats, configurationSettings.ExecutionHost, perfTestUtils.GenerateTemplateReport)
		} else {
			fmt.Println("System is not ready for testing. Attempting to run training mode....")
			runInTrainingMode(configurationSettings.ExecutionHost, false)
			readyForTest, basePerfStats = isReadyForTest(configurationSettings.ExecutionHost)
			if readyForTest {
				runInTestingMode(basePerfStats, configurationSettings.ExecutionHost, perfTestUtils.GenerateTemplateReport)
			} else {
				fmt.Println("System is not ready for testing. Check logs for more details.")
				os.Exit(1)
			}
		}
	}
}

func runInTrainingMode(host string, reBaseAll bool) {
	fmt.Println("Running Perf test in Training mode for host ", host)
	testStratTime := time.Now().UnixNano()

	var basePerfstats *perfTestUtils.BasePerfStats
	if reBaseAll {
		fmt.Println("Performing full rebase of perf stats for host ", host)

		basePerfstats = &perfTestUtils.BasePerfStats{
			BaseServiceResponseTimes: make(map[string]int64),
			MemoryAudit:              make([]uint64, 0),
		}
	} else {
		//Check to see if this server already has a base perf file defined.
		//If so, only values not previously populated will be set.
		//if not, a default base perf struct is created with nil values for all fields
		basePerfstats, _ = perfTestUtils.ReadBasePerfFile(host, configurationSettings.BaseStatsOutputDir)
	}

	//initilize Performance statistics struct for this test run
	perfStatsForTest := &perfTestUtils.PerfStats{ServiceResponseTimes: make(map[string]int64)}

	//Run the test
	runTests(perfStatsForTest, TRAINING_MODE)
	perfTestUtils.GenerateEnvBasePerfOutputFile(perfStatsForTest, basePerfstats, configurationSettings)

	testRunTime := time.Now().UnixNano() - testStratTime
	fmt.Println("Training mode completed successfully. ", getExecutionTimeDisplay(testRunTime))
}

func runInTestingMode(basePerfstats *perfTestUtils.BasePerfStats, host string, frg func(*perfTestUtils.BasePerfStats, *perfTestUtils.PerfStats, *perfTestUtils.Config)) {
	fmt.Println("Running Perf test in Testing mode for host ", host)
	testStratTime := time.Now().UnixNano()

	//initilize Performance statistics struct for this test run
	perfStatsForTest := &perfTestUtils.PerfStats{ServiceResponseTimes: make(map[string]int64)}
	perfStatsForTest.TestDate = time.Now()

	runTests(perfStatsForTest, TESTING_MODE)
	assertionFailures := runAssertions(basePerfstats, perfStatsForTest)
	frg(basePerfstats, perfStatsForTest, configurationSettings)

	fmt.Println("=================== TEST RESULTS ===================")
	if len(assertionFailures) > 0 {
		fmt.Println("Failures : ", len(assertionFailures))
		//Print assertion failures
		for _, failure := range assertionFailures {
			fmt.Println(failure)
		}
	} else {
		fmt.Println("Testing mode completed successfully")
	}

	testRunTime := time.Now().UnixNano() - testStratTime
	fmt.Println(getExecutionTimeDisplay(testRunTime))
	fmt.Println("=====================================================")

	if len(assertionFailures) > 0 {
		os.Exit(1)
	}
}

func getExecutionTimeDisplay(executionTime int64) string {
	timeInMilliSeconds := executionTime / 1000000
	seconds := (timeInMilliSeconds / 1000)
	secondsDisplay := seconds % 60
	minutes := seconds / 60
	minutesDisplay := minutes % 60

	displayStatement := []byte("Execution Time: ")
	displayStatement = append(displayStatement, []byte(strconv.FormatInt(minutesDisplay, 10))...)
	displayStatement = append(displayStatement, []byte(":")...)
	if secondsDisplay <= 9 {
		displayStatement = append(displayStatement, []byte("0")...)
	}
	displayStatement = append(displayStatement, []byte(strconv.FormatInt(secondsDisplay, 10))...)
	return string(displayStatement)
}

func isReadyForTest(host string) (bool, *perfTestUtils.BasePerfStats) {

	//1) read in perf base stats
	basePerfstats, err := perfTestUtils.ReadBasePerfFile(host, configurationSettings.BaseStatsOutputDir)
	if err != nil {
		fmt.Println("Failed to read env stats for " + host + ". Error:" + err.Error() + ".")
		return false, nil
	}

	//2) validate content  of base stats file
	isBasePerfStatsValid := validateBasePerfStat(basePerfstats)
	if !isBasePerfStatsValid {
		fmt.Println("Base Perf stats are not fully populated for  " + host + ".")
		return false, nil
	}
	//3) Verify the number of base test cases is equal to the number of service test cases.
	correctNumberOfTests := perfTestUtils.ValidateTestDefinitionAmount(len(basePerfstats.BaseServiceResponseTimes), configurationSettings)
	if !correctNumberOfTests {
		return false, nil
	}

	return true, basePerfstats
}

func validateBasePerfStat(basePerfstats *perfTestUtils.BasePerfStats) bool {
	isBasePerfStatsValid := true

	if basePerfstats.BasePeakMemory <= 0 {
		isBasePerfStatsValid = false
	}
	if basePerfstats.GenerationDate == "" {
		isBasePerfStatsValid = false
	}
	if basePerfstats.ModifiedDate == "" {
		isBasePerfStatsValid = false
	}
	if len(basePerfstats.MemoryAudit) <= 0 {
		isBasePerfStatsValid = false
	}
	if basePerfstats.BaseServiceResponseTimes != nil {
		for _, baseResponseTime := range basePerfstats.BaseServiceResponseTimes {
			if baseResponseTime <= 0 {
				isBasePerfStatsValid = false
				break
			}
		}
	} else {
		isBasePerfStatsValid = false
	}
	return isBasePerfStatsValid
}

//This function does two thing,
//1 Start a go routine to preiodically grab the memory foot print and set the peak memory value
//2 Run all test using mock servers and gather performance stats
func runTests(perfStatsForTest *perfTestUtils.PerfStats, mode int) {

	var peakMemoryAllocation = new(uint64)
	//var lastServiceName = "StartUp"
	//var currentServiceName = "StartUp"

	memoryAudit := make([]uint64, 0)
	testPartitions := make([]perfTestUtils.TestPartition, 0)
	counter := 0
	testPartitions = append(testPartitions, perfTestUtils.TestPartition{Count: counter, TestName: "StartUp"})

	//Start go routine to grab memory in use
	//Peak memory is stored in peakMemoryAlocation variable.
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:

				memoryStatsUrl := "http://" + configurationSettings.TargetHost + ":" + configurationSettings.TargetPort + "/debug/vars"
				resp, err := http.Get(memoryStatsUrl)
				if err != nil {
					//log.Error("Memory analysis unavailable. Failed to retrieve memory Statistics from endpoint ", memoryStatsUrl)
					fmt.Println("Memory analysis unavailable. Failed to retrieve memory Statistics from endpoint ", memoryStatsUrl, ". Error:", err)
					quit <- true
				} else {

					body, _ := ioutil.ReadAll(resp.Body)

					defer resp.Body.Close()

					m := new(perfTestUtils.Entry)
					unmarshalErr := json.Unmarshal(body, m)
					if unmarshalErr != nil {
						//log.Error("Memory analysis unavailable. Failed to unmarshal memory statistics. ", unmarshalErr)
						fmt.Println("Memory analysis unavailable. Failed to unmarshal memory statistics. ", unmarshalErr)
						quit <- true
					} else {
						if m.Memstats.Alloc > *peakMemoryAllocation {
							*peakMemoryAllocation = m.Memstats.Alloc
						}
						memoryAudit = append(memoryAudit, m.Memstats.Alloc)
						counter++
						time.Sleep(time.Millisecond * 200)
					}
				}
			}
		}
	}()

	//Read test case files from test definition directory
	d, err := os.Open(configurationSettings.TestDefinitionsDir)
	if err != nil {
		//log.Error("Failed to open test definitions directory. Error:", err)
		fmt.Println("Failed to open test definitions directory. Error:", err)
		os.Exit(1)
	}
	defer d.Close()

	testSuite := new(perfTestUtils.TestSuite)
	if configurationSettings.TestSuite == "" {
		//If no test suite has been defined, treat and test case files as the suite
		fi, err := d.Readdir(-1)
		if err != nil {
			//log.Error("Failed to read files in test definitions directory. Error:", err)
			fmt.Println("Failed to read files in test definitions directory. Error:", err)
			os.Exit(1)
		}
		if len(fi) == 0 {
			//log.Error("No test case files found in specified directory ", configurationSettings.TestDefinitionsDir)
			fmt.Println("No test case files found in specified directory ", configurationSettings.TestDefinitionsDir)
			os.Exit(1)
		}
		testSuite.Name = "Default"
		for _, fi := range fi {
			bs, err := ioutil.ReadFile(configurationSettings.TestDefinitionsDir + "/" + fi.Name())
			if err != nil {
				//log.Error("Failed to read test file. Filename: ", fi.Name(), err)
				fmt.Println("Failed to read test file. Filename: ", fi.Name(), err)
				continue
			}

			testDefinition := new(perfTestUtils.TestDefinition)
			xml.Unmarshal(bs, &testDefinition)
			testSuite.TestCases = append(testSuite.TestCases, testDefinition)
		}
	} else {
		//If a test suite has been defined, load in all tests associated with the test suite.
		bs, err := ioutil.ReadFile(configurationSettings.TestDefinitionsDir + "/" + configurationSettings.TestSuite)
		if err != nil {
			//log.Error("Failed to read test file. Filename: ", fi.Name(), err)
			fmt.Println("Failed to read test file. Filename: ", configurationSettings.TestSuite, err)
		}
		xml.Unmarshal(bs, &testSuite)
	}

	//Determine load per concurrent user
	loadPerUser := int(configurationSettings.NumIterations / configurationSettings.ConcurrentUsers)
	remainder := configurationSettings.NumIterations % configurationSettings.ConcurrentUsers

	//Add a 1 second delay before running test case to allow the graph get some initial memory data before test cases are executed.
	time.Sleep(time.Second * 1)

	for index, testDefinition := range testSuite.TestCases {

		//log.Info("Running Test case [Name:", testDefinition.TestName, ", File name:", fi.Name(), "]")
		fmt.Println("Running Test case ", index, " [Name:", testDefinition.TestName, ", File name:", fi.Name(), "]")
		testPartitions = append(testPartitions, perfTestUtils.TestPartition{Count: counter, TestName: testDefinition.TestName})
		averageResponseTime := executeServiceTest(testDefinition, loadPerUser, remainder)
		if averageResponseTime > 0 {
			perfStatsForTest.ServiceResponseTimes[testDefinition.TestName] = averageResponseTime
		} else {
			if mode == TRAINING_MODE {
				//Fail fast on training mode if any requests fail. If training fails we cannot guarantee the results.
				fmt.Println("Training mode failed due to invalid response on service [Name:", testDefinition.TestName, ", File name:", fi.Name(), "]")
				os.Exit(1)
			}
		}
	}

	time.Sleep(time.Second * 1)
	perfStatsForTest.PeakMemory = *peakMemoryAllocation
	perfStatsForTest.MemoryAudit = memoryAudit
	perfStatsForTest.TestPartitions = testPartitions
}

//Single execution function for all service test.
//Runs multiple invocations of the test based on num iterations parameter
func executeServiceTest(testDefinition *perfTestUtils.TestDefinition, loadPerUser int, remainder int) int64 {

	averageResponseTime := int64(0)

	//responseTimes := make(perfTestUtils.RspTimes, configurationSettings.NumIterations)
	responseTimes := make([]int64, 0)

	subsetOfResponseTimesChan := make(chan perfTestUtils.RspTimes, 1)

	//Execute the test in a loop

	var wg sync.WaitGroup
	wg.Add(configurationSettings.ConcurrentUsers)
	for i := 0; i < configurationSettings.ConcurrentUsers; i++ {
		go buildAndSendUserRequests(subsetOfResponseTimesChan, loadPerUser, testDefinition)
		go aggregateResponseTimes(&responseTimes, subsetOfResponseTimesChan, &wg)
	}
	if remainder > 0 {
		wg.Add(1)
		go buildAndSendUserRequests(subsetOfResponseTimesChan, remainder, testDefinition)
		go aggregateResponseTimes(&responseTimes, subsetOfResponseTimesChan, &wg)
	}

	wg.Wait()

	if len(responseTimes) == configurationSettings.NumIterations {
		averageResponseTime = perfTestUtils.CalcAverageResponseTime(responseTimes, configurationSettings.NumIterations)
	}
	return averageResponseTime
}

func buildAndSendUserRequests(subsetOfResponseTimesChan chan perfTestUtils.RspTimes, loadPerUser int, testDefinition *perfTestUtils.TestDefinition) {
	responseTimes := make(perfTestUtils.RspTimes, loadPerUser)
	loopExecutedToCompletion := true

	for i := 0; i < loadPerUser; i++ {

		var req *http.Request

		if !testDefinition.Multipart {
			if testDefinition.Payload != "" {
				paylaod := testDefinition.Payload
				newPayload := substituteRequestValues(&paylaod)
				req, _ = http.NewRequest(testDefinition.HttpMethod, "http://"+configurationSettings.TargetHost+":"+configurationSettings.TargetPort+testDefinition.BaseUri, strings.NewReader(newPayload))
			} else {
				req, _ = http.NewRequest(testDefinition.HttpMethod, "http://"+configurationSettings.TargetHost+":"+configurationSettings.TargetPort+testDefinition.BaseUri, nil)
			}
		} else {
			if testDefinition.HttpMethod != "POST" {
				//log.Fatal("Multipart request has to be 'POST' method.")
				fmt.Println("Multipart request has to be 'POST' method.")
			} else {
				body := new(bytes.Buffer)
				writer := multipart.NewWriter(body)
				for _, field := range testDefinition.MultipartPayload {
					if field.FileName == "" {
						writer.WriteField(field.FieldName, field.FieldValue)
					} else {
						part, _ := writer.CreateFormFile(field.FieldName, field.FileName)
						io.Copy(part, bytes.NewReader(field.FileContent))
					}
				}
				writer.Close()
				req, _ = http.NewRequest(testDefinition.HttpMethod, "http://"+configurationSettings.TargetHost+":"+configurationSettings.TargetPort+testDefinition.BaseUri, body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
			}
		}

		//add headers
		for _, v := range testDefinition.Headers {
			req.Header.Add(v.Key, v.Value)
		}
		startTime := time.Now()
		if resp, err := (&http.Client{}).Do(req); err != nil {
			//log.Error("Error by firing request: ", req, "Error:", err)
			fmt.Println("Error by firing request: ", req, "Error:", err)
			loopExecutedToCompletion = false
			break
		} else {

			timeTaken := time.Since(startTime)

			body, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			//Validate service response
			contentLengthOk := perfTestUtils.ValidateResponseBody(body, testDefinition.TestName)
			responseCodeOk := perfTestUtils.ValidateResponseStatusCode(resp.StatusCode, testDefinition.ResponseStatusCode, testDefinition.TestName)
			responseTimeOK := perfTestUtils.ValidateServiceResponseTime(timeTaken.Nanoseconds(), testDefinition.TestName)

			if contentLengthOk && responseCodeOk && responseTimeOK {
				responseTimes[i] = timeTaken.Nanoseconds()
				extracResponseValues(testDefinition.TestName, body, testDefinition.ResponseProperties)
			} else {
				loopExecutedToCompletion = false
				break
			}
		}
	}

	if loopExecutedToCompletion {
		subsetOfResponseTimesChan <- responseTimes
	} else {
		subsetOfResponseTimesChan <- nil
	}
}

func substituteRequestValues(requestBody *string) string {

	requestPayloadCopy := *requestBody

	r := regexp.MustCompile("{{(.+)?}}")
	res := r.FindAllString(*requestBody, -1)

	if len(res) > 0 {
		for _, property := range res {
			//remove placeholder syntax
			cleanedPropertyName := strings.TrimPrefix(property, "{{")
			cleanedPropertyName = strings.TrimSuffix(cleanedPropertyName, "}}")
			//lookup value in the globals map
			value := globals[cleanedPropertyName]
			if value != "" {
				requestPayloadCopy = strings.Replace(requestPayloadCopy, property, value, 1)
			}
		}

	}
	return requestPayloadCopy
}

func extracResponseValues(testCaseName string, body []byte, resposneProperties []string) {
	for _, name := range resposneProperties {
		if globals[testCaseName+"."+name] == "" {
			r := regexp.MustCompile("<(.+)?:" + name + ">(.+)?</(.+)?:" + name + ">")
			res := r.FindStringSubmatch(string(body))
			globals[testCaseName+"."+name] = res[2]
		}
	}
}

func aggregateResponseTimes(responseTimes *[]int64, subsetOfResponseTimesChan chan perfTestUtils.RspTimes, wg *sync.WaitGroup) {
	subsetOfResponseTimes := <-subsetOfResponseTimesChan
	if subsetOfResponseTimes != nil {
		*responseTimes = append(*responseTimes, subsetOfResponseTimes...)
	}
	wg.Done()
}

//This function runs the assertions to ensure memory and service have not deviated past the allowed variance
func runAssertions(basePerfstats *perfTestUtils.BasePerfStats, perfStats *perfTestUtils.PerfStats) []string {

	assertionFailures := make([]string, 0)

	//Asserts Peak memory growth has not exceeded the allowable variance
	peakMemoryVariancePercentage := perfTestUtils.CalcPeakMemoryVariancePercentage(basePerfstats.BasePeakMemory, perfStats.PeakMemory)
	varianceOk := perfTestUtils.ValidatePeakMemoryVariance(configurationSettings.AllowablePeakMemoryVariance, peakMemoryVariancePercentage)
	if !varianceOk {
		assertionFailures = append(assertionFailures, fmt.Sprintf("Memory Failure: Peak variance exceeded by %3.2f %1s", peakMemoryVariancePercentage, "%"))
	}

	for serviceName, baseResponseTime := range basePerfstats.BaseServiceResponseTimes {
		averageServiceResponseTime := perfStats.ServiceResponseTimes[serviceName]
		if averageServiceResponseTime == 0 {
			assertionFailures = append(assertionFailures, fmt.Sprintf("Service Failure: Service test %-60s did not execute correctly. See logs for more details.", serviceName))
		}

		responseTimeVariancePercentage := perfTestUtils.CalcAverageResponseVariancePercentage(averageServiceResponseTime, baseResponseTime)
		varianceOk := perfTestUtils.ValidateAverageServiceResponeTimeVariance(configurationSettings.AllowableServiceResponseTimeVariance, responseTimeVariancePercentage, serviceName)
		if !varianceOk {
			assertionFailures = append(assertionFailures, fmt.Sprintf("Service Failure: Service test %-60s response time variance exceeded by %3.2f %1s", serviceName, responseTimeVariancePercentage, "%"))
		}

	}
	return assertionFailures
}
