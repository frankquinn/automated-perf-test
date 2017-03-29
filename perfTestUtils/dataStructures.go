package perfTestUtils

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"runtime"
	"strings"
	"time"
)

const (
	defaultAPIName                              = "Default_API_NAME"
	defaultTargetHost                           = "localhost"
	defaultTargetPort                           = "8080"
	defaultNumIterations                        = 1000
	defaultAllowablePeakMemoryVariance          = float64(15)
	defaultAllowableServiceResponseTimeVariance = float64(15)
	defaultTestCaseDir                          = "./definitions/testCases"
	defaultTestSuiteDir                         = "./definitions/testSuites"
	defaultBaseStatsOutputDir                   = "./envStats"
	defaultReportOutputDir                      = "./"
	defaultConcurrentUsers                      = 1
	defaultTestSuite                            = ""
	defaultMemoryEndpoint                       = "/debug/vars"
	defaultRequestDelay                         = 1
	defaultTPSFreq                              = 5

	// Template file
	defaultTestFileFormat                       = "xml"
)

type Config struct {
	APIName                              string  `xml:"apiName"`
	TargetHost                           string  `xml:"targetHost"`
	TargetPort                           string  `xml:"targetPort"`
	NumIterations                        int     `xml:"numIterations"`
	AllowablePeakMemoryVariance          float64 `xml:"allowablePeakMemoryVariance"`
	AllowableServiceResponseTimeVariance float64 `xml:"allowableServiceResponseTimeVariance"`
	TestCaseDir                          string  `xml:"testCaseDir"`
	TestSuiteDir                         string  `xml:"testSuiteDir"`
	BaseStatsOutputDir                   string  `xml:"baseStatsOutputDir"`
	ReportOutputDir                      string  `xml:"reportOutputDir"`
	ConcurrentUsers                      int     `xml:"concurrentUsers"`
	TestSuite                            string  `xml:"testSuite"`
	MemoryEndpoint                       string  `xml:"memoryEndpoint"`
	RequestDelay                         int     `xml:"requestDelay"`
	TPSFreq                              int     `xml:"TPSFreq"`

	//These value can only be set by command line arguments as they control each training and test run.
	GBS          bool
	ReBaseMemory bool
	ReBaseAll    bool

	//This value is determined by the environment/machine on which the test is being run.
	ExecutionHost string

	//Template file
	ReportTemplateFile string `xml:"reportTemplateFile,omitempty"`
	ConfigFileFormat   string
	TestFileFormat     string
}

func (c *Config) SetDefaults() {
	c.APIName = defaultAPIName
	c.TargetHost = defaultTargetHost
	c.TargetPort = defaultTargetPort
	c.NumIterations = defaultNumIterations
	c.AllowablePeakMemoryVariance = defaultAllowablePeakMemoryVariance
	c.AllowableServiceResponseTimeVariance = defaultAllowableServiceResponseTimeVariance
	c.TestCaseDir = defaultTestCaseDir
	c.TestSuiteDir = defaultTestSuiteDir
	c.BaseStatsOutputDir = defaultBaseStatsOutputDir
	c.ReportOutputDir = defaultReportOutputDir
	c.ConcurrentUsers = defaultConcurrentUsers
	c.TestSuite = defaultTestSuite
	c.TestFileFormat = defaultTestFileFormat
	c.MemoryEndpoint = defaultMemoryEndpoint
	c.RequestDelay = defaultRequestDelay
	c.TPSFreq = defaultTPSFreq

	c.GBS = false
	c.ReBaseMemory = false
	c.ReBaseAll = false
}

func (c Config) PrintAndValidateConfig() {

	if strings.TrimSpace(c.APIName) == "" {
		c.APIName = defaultAPIName
	}
	if strings.TrimSpace(c.TargetHost) == "" {
		c.TargetHost = defaultTargetHost
	}
	if strings.TrimSpace(c.TargetPort) == "" {
		c.TargetPort = defaultTargetPort
	}
	if c.NumIterations < 1 {
		c.NumIterations = defaultNumIterations
	}
	if c.ConcurrentUsers < 1 {
		c.ConcurrentUsers = defaultConcurrentUsers
	}
	if c.AllowablePeakMemoryVariance <= 0.0 {
		c.AllowablePeakMemoryVariance = defaultAllowablePeakMemoryVariance
	}
	if c.AllowableServiceResponseTimeVariance <= 0.0 {
		c.AllowableServiceResponseTimeVariance = defaultAllowableServiceResponseTimeVariance
	}
	if strings.TrimSpace(c.TestCaseDir) == "" {
		c.TestCaseDir = defaultTestCaseDir
	}
	if strings.TrimSpace(c.BaseStatsOutputDir) == "" {
		c.BaseStatsOutputDir = defaultBaseStatsOutputDir
	}
	if strings.TrimSpace(c.ReportOutputDir) == "" {
		c.ReportOutputDir = defaultReportOutputDir
	}
	if strings.TrimSpace(c.MemoryEndpoint) == "" {
		c.MemoryEndpoint = defaultMemoryEndpoint
	}
	if c.RequestDelay < 1 {
		c.MemoryEndpoint = defaultMemoryEndpoint
	}
	if c.TPSFreq < 1 {
		c.TPSFreq = defaultTPSFreq
	}

	configOutput := []byte("")
	configOutput = append(configOutput, []byte("\n============== Configuration Settings =========\n")...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90s %2s", "apiName", c.APIName, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90s %2s", "targetHost", c.TargetHost, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90s %2s", "targetPort", c.TargetPort, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90d %2s", "numIterations", c.NumIterations, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90d %2s", "concurrentUsers", c.ConcurrentUsers, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90.2f %2s", "allowablePeakMemoryVariance", c.AllowablePeakMemoryVariance, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90.2f %2s", "allowableServiceResponseTimeVariance", c.AllowableServiceResponseTimeVariance, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90s %2s", "testCaseDir", c.TestCaseDir, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90s %2s", "testSuiteDir", c.TestSuiteDir, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90s %2s", "testSuite", c.TestSuite, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90s %2s", "memoryEndpoint", c.MemoryEndpoint, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90d %2s", "requestDelay", c.RequestDelay, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90s %2s", "baseStatsOutputDir", c.BaseStatsOutputDir, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90s %2s", "reportOutputDir", c.ReportOutputDir, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90t %2s", "gbs", c.GBS, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90t %2s", "reBaseMemory", c.ReBaseMemory, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90t %2s", "reBaseAll", c.ReBaseAll, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90s %2s", "ExecutionHost", c.ExecutionHost, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90d %2s", "RequestDelay", c.RequestDelay, "\n"))...)
	configOutput = append(configOutput, []byte(fmt.Sprintf("%-45s %-90d %2s", "TPSFreq", c.TPSFreq, "\n"))...)
	configOutput = append(configOutput, []byte("\n=================================================\n")...)
	log.Info(string(configOutput))
}

//This struct defines the base performance statistics
type BasePerfStats struct {
	GenerationDate           string           `json:"GenerationDate"`
	ModifiedDate             string           `json:"ModifiedDate"`
	BasePeakMemory           uint64           `json:"BasePeakMemory"`
	BaseServiceResponseTimes map[string]int64 `json:"BaseServiceResponseTimes"`
	MemoryAudit              []uint64         `json:"MemoryAudit"`
}

//This struct defines the performance statistics for this test run
type PerfStats struct {
	PeakMemory           uint64
	ServiceResponseTimes map[string]int64
	ServiceTransCount    map[string]*uint64
	ServiceTPS           map[string]float64
	OverAllTransCount    uint64
	OverAllTPS           float64
	MemoryAudit          []uint64
	TestPartitions       []TestPartition
	TestDate             time.Time
}

func (ps *PerfStats) GetTestTime() string {
	return ps.TestDate.Format(time.RFC850)
}

type TestPartition struct {
	Count    int
	TestName string
}

type Entry struct {
	Cmdline  []string         `json:"cmdline"`
	Memstats runtime.MemStats `json:"memstats"`
}

type RspTimes []int64
