package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "time"
)

func main() {
    port := portcheck()

    fmt.Fprintf(os.Stdout, "Start listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        config := readconf()
        zeit := time.Now().Format("2006/01/02 15:04:05")
        fmt.Fprintf(os.Stdout, "%s stdout for logging: %s\n", zeit, config)
        fmt.Fprintf(w, "Hostname: \n%s\n", hostname)
        fmt.Fprintf(w, "Port: %s\n\n", port)
        fmt.Fprintf(w, "Configfileinput: \n%s", config)
    })
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

func readconf() (string) {
    // read configfile
    b, err := ioutil.ReadFile("config.file")
    if err != nil {
        fmt.Print(err)
    }
    str := string(b)
    return str
}
