package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

const validTestCase = `
{
   "testname":"TestCaseService",
   "description":"desc",
   "overrideHost":"host",
   "overridePort":"9191",
   "httpMethod":"GET",
   "baseUri": "path/to/URI",
   "multipart":false,
   "payload": "payload",
   "responseStatusCode":200,
   "responseContentType": "JSON" ,
   "preThinkTime": 1000,
   "postThinkTime":2000,
   "execWeight": "Sparse",
   "headers":[{
   	 "key": "Authorization",
     "value" :"Header-Value"
   }],
  "responseValues":[{
     "value":"Res-Value",
     "extractionKey": "Res-Key"
  }],
  "multipartPayload":[{
     "fieldName": "F-Name",
   	 "fieldValue":"PayloadName",
     "fileName": "file-name"
  }]
}
`

const TestCaseNoName = `
{
   "testname":"",
   "description":"desc",
   "overrideHost":"host",
   "overridePort":"9191",
   "httpMethod":"GET",
   "baseUri": "path/to/URI",
   "multipart":false,
   "payload": "payload",
   "responseStatusCode":200,
   "responseContentType": "JSON" ,
   "preThinkTime": 1000,
   "postThinkTime":2000,
   "execWeight": "Sparse",
   "headers":[{
   	 "key": "Authorization",
     "value" :"Header-Value"
   }],
  "responseValues":[{
     "value":"Res-Value",
     "extractionKey": "Res-Key"
  }],
  "multipartPayload":[{
     "fieldName": "F-Name",
   	 "fieldValue":"PayloadName",
     "fileName": "file-name"
  }]
}
`

const TestCaseMissingRequired = `
{
   "testname":"",
   "description":"",
   "overrideHost":"",
   "overridePort":"",
   "baseUri": "path/to/URI",
   "multipart":false,
   "payload": "payload",
   "responseStatusCode":200,
   "responseContentType": "JSON" ,
   "preThinkTime": 1000,
   "postThinkTime":2000,
   "execWeight": "Sparse"
}
`

const testCaseForDeletion = `
{
   "testname":"TestCaseService2",
   "description":"desc",
   "overrideHost":"host",
   "overridePort":"9191",
   "httpMethod":"GET",
   "baseUri": "path/to/URI",
   "multipart":false,
   "payload": "payload",
   "responseStatusCode":200,
   "responseContentType": "JSON" ,
   "preThinkTime": 1000,
   "postThinkTime":2000,
   "execWeight": "Sparse",
   "headers":[{
   	 "key": "Authorization",
     "value" :"Header-Value"
   }],
  "responseValues":[{
     "value":"Res-Value",
     "extractionKey": "Res-Key"
  }],
  "multipartPayload":[{
     "fieldName": "F-Name",
   	 "fieldValue":"PayloadName",
     "fileName": "file-name"
  }]
}
`

func TestValidTestCasePost(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validTestCase)
	r.HandleFunc("/test-cases", postTestCase)

	os.Remove(os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/TestCaseService.xml")

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test"
	request, err := http.NewRequest(http.MethodPost, "/test-cases", reader)
	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, w.Code, "Error: Did Not Successfully Post")
}

func TestCasePostWithExistingFileName(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validTestCase)
	r.HandleFunc("/test-cases", postTestCase)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test"
	request, err := http.NewRequest(http.MethodPost, "/test-cases", reader)
	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should not have Successfully posted")
}

func TestCasePostNoHeader(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validTestCase)
	r.HandleFunc("/test-cases", postTestCase)

	filePath := ""
	request, err := http.NewRequest(http.MethodPost, "/test-cases", reader)
	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should not have Successfully posted")
}

func TestPostTestCaseMissingRequiredValues(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(TestCaseMissingRequired)
	r.HandleFunc("/test-cases", postTestCase)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test"
	request, err := http.NewRequest(http.MethodPost, "/test-cases", reader)
	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should not have Successfully posted")
}

func TestValidTestCasePut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())
	reader := strings.NewReader(validTestCase)

	//r.HandleFunc("/test-cases", putTestCase)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPut, "/test-cases/TestCaseService", reader)

	assert.NoError(t, err)

	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusNoContent, w.Code, "Did Not successfully Update")
}

func TestTestCasePutMissingRequired(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(TestCaseMissingRequired)
	r.HandleFunc("/test-cases", putTestCase)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPut, "/test-cases/TestCaseService", reader)
	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should not have successfully updated")
}

func TestInvalidUrlTestCasePut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validTestCase)
	r.HandleFunc("/test-suites", putTestCase)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPut, "/test-cases/xxxxxxxxxxxzzx", reader)
	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusNotFound, w.Code, "Sucessfully updated. Should not have updated")
}

func TestNoUrlTestCasePut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validTestCase)
	r.HandleFunc("/test-cases", putTestCase)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPut, "", reader)
	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusNotFound, w.Code, "Sucessfully updated. Should not have updated")
}

func TestCasePutWithNoPathSlash(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validTestCase)
	r.HandleFunc("/test-cases", putTestCase)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test"
	request, err := http.NewRequest(http.MethodPut, "/test-cases/TestCaseService", reader)
	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusNoContent, w.Code, "Did not update")
}

func TestNoPathTestCasePut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(validTestCase)
	r.HandleFunc("/test-cases", putTestCase)

	filePath := ""
	request, err := http.NewRequest(http.MethodPut, "/test-cases/TestCaseService", reader)
	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code, "Successfully updated. Should not have worked due to no filepath")
}

func TestNoNameTestCasePut(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(TestCaseNoName)
	r.HandleFunc("/test-cases", putTestCase)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test"
	request, err := http.NewRequest(http.MethodPut, "/test-cases/TestCaseService", reader)
	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code, "Successfully updated. Should not have worked due to no filepath")
}

func TestSuccessfulGetAllCases(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	directoryPath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodGet, "/test-cases", nil)
	assert.NoError(t, err)

	request.Header.Set("path", directoryPath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code, "Did not get all test casess")
}

func TestGetAllCasesNoHeader(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	request, err := http.NewRequest(http.MethodGet, "/test-cases", nil)

	request.Header.Set("path", "")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should not have retrieved all test cases")
}

func TestGetTestCaseNoHeader(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	filePath := ""
	request, err := http.NewRequest(http.MethodGet, "/test-cases/Case1", nil)

	request.Header.Set("path", filePath)
	request.Header.Get("path")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should not return data")
}

func TestGetTestCaseFileNotFound(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/cases"
	request, err := http.NewRequest(http.MethodGet, "/test-cases/xxx", nil)

	request.Header.Set("path", filePath)
	request.Header.Get("path")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusNotFound, w.Code, "Should not return data")
}

func TestDeleteAllCasesSuccess(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	directory := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/cases"
	err := ioutil.WriteFile(fmt.Sprintf("%s%s.xml", directory, "test"), nil, 0666)
	if err != nil {
		logrus.Errorf("Error trying to create a file: %s", err)
	}

	request, err := http.NewRequest(http.MethodDelete, "/test-cases", nil)
	assert.NoError(t, err)

	request.Header.Set("path", directory)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code, "All files have been DELETED")

}

func TestDeleteAllCasesNoHeader(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	DirectoryPath := ""
	request, err := http.NewRequest(http.MethodDelete, "/test-cases", nil)
	if err != nil {
		logrus.Warnf("Error creating the request %s", err)
	}

	request.Header.Set("path", DirectoryPath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Did not DELETE the files")
}

func TestDeleteAllCasesEmptyDirectory(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	DirectoryPath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/cases/"
	request, err := http.NewRequest(http.MethodDelete, "/test-cases/", nil)
	if err != nil {
		logrus.Warnf("Error creating the request %s", err)
		return
	}

	request.Header.Set("path", DirectoryPath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusNotFound, w.Code, "Empty Directory")
}

func TestSuccessfulCaseDelete(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	reader := strings.NewReader(testCaseForDeletion)
	r.HandleFunc("/test-cases", postTestCase)

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodPost, "/test-cases", reader)
	assert.NoError(t, err)
	request.Header.Set("path", filePath)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, w.Code, "Error: Did Not Successfully Post")

	r.HandleFunc("/test-cases", deleteTestCase)

	filePath = os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err = http.NewRequest(http.MethodDelete, "/test-cases/TestCaseService2", nil)
	assert.NoError(t, err)
	request.Header.Set("path", filePath)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusNoContent, w.Code, "Error. Did not successfully Delete")
}

func TestDeleteCaseFileNotFound(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	filePath := os.Getenv("GOPATH") + "/src/github.com/xtracdev/automated-perf-test/uiServices/test/"
	request, err := http.NewRequest(http.MethodDelete, "/test-cases/xxx", nil)

	request.Header.Set("path", filePath)
	request.Header.Get("path")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusNotFound, w.Code, "Should not have successfully deleted")
}

func TestDeleteCaseWithNoHeader(t *testing.T) {
	r := chi.NewRouter()
	r.Mount("/", GetIndexPage())

	filePath := ""
	request, err := http.NewRequest(http.MethodDelete, "/test-cases/TestCaseService", nil)

	request.Header.Set("path", filePath)
	request.Header.Get("path")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should not have successfully deleted")
}
