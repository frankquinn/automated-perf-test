Feature: Test Case Scenarios
  As an API user
  I want to be able use various requests for test cases
  So that I can test my application using custom metrics


                                ###################################
                                #######    POST REQUESTS    #######
                                ###################################

  Scenario: Successful creation of Test Case
    Given there is no existing test file "TestCaseSAMPLE.xml"
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test"
    When I send "POST" request to "/test-cases" with a body:
    """
      {
       "testname":"TestCaseSAMPLE",
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
    """
    Then the response code should be 201
    And the response body should be empty

  Scenario: Unsuccessful creation of Test Case (file already exists )
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test"
    When I send "POST" request to "/test-cases" with a body:
    """
      {
       "testname":"TestCaseSAMPLE",
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
    """
    Then the response code should be 201
    And the header "path" is "/uiServices/test/"
    When I send "POST" request to "/test-cases" with a body:
    """
      {
       "testname":"TestCaseSAMPLE",
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
    """
    Then the response code should be 400


  Scenario: Unsuccessful creation of test Case ( Missing Required Fields )
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/"
    When I send "POST" request to "/test-suites" with a body:
       """
      {
       "testname":"TestCaseSAMPLE",
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
    """
    Then the response code should be 400


  Scenario: Unsuccessful creation of Test Case ( No header defined )
    Given the automated performance ui server is available
    And the header "path" is ""
    When I send "POST" request to "/test-cases" with a body:
       """
      {
       "testname":"TestCaseSAMPLE",
       "description":"desc",
       "overrideHost":"host",
       "overridePort":"9191",
       "HttpMethod":"GET",
       "baseUri": "path/to/URI",
       "multipart":false,
       "payload": "payload",
       "responseStatusCode":200,
       "responseContentType": "JSON" ,
       "preThinkTime": 1000,
       "postThinkTime":2000,
       "execWeight": "Sparse",
       "headers":[{
   	     "Key": "Authorization",
         "Value" :"Header-Value"
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
    """
    Then the response code should be 400



                                ###################################
                                #######    PUT REQUESTS    #######
                                ###################################

  Scenario: Unsuccessful update of test-case file with PUT request (No File Path)
    Given the file "TestCaseSAMPLE.xml" exists at "/uiServices/test/"
    Given the automated performance ui server is available
    And the header "path" is ""
    When I send "PUT" request to "/test-suites/GodogTestSuite" with body:
  """
     {
   "testname":"TestCaseSAMPLE",
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
    """
    Then the response code should be 400


  Scenario: Unsuccessful update of test-case file with PUT request (Incorrect File Name)
    Given the file "TestCaseSAMPLE.xml" exists at "/uiServices/test/"
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/"
    When I send "PUT" request to "/test-cases/xxx" with body:
    """
     {
   "testname":"TestCaseSAMPLE",
   "description":"desc",
   "overrideHost":"host",
   "overridePort":"9191",
   "HttpMethod":"GET",
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
    """
    Then the response code should be 404


  Scenario: Unsuccessful update of test-case file with PUT request (No File Name)
    Given the file "TestCaseSAMPLE.xml" exists at "/uiServices/test/"
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/"
    When I send "PUT" request to "/test-suites/" with body:
    """
     {
   "testname":"TestCaseSAMPLE",
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
    """
    Then the response code should be 404

  Scenario: Unsuccessful update of test-suite file with PUT request (Missing Required Fields)
    Given the file "TestCaseSAMPLE.xml" exists at "/uiServices/test/"
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/"
    When I send "PUT" request to "/test-suites/TestCaseSAMPLE" with body:
    """
     {
   "testname":"TestCaseSAMPLE",
   "description":"",
   "overridePort":"",
   "httpMethod":"",
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
    """
    Then the response code should be 400


  Scenario: Successful update of test-case file with PUT request
    Given the file "TestCaseSAMPLE.xml" exists at "/uiServices/test/"
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/"
    When I send "PUT" request to "/test-cases/TestCaseSAMPLE" with body:
     """
     {
   "testname":"TestCaseSAMPLE",
   "description":"desc",
   "overrideHost":"host",
   "overridePort":"1001",
   "httpMethod":"POST",
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
    """
    Then the response code should be 204
    And the response body should be empty


                                ###################################
                                #######    GET ALL REQUESTS #######
                                ###################################
  Scenario: Successful retrieval all test cases with valid "GET" request
    ##Add additional file first so there are multiple files to GET
    Given there is no existing test file "TestCaseSAMPLE2.xml"
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/"
    When I send "POST" request to "/test-cases" with a body:
    """
      {
       "testname":"TestCaseSAMPLE2",
       "description":"desc2",
       "overrideHost":"host",
       "overridePort":"9191",
       "httpMethod":"PUT",
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
    """
    Then the response code should be 201
    And the response body should be empty
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/"
    When I send a "GET" request to "/test-cases"
    Then the response code should be 200
    And the test case collection response body should match json:
      """
        [
          {
          "name": "TestCaseSAMPLE,
          "description": "Case Desc",
          "httpMethod": "GET"
          },
          {
          "name": "TestCaseSAMPLE2,
          "description": "Case Desc2",
          "httpMethod": "PUT"
          }
        ]
    """


  Scenario: Unsuccessful retrieval of test-cases (No Header)
    Given the automated performance ui server is available
    And the header "path" is ""
    When I send a "GET" request to "/test-cases"
    Then the response code should be 400

                                ###################################
                                #######    GET REQUESTS #######
                                ###################################

  Scenario: Unsuccessful retrieval of test-Cases file (File Not Found)
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/"
    When I send a "GET" request to "/test-cases/xxx"
    Then the response code should be 404


  Scenario: Unsuccessful retrieval of test-suites file (No Header)
    Given the automated performance ui server is available
    And the header "path" is ""
    When I send a "GET" request to "/test-suites/TestCaseSAMPLE"
    Then the response code should be 400

  Scenario: Retrieve Test Case file with valid "GET" request
    Given there is no existing test file "TestCaseSAMPLE.xml"
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test"
    When I send "POST" request to "/test-cases" with a body:
    """
      {
       "testname":"TestCaseSAMPLE",
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
    """
    Then the response code should be 201
    Given the file "TestCaseSAMPLE.xml" exists at "/uiServices/test/samples/"
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/samples/"
    And the file name is "TestCaseSAMPLE.xml"
    When I send a "GET" request to "/test-cases/testCase"
    Then the response code should be 200
    And the test case response body should match json:
    """
   {
      "TestName": TestCaseSAMPLE",
      "OverrideHost": "host",
      "OverridePort": "9191",
      "HTTPMethod": "GET",
      "Description": "desc",
      "baseUri": "",
      "Multipart": false,
      "Payload": "",
      "MultipartPayload": null,
      "ResponseStatusCode": 0,
      "ResponseContentType": "",
      "headers": null,
      "responseValues": null,
      "PreThinkTime": 0,
      "PostThinkTime": 0,
      "ExecWeight": ""
    }
    """

                                ###################################
                                #######    DELETE REQUESTS ########
                                ###################################



  Scenario:  Fail to remove Test Case file with "DELETE" request (File Not Found)
    Given the file "TestCaseSAMPLE.xml" exists at "/uiServices/test/"
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/"
    When I send a "DELETE" request to "/test-cases/xxxx"
    Then the response code should be 404

  Scenario:  Fail to remove Test Case file with "DELETE" request (No Header)
    Given the file "TestCaseSAMPLE.xml" exists at "/uiServices/test/"
    Given the automated performance ui server is available
    And the header "path" is ""
    When I send a "DELETE" request to "/test-cases/xxxx"
    Then the response code should be 400

  Scenario:  Successful removal Test Case file with "DELETE" request
    #create file to delete
    Given there is no existing test file "TestCaseSAMPLE.xml"
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/"
    When I send "POST" request to "/test-cases" with a body:
    """
      {
       "testname":"TestCaseSAMPLE",
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
    """
    Then the response code should be 201
    #Delete
    Given the file "TestCaseSAMPLE.xml" exists at "/uiServices/test/"
    Given the automated performance ui server is available
    And the header "path" is "/uiServices/test/"
    When I send a "DELETE" request to "/test-cases/TestCaseSAMPLE"
    Then the response code should be 204


Scenario: Unsuccessful deleting of test-case (No Header)
  Given the automated performance ui server is available
  And the header "path" is ""
  When I send a "DELETE" request to "/test-cases"
  Then the response code should be 400


Scenario: Unsuccessful deleting of test-case file (Empty Directory)
  Given the automated performance ui server is available
  And the header "path" is "/uiServices/test/cases"
  When I send a "DELETE" request to "/test-cases/"
  Then the response code should be 404


Scenario: Successful deleting of test-case file with DELETE request
  Given the file "deleteTest.xml" exists at "/uiServices/test/cases"
  Given the automated performance ui server is available
  And the header "path" is "/uiServices/test/cases"
  When I send a "DELETE" request to "/test-cases"
  Then the response code should be 200