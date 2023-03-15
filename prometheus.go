// Copyright (c) 2017, Tax Products Group, LLC.
//
// All rights reserved.

package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

/**
 * Global Variables
 */

// Prometheus Metrics  https://prometheus.io/docs/concepts/metric_types/
// https://prometheus.io/docs/practices/naming/
var (
	// A counter is a cumulative metric that represents a single numerical
	// value that only ever goes up. A counter is typically used to count
	// requests served, tasks completed, errors occurred, etc.
	metricEmails = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "sbtpg",
		Subsystem: "collection_machine",
		Name:      "emails_sent_total",
		Help:      "The total number of emails sent.",
	})
	metricErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "sbtpg",
		Subsystem: "collection_machine",
		Name:      "email_errors_total",
		Help:      "The total number of email errors.",
	})
	// A gauge is a metric that represents a single numerical value that can
	// arbitrarily go up and down.
	metricDbConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "sbtpg",
		Subsystem: "collection_machine",
		Name:      "database_connections",
		Help:      "The number of database connections.",
	})

	// A histogram samples observations (usually things like request durations
	// or response sizes) and counts them in configurable buckets. It also
	// provides a sum of all observed values.

	// Similar to a histogram, a summary samples observations (usually things
	// like request durations and response sizes). While it also provides a
	// total count of observations and a sum of all observed values, it
	// calculates configurable quantiles over a sliding time window.
)

/**
 * Functions
 */

// prometheusInit initailizes Prometheus stats
func prometheusInit() {
	prometheus.MustRegister(metricEmails)
	prometheus.MustRegister(metricErrors)
	prometheus.MustRegister(metricDbConnections)
}
