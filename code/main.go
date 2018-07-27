package main

import (
  "errors"
  "fmt"
  "github.com/prometheus/client_golang/prometheus"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "time"
)

var (
  curls = prometheus.NewCounter(prometheus.CounterOpts{
    Name:   "webtest_connections_total",
    Help:   "Number of checks",
    ConstLabels: prometheus.Labels{"type": "http"},
  })
)

func init() {
  prometheus.MustRegister(curls)
}

func main() {
    port := portcheck()
    fmt.Fprintf(os.Stdout, "Start listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        config, _ := readconf()
        timestamp := time.Now().Format("02/01/2006 15:04:05")
        fmt.Fprintf(os.Stdout, "%s stdout for logging: %s\n", timestamp, config)
        fmt.Fprintf(w, "Hostname: \n%s\n", hostname)
        fmt.Fprintf(w, "Port: %s\n\n", port)
        fmt.Fprintf(w, "Configfileinput: \n%s", config)
        curls.Inc()
    })
    metrics()
    log.Fatal(http.ListenAndServe(":" + port, nil))
}

func portcheck() (string) {
    // check env vars for specific port
    port := os.Getenv("PORT")
    if port == "" {
       port = "8000"
    }
    return port
}

func readconf() (string, error) {
    // read configfile
    b, err := ioutil.ReadFile("config.file")
    if err != nil {
        return "OH NO", errors.New("file not available.")
    } else {
        str := string(b)
        return str, nil
    }
}

func metrics() {
    // expose metrics @ /metrics
    http.Handle("/metrics", prometheus.Handler())
}
