import { Injectable } from "@angular/core";
import { HttpModule } from "@angular/http";
import { Http } from "@angular/http";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { ConfigurationsComponent } from "./configurations/configurations.component";
import { Headers } from "@angular/http/src/headers";
import { Observable } from "rxjs/Observable";

const httpOptions = {
  headers: new HttpHeaders({
    "Content-Type": "application/json",
    configPathDir:
      "C:/Users/a586754/go/src/github.com/xtracdev/automated-perf-test/config/"
  })
};

@Injectable()
export class AutomatedUIServices {
  constructor(private http: HttpClient) {}

   private url = "http://localhost:9191/configs";

  postConfig$(configData): Observable<any> {
    return this.http.post(this.url, configData, httpOptions);
  }
}