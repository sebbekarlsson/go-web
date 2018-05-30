package main


import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "./stringutils"
)

func route(w http.ResponseWriter, r *http.Request) {
    template := "index.html"

    if r.URL.Path != "/" {
        template = stringutils.RmFirstChar(r.URL.Path)
    }

    s, err := ioutil.ReadFile(template)

    response := string(s)

    if err != nil {
        response = err.Error()
    }

    fmt.Fprintf(w, response)
}

func main() {
    http.HandleFunc("/", route)
    log.Fatal(http.ListenAndServe(":5000", nil))
}
