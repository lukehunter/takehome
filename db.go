package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
    "strings"
    "strconv"
    "os"
)

var db *sql.DB

func InitDB() {
    var err error

    db, err = sql.Open("mysql", os.Getenv("TAKEHOME_DB_CONN"))
    if err != nil {
        log.Panic(err)
    }

    if err = db.Ping(); err != nil {
        log.Panic(err)
    }
}

func GetScans(ip string) (Scans, error) {
    if !ValidTarget(ip) {
        return nil, fmt.Errorf("Invalid target %v", ip)
    }

    query := fmt.Sprintf("SELECT ip, date_performed, start_port, end_port, open_ports FROM " +
        "scan JOIN `host` ON `host`.host_id=scan.host_id WHERE ip=? ORDER BY date_performed DESC")
    rows, err := db.Query(query, ip)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    scans := Scans{}

    for rows.Next() {
        scan := new(Scan)
        err := rows.Scan(&scan.Ip, &scan.DatePerformed, &scan.StartPort, &scan.EndPort, &scan.RawPorts)

        if err != nil {
            return nil, err
        }

        openPortsArr := strings.Split(scan.RawPorts, ",")

        for _, portStr := range openPortsArr {
            port, err := strconv.Atoi(portStr)
            if err != nil {
                return nil, err
            }
            scan.OpenPorts = append(scan.OpenPorts, port)
        }

        scans = append(scans, *scan)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return scans, nil
}

func SaveScan(scan Scan) error {
    insertHost := fmt.Sprintf("INSERT IGNORE INTO `host` SET ip=?")
    getHost := fmt.Sprintf("SELECT host_id FROM `host` WHERE ip=?")

    _, err := db.Exec(insertHost, scan.Ip)

    if err != nil {
        return err
    }

    hostId := -1
    row := db.QueryRow(getHost, scan.Ip)

    err = row.Scan(&hostId)

    if err != nil {
        return err
    }

    insertScan := fmt.Sprintf("INSERT INTO scan (host_id, date_performed, start_port, end_port, open_ports) " +
                                "VALUES (?, ?, ?, ?, ?)")

    _, err = db.Exec(insertScan, hostId, scan.DatePerformed, scan.StartPort, scan.EndPort, scan.RawPorts)

    return err
}
