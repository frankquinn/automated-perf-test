package perfTestUtils

import (
	"encoding/json"
	"fmt"
	//log "github.com/Sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

//FileSystem is an interface to access os filesystem or mock it
type FileSystem interface {
	Open(name string) (File, error)
	Create(name string) (File, error)
}

//File is an interface to access os.File or mock it
type File interface {
	Readdir(n int) (fi []os.FileInfo, err error)
	io.WriteCloser
	Read(p []byte) (n int, err error)
}

// OsFS implements fileSystem using the local disk.
type OsFS struct{}

//Open calls os function
func (OsFS) Open(name string) (File, error) { return os.Open(name) }

//Create calls os function
func (OsFS) Create(name string) (File, error) { return os.Create(name) }

//=============================
//Testing run utility functions
//=============================
//This function reads a base perf and converts it to a base perf struct
func ReadBasePerfFile(r io.Reader) (*BasePerfStats, error) {
	basePerfstats := &BasePerfStats{
		BaseServiceResponseTimes: make(map[string]int64),
		MemoryAudit:              make([]uint64, 0),
	}
	var errorFound error

	content, err := ioutil.ReadAll(r)
	if err != nil {
		errorFound = err
	} else {
		jsonError := json.Unmarshal(content, basePerfstats)
		if jsonError != nil {
			errorFound = jsonError
		}
	}
	return basePerfstats, errorFound
}

func ValidateTestDefinitionAmount(baselineAmount int, configurationSettings *Config, fs FileSystem) bool {

	d, err := fs.Open(configurationSettings.TestDefinitionsDir)
	if err != nil {
		fmt.Println("Failed to open test definitions directory. Error:", err)
		os.Exit(1)
	}
	defer d.Close()
	fi, err := d.Readdir(-1)
	if err != nil {
		fmt.Println("Failed to read files in test definitions directory. Error:", err)
		os.Exit(1)
	}

	definitionAmount := len(fi)

	fmt.Println("Number of defined test cases:", definitionAmount)
	fmt.Println("Number of base line test cases:", baselineAmount)
	if definitionAmount != baselineAmount {
		fmt.Printf("Amount of test definition: %d does not equal to baseline amount: %d.\n", definitionAmount, baselineAmount)
		return false
	}
	return true
}

//=====================
//Calc Memory functions
//=====================
func CalcPeakMemoryVariancePercentage(basePeakMemory uint64, peakMemory uint64) float64 {

	peakMemoryVariancePercentage := float64(0)

	if basePeakMemory < peakMemory {
		peakMemoryDelta := peakMemory - basePeakMemory
		temp := float64(float64(peakMemoryDelta) / float64(basePeakMemory))
		peakMemoryVariancePercentage = temp * 100
	} else {
		peakMemoryDelta := basePeakMemory - peakMemory
		temp := float64(float64(peakMemoryDelta) / float64(basePeakMemory))
		peakMemoryVariancePercentage = (temp * 100) * -1
	}

	return peakMemoryVariancePercentage
}

//============================
//Calc Response time functions
//============================
func CalcAverageResponseTime(responseTimes RspTimes, numIterations int) int64 {

	averageResponseTime := int64(0)

	//Remove the highest 5% to take out anomolies
	sort.Sort(responseTimes)
	numberToRemove := int(float32(numIterations) * float32(0.05))
	fmt.Printf("resp times length: %v\n", len(responseTimes))
	fmt.Printf("To remove: %v\n", numberToRemove)
	responseTimes = responseTimes[0 : len(responseTimes)-numberToRemove]

	totalOfAllresponseTimes := int64(0)
	for _, val := range responseTimes {
		totalOfAllresponseTimes = totalOfAllresponseTimes + val
	}
	fmt.Printf("totalOfAllresponseTimes = %v\n", totalOfAllresponseTimes)
	averageResponseTime = int64(float64(totalOfAllresponseTimes) / float64(numIterations-numberToRemove))

	return averageResponseTime
}

func CalcAverageResponseVariancePercentage(averageResponseTime int64, baseResponseTime int64) float64 {

	responseTimeVariancePercentage := float64(0)

	if baseResponseTime < averageResponseTime {
		delta := uint64(averageResponseTime) - uint64(baseResponseTime)
		temp := float64(float64(delta) / float64(baseResponseTime))
		responseTimeVariancePercentage = temp * 100
	} else {
		delta := baseResponseTime - averageResponseTime
		temp := float64(float64(delta) / float64(baseResponseTime))
		responseTimeVariancePercentage = (temp * 100) * -1
	}

	return responseTimeVariancePercentage
}

//=====================================
//Service response validation functions
//=====================================
func ValidateResponseBody(body []byte, testName string) bool {

	isResponseBodyValid := false
	if len(body) > 0 {
		isResponseBodyValid = true
	} else {
		fmt.Printf("Incorrect Content lenght (%d) returned for service %s", len(body), testName)
	}
	return isResponseBodyValid
}

func ValidateResponseStatusCode(responseStatusCode int, expectedStatusCode int, testName string) bool {

	isResponseStatusCodeValid := false
	if responseStatusCode == expectedStatusCode {
		isResponseStatusCodeValid = true
	} else {
		fmt.Printf("Incorrect status code of %d retruned for service %s. %d expected", responseStatusCode, testName, expectedStatusCode)
	}
	return isResponseStatusCodeValid
}

func ValidateServiceResponseTime(responseTime int64, testName string) bool {

	isResponseTimeValid := false
	if responseTime > 0 {
		isResponseTimeValid = true
	} else {
		fmt.Printf("Time taken to complete request %s was 0 nanoseconds", testName)
	}
	return isResponseTimeValid
}

//=====================================
//Test Assertion functions
//=====================================
func ValidatePeakMemoryVariance(allowablePeakMemoryVariance float64, peakMemoryVariancePercentage float64) bool {

	if allowablePeakMemoryVariance >= peakMemoryVariancePercentage {
		return true
	} else {
		return false
	}
}

/*func ValidateTestCaseCount(baseTestCaseCount int, testTestCaseCount int) bool {

	isTestCaseCountValid := false
	if baseTestCaseCount == testTestCaseCount {
		isTestCaseCountValid = true
	} else {
		fmt.Printf("Number of service tests in base is differnet to the number of services for this test run.")
	}
	return isTestCaseCountValid
}*/

func ValidateAverageServiceResponeTimeVariance(allowableServiceResponseTimeVariance float64, serviceResponseTimeVariancePercentage float64, serviceName string) bool {
	if allowableServiceResponseTimeVariance >= serviceResponseTimeVariancePercentage {
		return true
	} else {
		return false
	}
}

//=====================================
//Response times sort functions
//=====================================
func (a RspTimes) Len() int           { return len(a) }
func (a RspTimes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a RspTimes) Less(i, j int) bool { return a[i] < a[j] }

//==============================================
//Generate base environment stats file functions
//==============================================
func populateBasePerfStats(perfStatsForTest *PerfStats, basePerfstats *BasePerfStats, reBaseMemory bool) {
	modified := false

	//Setting memory data
	if basePerfstats.BasePeakMemory == 0 || reBaseMemory {
		basePerfstats.BasePeakMemory = perfStatsForTest.PeakMemory
		modified = true
	}
	if basePerfstats.MemoryAudit == nil || len(basePerfstats.MemoryAudit) == 0 || reBaseMemory {
		basePerfstats.MemoryAudit = perfStatsForTest.MemoryAudit
		modified = true
	}

	//Setting service response time data
	for serviceName, responseTime := range perfStatsForTest.ServiceResponseTimes {
		serviceBaseResponseTime := basePerfstats.BaseServiceResponseTimes[serviceName]
		if serviceBaseResponseTime == 0 {
			basePerfstats.BaseServiceResponseTimes[serviceName] = responseTime
			modified = true
		}
	}

	//Setting time stamps
	currentTime := time.Now().Format(time.RFC850)
	if basePerfstats.GenerationDate == "" {
		basePerfstats.GenerationDate = currentTime
	}
	if modified {
		basePerfstats.ModifiedDate = currentTime
	}
}

func GenerateEnvBasePerfOutputFile(perfStatsForTest *PerfStats, basePerfstats *BasePerfStats, configurationSettings *Config, exit func(code int), fs FileSystem) {

	//Set base performance based on training test run
	populateBasePerfStats(perfStatsForTest, basePerfstats, configurationSettings.ReBaseMemory)

	//Convert base perf stat to Json and write out to file
	basePerfstatsJson, err := json.Marshal(basePerfstats)
	if err != nil {
		fmt.Println("Failed to marshal to Json. Error:", err)
		exit(1)
	}
	file, err := fs.Create(configurationSettings.BaseStatsOutputDir + "/" + configurationSettings.ExecutionHost + "-perfBaseStats")
	if err != nil {
		fmt.Println("Failed to create output file. Error:", err)
		exit(1)
	}
	if file != nil {
		defer file.Close()
		file.Write(basePerfstatsJson)
	}
}
