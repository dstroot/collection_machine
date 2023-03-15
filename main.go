// Copyright (c) 2017, Tax Products Group, LLC.
//
// All rights reserved.

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	u "github.com/dstroot/utility"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	db *sql.DB // global database to share
)

func main() {
	err := initialize()
	u.Check(err)

	err1 := setupDatabase()
	u.Check(err1)
	defer db.Close()

	prometheusInit()

	// Report metrics (JSON via http) via separate goroutine
	go func() {
		metricDbConnections.Set(float64(db.Stats().OpenConnections))
		http.Handle("/metrics", prometheus.Handler())
		http.Handle("/", http.HandlerFunc(status))
		err := http.ListenAndServe(":"+cfg.Port, nil)
		u.Check(err)
	}()

	// SIGINT or SIGTERM handling
	sigs := make(chan os.Signal)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		if cfg.Debug {
			fmt.Println(sig)
		}
		done <- true
	}()

	// for closing goroutine down cleanly
	closer := make(chan struct{})
	wg := new(sync.WaitGroup)

	// Process goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// Listen for close message
			select {
			case _ = <-closer:
				return

			default:
			}

			// main processing
			err := processCollections()
			u.Check(err)

			// throttle loops
			time.Sleep(cfg.LoopTime)
		}
	}()

	// Wait until we get the quit message
	<-done

	// Shutdown goroutine
	log.Println("Received quit. Shutting down...")

	// close main goroutine
	close(closer)
	wg.Wait()

	// Log running duration
	log.Printf("Running for: %s\n", u.RoundDuration(time.Since(start), time.Second))
	log.Println("Goodbye!")
}
