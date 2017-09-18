package main

import (
    "regexp"
    "errors"
)

type Scan struct {
    Ip            string
    DatePerformed string
    StartPort     int
    EndPort       int
    OpenPorts     []int
    RawPorts      string
    Responding    bool
}

type Scans []Scan

var hostnameRegex = regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`)
var ipv4Regex = regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)

func ValidTarget(target string) bool {
    return ipv4Regex.MatchString(target) || (hostnameRegex.MatchString(target) && len(target) <= 255)
}

func DoScan(target string) (Scan, error) {
    scan := Scan{}

    // check input valid (essential to prevent command injection when running nmap)
    if !ValidTarget(target) {
        return scan, errors.New("Invalid hostname/ip")
    }

    // call nmap
    scan, err := nmapScan(target, 1, 1000)

    if err != nil {
        return scan, err
    }

    // don't save anything to db if host is not responding
    if !scan.Responding {
        return scan, nil
    }

    // save to db
    err = SaveScan(scan)

    return scan, err
}