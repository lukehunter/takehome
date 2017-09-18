package main

import (
    "os/exec"
    "fmt"
    "encoding/xml"
    "bytes"
    "strings"
    "time"
)

type xnDoc struct {
    XMLName xml.Name `xml:"nmaprun"`
    Hosts []xnHost `xml:"host"`
    RunStats xnStats `xml:"runstats"`
}

type xnStats struct {
    HostStats xnHostStats `xml:"hosts"`
}

type xnHostStats struct {
    Up bool `xml:"up,attr"`
}
type xnHost struct {
    Addrs []xnAddr `xml:"address"`
    Ports xnPorts `xml:"ports"`
}

type xnAddr struct {
    Addr string `xml:"addr,attr"`
}

type xnPorts struct {
    Port []xnPort `xml:"port"`
}

type xnPort struct {
    PortNum int `xml:"portid,attr"`
}

func nmapScan(host string, startPort, endPort int) (Scan, error) {
    scan := Scan{StartPort: startPort, EndPort: endPort}

    if endPort < startPort || startPort < 1 || endPort < 1 || startPort > 65535 || endPort > 65535 {
        return scan, fmt.Errorf("Invalid port args: start %q end %q", startPort, endPort)
    }

    // todo read nmap path from config
    ports := fmt.Sprintf("%v-%v", startPort, endPort)
    c := exec.Command("nmap", "-oX", "-", host, "-p", ports, "--open")
    var stderr bytes.Buffer
    c.Stderr = &stderr
    //err := cmd.Run()
    //if err != nil {
    //    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    //    return
    //}
    cout, cerr := c.Output()
    if cerr != nil {
        fmt.Println(stderr.String())
        return scan, cerr
    }

    doc := xnDoc{}

    dec := xml.NewDecoder(bytes.NewReader(cout))
    derr := dec.Decode(&doc)
    if derr != nil {
        return scan, derr
    }

    scan.Responding = doc.RunStats.HostStats.Up

    if !scan.Responding {
        return scan, nil
    }

    if len(doc.Hosts) != 1 {
        return scan, fmt.Errorf("Expected 1 host, found %q", len(doc.Hosts))
    }

    if len(doc.Hosts[0].Addrs) < 1 {
        return scan, fmt.Errorf("Found no addresses")
    }

    scan.Ip = doc.Hosts[0].Addrs[0].Addr

    for _, port := range doc.Hosts[0].Ports.Port {
        scan.OpenPorts = append(scan.OpenPorts, port.PortNum)
    }

    scan.RawPorts = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(scan.OpenPorts)), ","), "[]")
    scan.DatePerformed = time.Now().Format("2006-01-02 15:04:05")

    return scan, nil
}

