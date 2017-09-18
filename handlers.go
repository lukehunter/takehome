package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
    "net/url"
    "encoding/json"
)

var t, _ = template.ParseFiles("./templates/index.html")

func Index(w http.ResponseWriter, r *http.Request) {
    error := r.URL.Query().Get("error")

    if len(error) != 0 {
        t.Execute(w, "Error: " + error)
    } else {
        t.Execute(w, "")
    }
}

func FormScan(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    target := r.Form["target"][0]

    scan, error := DoScan(target)

    if error != nil {
        ErrorRedirect(w, r, error.Error())
        return
    }

    scans, err := GetScans(scan.Ip)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintln(w, "host: ", target)

    if !scan.Responding {
        fmt.Fprintln(w, "Host is not responding. Showing results of earlier scans, if any.")
    }

    if len(scans) == 0 {
        fmt.Fprintln(w, "No previous scans")
    }

    if len(scans) > 1 {

    }

    for _, scan := range scans {

    }
}

func PrintDiff(w http.ResponseWriter, scan0, scan1 Scan) {
    m := make(map[int]int)
    opened := []int{}
    closed := []int{}

    for _, port := range scan0.OpenPorts {
        m[port]++
    }

    for _, port := range scan1.OpenPorts {
        m[port]--
    }

    for k, v := range m {
        if v > 0 {
            opened = append(opened, k)
        }
        if v < 0 {
            closed = append(closed, k)
        }
    }
}

func PrintScan(w http.ResponseWriter, scan Scan) {
    fmt.Fprintln(w, scan.Ip, scan.DatePerformed, scan.RawPorts)
}

func ErrorRedirect(w http.ResponseWriter, r *http.Request, error string) {
    http.Redirect(w, r, fmt.Sprintf("/?error=%q", url.QueryEscape(error)) , 302)
}

func WsScan(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    target := vars["target"]

    scans, err := GetScans(target)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = json.NewEncoder(w).Encode(scans)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}