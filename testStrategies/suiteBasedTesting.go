package testStrategies

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/xtracdev/automated-perf-test/perfTestUtils"
	"sync"
	"sync/atomic"
	"time"
)

// ExecuteTestSuiteWrapper executes suites using concurrent goroutines and
// returns response time metrics.
func ExecuteTestSuiteWrapper(
	testSuite *TestSuite,
	configSettings *perfTestUtils.Config,
	perfStatsForTest *perfTestUtils.PerfStats,
	scenarioTimeStart time.Time,
) map[string][]int64 {
	allServicesResponseTimesMap := make(map[string][]int64, 0)
	testSuiteResponseTimesChan := make(chan []map[string]int64, 1)
	var suiteWaitGroup sync.WaitGroup

	// Set concurrency control:
	suiteWaitGroup.Add(configSettings.ConcurrentUsers)

	// Run the test suites concurrently.
	for i := 0; i < configSettings.ConcurrentUsers; i++ {
		// If RampUsers is set, start a batch of user threads, then wait the
		// specified delay time. Otherwise, skip the delay and start all
		// threads simultaneously.
		if (i != 0) && (configSettings.RampUsers != 0) && (i%configSettings.RampUsers == 0) {
			time.Sleep(time.Duration(configSettings.RampDelay) * time.Second)
		}
		go executeTestSuite(testSuiteResponseTimesChan, testSuite, configSettings, i, GlobalsLockCounter, perfStatsForTest)
		go aggregateSuiteResponseTimes(testSuiteResponseTimesChan, allServicesResponseTimesMap, &suiteWaitGroup)
	}

	// Display the ongoing TPS to log.Info based on period specified in configurationSettings.TPSFreq:
	quitShowTPSChan := make(chan bool)
	go showCurrentTPS(quitShowTPSChan, configSettings, scenarioTimeStart, &perfStatsForTest.OverAllTransCount)

	suiteWaitGroup.Wait()
	quitShowTPSChan <- true

	return allServicesResponseTimesMap
}

//----- executeTestSuite ------------------------------------------------------
func executeTestSuite(
	testSuiteResponseTimesChan chan []map[string]int64,
	testSuite *TestSuite,
	configurationSettings *perfTestUtils.Config,
	userID int,
	globalsMap GlobalsMaps,
	perfStatsForTest *perfTestUtils.PerfStats,
) {
	log.Info("Test Suite started")

	allSuiteResponseTimes := make([]map[string]int64, 0)
	uniqueTestRunID := ""

	for i := 0; i < configurationSettings.NumIterations; i++ {
		// Run all services of the test suite NumIterations of times.
		uniqueTestRunID = fmt.Sprintf("User%dIter%d", userID, i)
		testSuiteResponseTimes := make(map[string]int64)

		// Set booleans for weighted load tags:
		// Determine whether "Infrequent" items should run this iteration.
		// [Currently set at 20% (mod 5)]
		skipInfrequent := false
		if i%5 != 0 {
			skipInfrequent = true
		}
		// Determine whether "Sparse" items should run this iteration.
		// [Currently set at 8% (mod 12)]
		skipSparse := false
		if i%12 != 0 {
			skipSparse = true
		}

		for _, testDefinition := range testSuite.TestCases {
			// Execute service based on weighted load:
			if testDefinition.ExecWeight == "Infrequent" && skipInfrequent {
				// Skip "Infrequent" items:
				continue
			}
			if testDefinition.ExecWeight == "Sparse" && skipSparse {
				// Skip "Sparse" items:
				continue
			}

			// DEBUG:
			if testDefinition.ExecWeight == "Infrequent" {
				log.Debug("ExecWeight = [", testDefinition.ExecWeight, "] testCase = [", testDefinition.TestName, "]")
			}
			if testDefinition.ExecWeight == "Sparse" {
				log.Debug("ExecWeight = [", testDefinition.ExecWeight, "] testCase = [", testDefinition.TestName, "]")
			}

			log.Info("Test case: [", testDefinition.TestName, "] UniqueRunID: [", uniqueTestRunID, "]")

			targetHost, targetPort := determineHostandPortforRequest(testDefinition, configurationSettings)
			responseTime := testDefinition.BuildAndSendRequest(configurationSettings.RequestDelay, targetHost, targetPort, uniqueTestRunID, globalsMap)
			testSuiteResponseTimes[testDefinition.TestName] = responseTime

			// Update the concurrent counters.
			// Overall counter:
			atomic.AddUint64(&perfStatsForTest.OverAllTransCount, 1)

			// Service-level counters:
			// Create the counters on the fly and increment atomically.
			if perfStatsForTest.ServiceTransCount[testDefinition.TestName] == nil {
				perfStatsForTest.ServiceTransCount[testDefinition.TestName] = new(uint64)
				atomic.StoreUint64(
					perfStatsForTest.ServiceTransCount[testDefinition.TestName],
					0,
				)
			}
			atomic.AddUint64(
				perfStatsForTest.ServiceTransCount[testDefinition.TestName],
				1,
			)
		}

		allSuiteResponseTimes = append(allSuiteResponseTimes, testSuiteResponseTimes)

		globalsMap.Lock()
		globalsMap.m[uniqueTestRunID] = nil
		globalsMap.Unlock()
	}

	testSuiteResponseTimesChan <- allSuiteResponseTimes
}

//----- aggregateSuiteResponseTimes -------------------------------------------
func aggregateSuiteResponseTimes(
	testSuiteResponseTimesChan chan []map[string]int64,
	allServicesResponseTimesMap map[string][]int64,
	suiteWaitGroup *sync.WaitGroup,
) {
	perUserSuiteResponseTimes := <-testSuiteResponseTimesChan
	for _, singleSuiteRunResponseTimes := range perUserSuiteResponseTimes {
		for serviceName, serviceResponseTime := range singleSuiteRunResponseTimes {
			if allServicesResponseTimesMap[serviceName] == nil {
				serviceResponseSlice := make([]int64, 0)
				allServicesResponseTimesMap[serviceName] = serviceResponseSlice
			}
			allServicesResponseTimesMap[serviceName] = append(allServicesResponseTimesMap[serviceName], serviceResponseTime)
		}
	}
	suiteWaitGroup.Done()
}

//----- showCurrentTPS -------------------------------------------------------------------------------------------------
// Print current TPS progress every period of time defined by configurationSettings.TPSFREQ.
func showCurrentTPS(
	chQuit chan bool,
	confgSettings *perfTestUtils.Config,
	scenarioStartTime time.Time,
	nNumberOfTrans *uint64,
) {
	for {
		// Concurrent controls:
		select {
		case <-chQuit:
			return
		default:
			// Set variables for convenience.
			durElapsedTime := time.Since(scenarioStartTime)
			numOps := atomic.LoadUint64(nNumberOfTrans)

			// We only want one output line during any given second. This
			// effectively sets the lower bound for TPSFreq to one second.
			time.Sleep(time.Second)

			// No need to display until at least one operation has completed.
			if numOps < uint64(1) {
				break
			}

			// No need to display if not within the period set in config:
			if int64(durElapsedTime.Seconds())%int64(confgSettings.TPSFreq) != 0 {
				break
			}

			// Print the display.
			tps := 0.0
			if int(durElapsedTime.Seconds()) > 0 {
				tps = (float64(numOps) / durElapsedTime.Seconds())
			}

			log.Infof("[showCurrentTPS] {\"TransCount\":\"%d\",\"ElapsedTime\":\"%v\",\"TPS\":\"%f\"}",
				numOps,
				durElapsedTime,
				tps,
			)
		}
	}
}
