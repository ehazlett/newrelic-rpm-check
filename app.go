package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "flag"
    "log"
    _"io/ioutil"
)

var (
    apiKey string
    appId   string
    threshold    int
    quiet       bool
)

type (
    Application struct {
        Info    ApplicationInfo `json:"application"`
    }

    ApplicationInfo struct {
        Name    string `json:"name"`
        Summary Summary `json:"application_summary"`
    }

    Summary struct {
        ResponseTime    float64 `json:"response_time"`
        Throughput      float64 `json:"throughput"`
    }
)

func init() {
    flag.StringVar(&apiKey, "k", "", "NewRelic API Key")
    flag.StringVar(&appId, "a", "", "NewRelic Application ID")
    flag.IntVar(&threshold, "t", 0, "Threshold for throughput (min required)")
    flag.BoolVar(&quiet, "q", false, "Only show status (STATUS:VALUE)")
}

func check(key string, id string) {
    client := &http.Client{}
    url := fmt.Sprintf("https://api.newrelic.com/v2/applications/%s.json", id)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Printf("Error sending request to New Relic: %s", err)
        return
    }
    req.Header.Add("X-Api-Key", key)
    resp, err := client.Do(req)
    defer resp.Body.Close()
    //content, _ := ioutil.ReadAll(resp.Body)
    //log.Println(string(content))
    var app Application
    d := json.NewDecoder(resp.Body)
    if err := d.Decode(&app); err != nil {
        log.Printf("Error parsing JSON from New Relic: %s", err)
        return
    }
    val := int(app.Info.Summary.Throughput)
    if val <= threshold {
        if quiet {
            log.Fatalf("CRITICAL:%d", val)
            return
        }
        log.Fatalf("CRITICAL: %s Throughput: %d rpm", app.Info.Name, val)
    }
    if quiet {
        log.Printf("OK:%d", int(app.Info.Summary.Throughput))
        return
    }
    log.Printf("OK: %s Throughput: %d rpm", app.Info.Name, int(app.Info.Summary.Throughput))
}

func main() {
    log.SetFlags(0)
    flag.Parse()
    check(apiKey, appId)
}