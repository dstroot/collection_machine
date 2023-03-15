// Copyright (c) 2017, Tax Products Group, LLC.
//
// All rights reserved.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/dstroot/utility"
)

var (
	report     Metrics            // global metrics
	start      = time.Now().UTC() // global start
	buildstamp = "not set"
	githash    = "not set"
)

// Metrics holds our metrics for reporting
type Metrics struct {
	Program           string
	Buildstamp        string
	GitHash           string
	GoVersion         string
	RunTime           string
	Activities        int
	EmailsSent        int
	SMSSent           int
	PaymentsCreated   int
	DirectMailCreated int
	DBconnections     int
	GiactError        int
}

// status writes a JSON object with the current metrics
func status(w http.ResponseWriter, req *http.Request) {
	report.RunTime = fmt.Sprintf("%v", utility.RoundDuration(time.Since(start), time.Second))
	report.DBconnections = db.Stats().OpenConnections

	js, err := json.MarshalIndent(report, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
