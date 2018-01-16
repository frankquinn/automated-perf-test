import { Injectable } from "@angular/core";
import { HttpModule } from "@angular/http";
import { Http } from "@angular/http";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { ConfigurationsComponent } from "./configurations/configurations.component";
import { Headers } from "@angular/http/src/headers";

const httpOptions = {
  headers: new HttpHeaders({
    "Content-Type": "application/json",
    configPathDir:
      //remove
      "C:/Users/a615194/go/src/github.com/xtracdev/automated-perf-test/config/"
  })
};

@Injectable()
export class AutomatedUIServices {
  constructor(private http: HttpClient) {}
  //remove
  private url = "http://localhost:9191/configs";

  createJsonFile(configData): void {
    console.log("Form", configData);
    //ngbootstrap alert note
    this.http.post(this.url, configData, httpOptions).subscribe(data => {
      console.log(data);
    });
  }
}
