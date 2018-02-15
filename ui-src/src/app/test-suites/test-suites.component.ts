import { Component, OnInit } from "@angular/core";
import { AutomatedUIServices } from "../automated-ui-services";
import { ToastsManager } from "ng2-toastr/ng2-toastr";
@Component({
  selector: "app-test-suites",
  templateUrl: "./test-suites.component.html",
  styleUrls: ["./test-suites.component.css"]
})
export class TestSuitesComponent {
  testSuiteData = {};
  testSuitePath = undefined;
  testSuiteFileName = undefined;
  constructor(
    private automatedUIServices: AutomatedUIServices,
    private toastr: ToastsManager
  ) { }
  testSchema = {
    type: "object",
    properties: {
      name: { type: "string" },
      testStrategy: {
        type: "string",
        "enum": [
          "SuiteBased",
          "ServiceBased"
        ]
      },
      description: { type: "string" }

    }
  }

  onAdd() {
    this.testSuiteData = undefined;
    //clear schema and get all in available  test cases from testSuitePath

  }

  onDelete() {
    //need delete service
  }

  onCancel() {
    //clear schema and moving info back into available (get method)
    this.automatedUIServices
      .getConfig$(this.testSuitePath, this.testSuiteFileName)
      .subscribe(
      data => {
        this.testSuiteData = data;
        this.toastr.success("Previous data reloaded!");
      },
      error => {
        this.testSuiteData = undefined;
      }
      );
  }

  onSave(testSuiteData) {
    this.automatedUIServices.postTestSuite$(testSuiteData, this.testSuitePath).subscribe(
      data => {
        this.toastr.success("Your data has been saved!", "Success!");
      },

      error => {
        switch (error.status) {
          case 500: {
            this.toastr.error("An error has occurred!", "Check the logs!");
            break;
          }
          case 400: {
            this.toastr.error(
              "Some of the fields do not conform to the schema!",
              "An error occurred!"
            );
            break;
          }
          default: {
            this.toastr.error("Your data did not save!", "An error occurred!");
          }
        }
      }
    );
  }

}
