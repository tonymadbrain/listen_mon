package main

import (
  "log"
  "os/exec"
  "strings"
  "net/http"
  "io"
  "encoding/json"
  "strconv"
)

func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
  })
}

func main() {

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    cmd_str := "netstat -an -ptcp | grep tcp4 | grep LISTEN | awk '{print $4}'"
    cmd := exec.Command("bash", "-c", cmd_str)

    netstat, err := cmd.CombinedOutput()
    if err != nil {
      log.Fatal(err)
    }

    lines := strings.Split(string(netstat), "\n")

    results := []map[string]string{}

    for _, line := range lines {
      parsed := strings.Split(line, ".")

      if len(parsed) == 5 {

        app_port, err := strconv.Atoi(parsed[4])
        if err != nil {
          log.Fatal(err)
        }

        if 1000 < app_port && app_port < 9999 {

          cmd_str = "lsof -n -i:" + parsed[4] + " | grep LISTEN | awk '{print $1}'"
          cmd = exec.Command("bash", "-c", cmd_str)

          lsof, err := cmd.CombinedOutput()
          if err != nil {
            log.Fatal(err)
          }

          lsof_trim := strings.TrimSpace(string(lsof))

          app := "No app found"

          if lsof_trim != "" {
            app = lsof_trim
          }

          results = append(results, map[string]string{parsed[4]:app})
        }
      }
    }
    response, err := json.Marshal(results)
    if err != nil {
      log.Fatal(err)
    }
    io.WriteString(w, string(response))
  })

  port := ":1234"
  log.Printf("Listen on http://127.0.0.1%s/", port)
  log.Fatal(http.ListenAndServe(port, Log(http.DefaultServeMux)))
}
