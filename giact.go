// Copyright (c) 2017, Tax Products Group, LLC.
//
// All rights reserved.

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gchaincl/dotsql"
	"github.com/pkg/errors"
)

/**
 * Globals
 */

var (
	giactAcct  giactReq
	response   giactRes
	client     = &http.Client{}
	startAPI   time.Time
	elapsedAPI time.Duration
	giactList  = make(map[int]bool)
)

/**
 * Structs
 */

// giactReq is a GIACT request
type giactReq struct {
	UniqueID       string
	Check          Check
	GVerifyEnabled bool
}

// giactRes is a GIACT response
type giactRes struct {
	ItemReferenceID      int       `json:"ItemReferenceID"`
	CreatedDate          time.Time `json:"CreatedDate"`
	ErrorMessage         string    `json:"ErrorMessage"`
	VerificationResponse int       `json:"VerificationResponse"`
	AccountResponseCode  int       `json:"AccountResponseCode"`
	BankName             string    `json:"BankName"`
}

// Check is the account we are validating
type Check struct {
	RoutingNumber string
	AccountNumber string
	CheckAmount   float64
	AccountType   int
}

/**
 * Functions
 */

// getGiactActions reads the SQL database for the list of GIACT codes.
func getGiactActions() (err error) {

	// Load queries
	dot, err := dotsql.LoadFromFile("sql/sql.sql")
	if err != nil {
		return errors.Wrap(err, "could not load sql")
	}

	// Get records
	rows, err := dot.Query(db, "GET_GIACT_CODES")
	if err != nil {
		return errors.Wrap(err, "get giact codes query failed")
	}
	defer rows.Close()

	var (
		giactCode int
		debitFlag bool
	)

	// interate the records
	for rows.Next() {
		err1 := rows.Scan(
			&giactCode,
			&debitFlag,
		)
		if err1 != nil {
			return errors.Wrap(err1, "row could not be scanned")
		}

		// Add row data to map
		giactList[giactCode] = debitFlag

	}
	err2 := rows.Err()
	if err2 != nil {
		return errors.Wrap(err2, "row processing error")
	}
	return nil
}

// verify takes an account, calls the API and populates the GIACT response
func verify(acct *giactReq, process int) (err error) {

	// Create the JSON payload to post
	j, err := json.Marshal(acct)
	if err != nil {
		return errors.Wrap(err, "JSON encode failed")
	}

	if cfg.GiactLogging {
		var prettyJSON bytes.Buffer
		error := json.Indent(&prettyJSON, j, "", "\t")
		if error != nil {
			log.Println("JSON parse error: ", error)
		}
		log.Printf("GIACT JSON Payload: \n%s", string(prettyJSON.Bytes()))
	}

	var auth string

	if process == 10 { // Intuit
		auth = cfg.GiactAuthIntuit
	}

	if process == 30 { // TaxSlayer
		auth = cfg.GiactAuthTaxSlayer
	}

	// Create new HTTP request
	req, err := http.NewRequest("POST", cfg.GiactURL, bytes.NewBuffer(j))
	if err != nil {
		return errors.Wrap(err, "creating a request failed")
	}

	// Add necessary headers to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", auth)

	// Timing the API
	startAPI = time.Now()

	// Send the request via a HTTP client
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "response failed")
	}
	defer resp.Body.Close()

	// If debugging lets see the response and elapsed time
	if cfg.GiactLogging {

		// API elapsed time
		elapsedAPI = time.Since(startAPI)
		log.Printf("GIACT API took %s, Response: %s", elapsedAPI, resp.Status)

		// dump, err := httputil.DumpResponse(resp, true)
		// if err != nil {
		// 	return errors.Wrap(err, "response dump failed")
		// }
		// log.Printf("%q", dump)
	}

	// Return an error if we don't get a 201/Created
	if resp.StatusCode != http.StatusCreated {
		err := errors.New("http response = " + resp.Status)
		return errors.Wrap(err, "GIACT api call failed")
	}

	// Populate the GIACT reponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return errors.Wrap(err, "JSON decode failed")
	}

	return nil
}
