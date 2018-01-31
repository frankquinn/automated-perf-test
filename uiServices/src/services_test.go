package services

import (
	"fmt"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"github.com/xtracdev/automated-perf-test/perfTestUtils"
)

const validJson = `{
        "apiName": "ServiceTestConfig",
       "targetHost": "localhost",
       "targetPort": "9191",
       "numIterations": 1000,
       "allowablePeakMemoryVariance": 30,
       "allowableServiceResponseTimeVariance": 30,
       "testCaseDir": "./definitions/testCases",
       "testSuiteDir": "./definitions/testSuites",
        "baseStatsOutputDir": "./envStats",
       "reportOutputDir": "./report",
       "concurrentUsers": 50,
       "testSuite": "Default-1",
       "memoryEndpoint": "/alt/debug/vars",
       "requestDelay": 5000,
       "TPSFreq": 30,
       "rampUsers": 5,
       "rampDelay": 15
       }`

const invalidJsonMissingFields = `{
        "apiName": "ServiceTestConfi",
       "targetHost": "localhost",
       "targetPort": "9191",
       "numIterations": 1000,
       "allowablePeakMemoryVariance": -1,
       "allowableServiceResponseTimeVariance": 30,
       "testCaseDir": "./definitions/testCases",
       "testSuiteDir": "./definitions/testSuites",
        "baseStatsOutputDir": "./envStats",
       "reportOutputDir": "./report",
       "concurrentUsers": 50,
       "testSuite": "",
       "memoryEndpoint": "/alt/debug/vars",
       "requestDelay": 5000,
       "TPSFreq": 30,
       "rampUsers": 5,
       "rampDelay": 0
       }`

const invalidJson = `{
        "apiName"://*()()(),
       "targetHost": 0,
       "targetPort": 0,
       "numIterations": "x",
       "allowablePeakMemoryVariance": 30,
       "allowableServiceResponseTimeVariance": 30,
       "testCaseDir": "./definitions/testCases",
       "testSuiteDir": "./definitions/testSuites",
        "baseStatsOutputDir": "./envStats",
       "reportOutputDir": "./report",
       "concurrentUsers": 50,
       "testSuite": "suiteFileName.xml",
       "memoryEndpoint": "/alt/debug/vars",
       "requestDelay": 5000,
       "TPSFreq": 30,
       "rampUsers": 5,
       "rampDelay": 15
       }`

func TestFilePathExist(t *testing.T) {
	path := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	actual := false
	fmt.Println(path)
	actual = FilePathExist(path)
	expected := true
	assert.Equal(t, expected, actual)
}

func TestFilePathDoesNotExist(t *testing.T) {
	path := "((((((("
	actual := FilePathExist(path)
	expected := false
	assert.Equal(t, expected, actual)
}

func TestInvalidJsonPostMissingRequiredField(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(invalidJsonMissingFields)
	r.HandleFunc("/configs", postConfigs)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPost, "/configs", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("TestValidJsonPost. Expected:", http.StatusBadRequest, " Got:", w.Code, "  Error. Did not succesfully post")
	}
}

func TestValidJsonPost(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validJson)
	r.HandleFunc("/configs", postConfigs)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPost, "/configs", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("TestValidJsonPost. Expected:", http.StatusCreated, " Got:", w.Code, "  Error. Did not succesfully post")
	}
}

func TestPostWithInvalidHeader (t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validJson)
	r.HandleFunc("/configs", postConfigs)

	filePath := "xxxxxx"
	request, err := http.NewRequest(http.MethodPost, "/configs", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("TestValidJsonPost. Expected:", http.StatusBadRequest ," Got:", w.Code, "  Error. Did not succesfully post")
	}
}

func TestFilePathEndsWIthSlash(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validJson)
	r.HandleFunc("/configs", postConfigs)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test"

	request, err := http.NewRequest(http.MethodPost, "/configs", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("TestFilePathEndsWith'/'.  Expected:", http.StatusCreated, " Got:", w.Code, "  Error. Did not succesfully post")
	}
}

func TestInvalidJsonPost(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(invalidJson)
	r.HandleFunc("/configs", postConfigs)

	request, err := http.NewRequest(http.MethodPost, "/configs", reader)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("TestInvalidJsonPost.  Expected:", http.StatusBadRequest, " Got:", w.Code, "Error. Did not succesfully post ")
	}
}

func TestWhenConfigPathDirEmpty(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validJson)
	r.HandleFunc("/configs", postConfigs)

	request, err := http.NewRequest(http.MethodPost, "/configs", reader)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("TestWhenConfigPathDirEmpty.  Expected:", http.StatusBadRequest, " Got:", w.Code, "Error. ConfigPathDir is Empty ")
	}
}

func TestInvalidURL(t *testing.T) {
	pt := perfTestUtils.Config{}
	writerXml(pt, "/path/xxx")
}


func TestSuccessfulGet(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	// create file to GET
	reader := strings.NewReader(validJson)
	r.HandleFunc("/configs", postConfigs)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPost, "/configs", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	//prepare GET request
	filePath = os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err = http.NewRequest(http.MethodGet, "/configs/ServiceTestConfig", nil)

	request.Header.Set("configPathDir", filePath)
	request.Header.Get("configPathDir")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("TestSuccessfulGET. Expected:", http.StatusOK, " Got:", w.Code, "  Error. Did not succesfully get")
	}
}

func TestSuccessfulGetPathWihoutSlash(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	r.HandleFunc("/configs", getConfigs)
	//no slash at end of filepath header
	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test"
	request, err := http.NewRequest(http.MethodGet, "/configs/ServiceTestConfig", nil)

	request.Header.Set("configPathDir", filePath)
	request.Header.Get("configPathDir")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Test Get Path ends with backslash. Expected:", http.StatusOK, " Got:", w.Code, "  Error. Did not succesfully get")
	}
}

func TestGetNoHeaderPath(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	r.HandleFunc("/configs", getConfigs)

	filePath := ""
	request, err := http.NewRequest(http.MethodGet, "/configs/serviceTestConfig", nil)

	request.Header.Set("configPathDir", filePath)
	request.Header.Get("configPathDir")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("Test No-Header-Get. Expected:", http.StatusBadRequest, " Got:", w.Code)
	}
}

func TestGetFileNotFound(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	r.HandleFunc("/configs", getConfigs)


	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodGet, "/configs/xxx.java", nil)

	request.Header.Set("configPathDir", filePath)
	request.Header.Get("configPathDir")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	if w.Code != http.StatusNotFound {
		t.Errorf("Test File Not Found. Expected:", http.StatusNotFound, " Got:", w.Code)
	}
}

func TestValidJsonPut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validJson)
	r.HandleFunc("/configs", putConfigs)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPut, "/configs/ServiceTestConfig", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, w.Code, http.StatusNoContent, "Did Not successfully Update")
}

func TestMissingFieldPut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(invalidJsonMissingFields)
	r.HandleFunc("/configs", putConfigs)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPut, "/configs/ServiceTestConfig", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, w.Code, http.StatusBadRequest, "Sucessfully updated. Field Should be missing so update shouldn't occur")
}

func TestInvalidJsonPut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(invalidJson)
	r.HandleFunc("/configs", putConfigs)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPut, "/configs/ServiceTestConfig", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, w.Code, http.StatusBadRequest, "Sucessfully updated. Field data type should have been incorrect so update should occur")
}

func TestInvalidUrlPut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validJson)
	r.HandleFunc("/configs", putConfigs)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPut, "/configs/xxx", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, w.Code, http.StatusConflict, "Sucessfully updated. Should have have worked using /configs/xxx")
}

func TestNoUrlPut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validJson)
	r.HandleFunc("/configs", putConfigs)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPut, "", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, w.Code, http.StatusNotFound, "Sucessfully updated. Should not have worked with no URL")
}

func TestSuccessfulPutWithNoPathSlash(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validJson)
	r.HandleFunc("/configs", putConfigs)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test"
	request, err := http.NewRequest(http.MethodPut, "/configs/ServiceTestConfig", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, w.Code, http.StatusNoContent, "Did not update. Should have added '/' to path")
}
func TestNoPathPut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validJson)
	r.HandleFunc("/configs", putConfigs)

	filePath := ""
	request, err := http.NewRequest(http.MethodPut, "/configs/ServiceTestConfig", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, w.Code, http.StatusBadRequest, "Successfully updated. Should not have worked due to no filepath")
}

func TestNoFileNamePut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validJson)
	r.HandleFunc("/configs", putConfigs)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test"
	request, err := http.NewRequest(http.MethodPut, "/configs", reader)
	request.Header.Set("configPathDir", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, w.Code, http.StatusNotFound, "Successfully updated. Should not have worked due to no file name given")
}

