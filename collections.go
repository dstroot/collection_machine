// Copyright (c) 2017, Tax Products Group, LLC.
//
// All rights reserved.

package main

import (
	"log"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	u "github.com/dstroot/utility"
	"github.com/gchaincl/dotsql"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

/**
 * Constants
 */

const (
	customerInitiatedFalse = 0
	achPayment             = 1
	achRetryTrue           = 1
	achRetryFalse          = 0
	paidTrue               = 1
	checkingAcct           = 1
	responsys              = 2 // QUEUE_VENDOR_ID
	sendgrid               = 3 // QUEUE_VENDOR_ID
	giactErr               = 0
)

/**
 * Global Variables
 */

// GIACTStopCode contains the GIACT stop activity step by process
// TODO Should *not* be hardcoded. Move to flag on process definition table?
var (
	GIACTStopCode = map[int]int{
		10: 75,  // process 10 GIACT stop step is 75
		30: 175, // process 30 GIACT stop step is 175
	}
)

/**
 * Structs
 */

// Process defines a process Record
type Process struct {
	processID   int
	processName string
	processDesc string
}

// Action defines a process action record
type Action struct {
	actionID      int
	actionDesc    string
	duration      int64
	autoDebit     bool
	directMail    bool
	debitRetry    bool
	customerRetry bool
	processEnd    bool
	fromName      string
	fromEmail     string
	subject       string
	email         string
	sms           string
}

// getActions reads the SQL database for a list of actions associated with
// a given process ID and loads them into a slice of actions. This way we
// only query the DB once.
func getActions(id int) (actions []Action, err error) {

	// make slice payments of Payment struct, zero length
	actions = make([]Action, 0)
	var a Action // a records will be added to actions slice

	// Load queries
	dot, err := dotsql.LoadFromFile("sql/sql.sql")
	u.Check(err)

	// Get records to process
	rows, err := dot.Query(db, "GET_PROCESS_ACTIONS", id)
	if err != nil {
		return nil, errors.Wrap(err, "get actions query failed")
	}
	defer rows.Close()

	// interate the records and fill the slice
	for rows.Next() {
		// scan the row into struct 'a'
		err1 := rows.Scan(
			&a.actionID,
			&a.actionDesc,
			&a.duration,
			&a.autoDebit,
			&a.directMail,
			&a.debitRetry,
			&a.customerRetry,
			&a.processEnd,
			&a.fromName,
			&a.fromEmail,
			&a.subject,
			&a.email,
			&a.sms)
		if err1 != nil {
			return nil, errors.Wrap(err1, "row could not be scanned")
		}

		// append 'a' to actions slice
		actions = append(actions, a)
	}
	err2 := rows.Err()
	if err2 != nil {
		return nil, errors.Wrap(err2, "row processing error")
	}
	return actions, nil
}

// getProcesses reads the SQL database for the list of processes.
func getProcesses() (processes []Process, err error) {

	// make slice processes of Process struct, zero length
	processes = make([]Process, 0)
	var p Process // p records will be added to processes slice

	// Load queries
	dot, err := dotsql.LoadFromFile("sql/sql.sql")
	u.Check(err)

	// Get records
	rows, err := dot.Query(db, "GET_PROCESSES")
	if err != nil {
		return nil, errors.Wrap(err, "get processes query failed")
	}
	defer rows.Close()

	// interate the records and fill the slice
	for rows.Next() {
		// scan the row into struct 'p'
		err1 := rows.Scan(
			&p.processID,
			&p.processName,
			&p.processDesc)
		if err1 != nil {
			return nil, errors.Wrap(err1, "row could not be scanned")
		}

		// append 'p' to processes slice
		processes = append(processes, p)
	}
	err2 := rows.Err()
	if err2 != nil {
		return nil, errors.Wrap(err2, "row processing error")
	}
	return processes, nil
}

// generateToken generates and a token and a hash of the token
func generateToken() (token, hash string, err error) {

	token = u.GenRandomString(21) // hex string
	test := []byte(token)

	hashedToken, err := bcrypt.GenerateFromPassword(test, bcrypt.DefaultCost)
	if err != nil {
		return "", "", errors.Wrap(err, "token hashing failed")
	}

	hash = string(hashedToken)
	return token, hash, nil
}

// processCollections is the main processing function
func processCollections() error {

	// Load queries
	dot, err := dotsql.LoadFromFile("sql/sql.sql")
	if err != nil {
		return errors.Wrap(err, "load sql failed")
	}

	// Get GIACT codes
	err = getGiactActions()
	if err != nil {
		return errors.Wrap(err, "get giact codes failed")
	}

	// get the processes
	processes, err := getProcesses()
	if err != nil {
		return errors.Wrap(err, "get processes failed")
	}

	if cfg.Debug {
		// log the process steps
		log.Println("================================================")
		log.Println("Processes")
		log.Println("================================================")
		for i := range processes {
			log.Println(strconv.Itoa(int(processes[i].processID)) + " " + processes[i].processName)
		}
	}

	// iterate through each process
	for i := range processes {

		// get the steps of our process
		actions, err := getActions(processes[i].processID)
		if err != nil {
			return errors.Wrap(err, "get actions failed")
		}

		if cfg.Debug {
			// log the process steps
			log.Println("================================================")
			log.Println("Process " + strconv.Itoa(int(processes[i].processID)) + " Actions:")
			log.Println("================================================")
			for i := range actions {
				log.Println(strconv.Itoa(int(actions[i].actionID)) + " " + actions[i].actionDesc[:7] +
					" Dur: " + strconv.Itoa(int(actions[i].duration)) +
					" AD: " + strconv.FormatBool(actions[i].autoDebit) +
					" DM: " + strconv.FormatBool(actions[i].directMail) +
					" DR: " + strconv.FormatBool(actions[i].debitRetry) +
					" CR: " + strconv.FormatBool(actions[i].customerRetry) +
					" END: " + strconv.FormatBool(actions[i].processEnd))
			}
		}

		// Get unfunded records to process by process id
		rows, err := dot.Query(db, "GET_COLLECTION_ITEMS", processes[i].processID)
		if err != nil {
			return errors.Wrap(err, "get collection items failed")
		}
		defer rows.Close()

		var (
			productID     int
			firstName     string
			middleInitial string
			lastName      string
			address1      string
			address2      string
			city          string
			state         string
			postalCode    string
			email         string
			phone         string
			amount        float64
			paymentURL    string
			rtn           string
			dan           string
			process       int
			activity      int
			timestamp     time.Time
		)

		// iterate the rows
		for rows.Next() {

			// scan the row into vars
			err1 := rows.Scan(
				&productID,
				&firstName,
				&middleInitial,
				&lastName,
				&address1,
				&address2,
				&city,
				&state,
				&postalCode,
				&email,
				&phone,
				&amount,
				&paymentURL,
				&rtn,
				&dan,
				&process,
				&activity,
				&timestamp)
			if err1 != nil {
				return errors.Wrap(err1, "row scan failed")
			}

			// Get the actions slice index of the current activity
			indexPosition, err2 := u.SliceIndex(len(actions), func(i int) bool { return actions[i].actionID == activity })
			if err2 != nil {
				return errors.Wrap(err2, "get slice index failed")
			}

			// Determine the next activity
			nextActivity := indexPosition + 1

			// Reset GIACT flag
			giactStopActivity := false

			// Calculate wait time
			wait := time.Duration(actions[indexPosition].duration) * time.Minute

			if cfg.Debug {
				log.Println("================================================")
				log.Println("Collection:      ", productID, firstName, lastName)
			}

			// Make sure we aren't at the end of the process
			// and sufficient time has passed to proceed
			if !actions[indexPosition].processEnd && time.Since(timestamp) > wait {

				if cfg.Debug {
					log.Println("================================================")
					log.Printf("Current process:  %d", processes[i].processID)
					log.Println("Current activity: " + strconv.Itoa(actions[indexPosition].actionID) + " - " + actions[indexPosition].actionDesc)
					log.Println("Next activity:    " + strconv.Itoa(actions[nextActivity].actionID) + " - " + actions[nextActivity].actionDesc)
					log.Printf("Timestamp:        %v\n", timestamp)
					log.Printf("Wait for:         %v\n", wait)
					log.Printf("Time since:       %v\n", time.Since(timestamp))
					log.Println("It's GO TIME!     " + strconv.Itoa(actions[nextActivity].actionID) + " - " + actions[nextActivity].actionDesc)
				}

				// You begin a transaction with a call to db.Begin(),
				// and close it with a Commit() or Rollback() method
				// on the resulting Tx variable. Under the covers,
				// the Tx gets a connection from the pool, and reserves
				// it for use only with that transaction. The methods
				// on the Tx map one-for-one to methods you can call
				// on the database itself, such as Query() and so forth.

				// Begin database transaction
				tx, err3 := db.Begin()
				if err3 != nil {
					log.Println(err3)
				}

				// prepare activity insert
				activityInsert, err4 := dot.Prepare(tx, "INSERT_ACTIVITY")
				if err4 != nil {
					log.Println(err4)
					tx.Rollback()
				}

				// prepare email insert
				emailInsert, err5 := dot.Prepare(tx, "INSERT_EMAIL")
				if err5 != nil {
					log.Println(err5)
					tx.Rollback()
				}

				// prepare sms insert
				smsInsert, err6 := dot.Prepare(tx, "INSERT_SMS")
				if err6 != nil {
					log.Println(err6)
					tx.Rollback()
				}

				// prepare direct mail insert
				dmInsert, err7 := dot.Prepare(tx, "INSERT_DIRECT_MAIL")
				if err7 != nil {
					log.Println(err7)
					tx.Rollback()
				}

				// prepare collections payment insert (Autodebit)
				paymentInsert, err8 := dot.Prepare(tx, "INSERT_COLLECTION_PAYMENT")
				if err8 != nil {
					log.Println(err8)
					tx.Rollback()
				}

				// prepare collections item update
				itemUpdate, err9 := dot.Prepare(tx, "UPDATE_COLLECTION_ITEM")
				if err9 != nil {
					log.Println(err9)
					tx.Rollback()
				}

				// prepare collections item update (mark paid)
				itemPaid, err10 := dot.Prepare(tx, "PAY_COLLECTION_ITEM")
				if err10 != nil {
					log.Println(err10)
					tx.Rollback()
				}

				// prepare collections item update (GIACT code)
				itemGiact, err11 := dot.Prepare(tx, "UPDATE_GIACT_CODE")
				if err10 != nil {
					log.Println(err11)
					tx.Rollback()
				}

				// prepare collections item update (mark paid)
				itemPaidGiact, err12 := dot.Prepare(tx, "PAY_GIACT_COLLECTION_ITEM")
				if err12 != nil {
					log.Println(err12)
					tx.Rollback()
				}

				// Do we have an email to send?  If the email address is
				// not null we should send an email.
				if actions[nextActivity].email != "" {

					// create payment URL if one doesn't already exist
					if paymentURL == "" {

						// How long before we expire the token?
						tokenLifespan := time.Duration(24*180) * time.Hour // six months

						// Typically I would use a normal timestamp but since
						// the web site will need to compare the expiration time
						// this is a better format. Unix() returns t as a Unix time
						// (the number of seconds elapsed since January 1, 1970 UTC).
						// So we have to multiply by 1000 because in Javascript
						// The Date.now() method returns the number of
						// *milliseconds* elapsed since 1 January 1970 00:00:00 UTC.

						// convert to UNIX time to Javascript time (milliseconds)
						expires := time.Now().Add(tokenLifespan).Unix() * 1000

						// create token & hash
						token, hash, err4 := generateToken()
						u.Check(err4)

						// make the url to send
						siteRoot := ""

						if processes[i].processID == 10 || processes[i].processID == 20 {
							siteRoot = cfg.Site.Intuit
						}
						if processes[i].processID == 30 || processes[i].processID == 40 {
							siteRoot = cfg.Site.TaxSayer
						}

						paymentURL = siteRoot + "/payment/" + strconv.Itoa(productID) + "/" + token + "/" + strconv.Itoa(processes[i].processID)

						// Update item record with URL, and hash of the token and expiration
						if cfg.Debug {
							log.Println("  - Updating Item Record...")
						}
						_, err = itemUpdate.Exec(
							paymentURL,
							hash,
							expires,
							productID)
						if err != nil {
							log.Println(err)
							tx.Rollback()
						}
					}

					// set email provider to either Responsys or SendGrid
					emailProvider := 0
					if processes[i].processID == 10 || processes[i].processID == 20 {
						emailProvider = responsys
					}
					if processes[i].processID == 30 || processes[i].processID == 40 {
						emailProvider = sendgrid
					}

					// insert email record
					if cfg.Debug {
						log.Println("  - Writing Email Record...")
					}
					_, err = emailInsert.Exec(
						productID,
						firstName,
						middleInitial,
						lastName,
						email,
						actions[nextActivity].subject,
						amount,
						paymentURL,
						actions[nextActivity].fromName,
						actions[nextActivity].fromEmail,
						actions[nextActivity].email, // template to use
						emailProvider)
					if err != nil {
						log.Println(err)
						tx.Rollback()
					}

					report.EmailsSent++
				}

				// Do we have an SMS to send? If SMS is not null we should
				// send an SMS.
				if actions[nextActivity].sms != "" {
					if cfg.Debug {
						log.Println("  - Writing SMS Record...")
					}
					_, err = smsInsert.Exec(
						productID,
						firstName,
						middleInitial,
						lastName,
						phone,
						actions[nextActivity].fromName,
						actions[nextActivity].sms)
					if err != nil {
						log.Println(err)
						tx.Rollback()
					}

					report.SMSSent++
				}

				// Is this an Autodebit step?
				if actions[nextActivity].autoDebit == true {

					// GIACT Data
					giactAcct.UniqueID = strconv.Itoa(productID)
					giactAcct.Check.RoutingNumber = rtn
					giactAcct.Check.AccountNumber = dan
					giactAcct.Check.CheckAmount = amount
					giactAcct.Check.AccountType = checkingAcct
					giactAcct.GVerifyEnabled = true

					// Pass in an account and process id and have the verify
					// function call the GIACT API and return the response.
					err1 := verify(&giactAcct, processes[i].processID)
					if err1 != nil {
						// Program needs to continue processing so lets
						// just log GIACT errors.
						log.Println(err1)
						report.GiactError++

						// NOTE if we get an error from GIACT we are going to assume
						// the account is bad and halt the auto debit process. Usually
						// we see 400 errors for accounts with characters

						// Update item record with GIACT err code
						if cfg.Debug {
							log.Println("  - Updating Item with GIACT err code...")
						}
						_, err = itemGiact.Exec(
							giactErr,
							productID)
						if err != nil {
							log.Println(err)
							tx.Rollback()
						}

						// Set GIACT stop
						giactStopActivity = true

					} else {

						// Lookup the GIACT account response code in the map
						// to get the boolean flag (if we should debit or not).
						autodebit, ok := giactList[response.AccountResponseCode]

						// Update item record with GIACT code
						if cfg.Debug {
							log.Println("  - Updating Item with GIACT code...")
						}
						_, err = itemGiact.Exec(
							response.AccountResponseCode,
							productID)
						if err != nil {
							log.Println(err)
							tx.Rollback()
						}

						if !ok || !autodebit {

							// Set GIACT Stop
							giactStopActivity = true
						} else {

							// Update item record as paid
							if cfg.Debug {
								log.Println("  - Updating Item Record as paid...")
							}
							_, err = itemPaidGiact.Exec(
								paidTrue,
								response.AccountResponseCode,
								productID)
							if err != nil {
								log.Println(err)
								tx.Rollback()
							}

							// Insert payment record
							if cfg.Debug {
								log.Println("  - Writing Collections Payment Record...")
							}
							_, err = paymentInsert.Exec(
								productID,
								amount,
								time.Now(),
								achPayment,
								customerInitiatedFalse,
								achRetryFalse)
							if err != nil {
								log.Println(err)
								tx.Rollback()
							}

							report.PaymentsCreated++
						}
					}
				}

				// Is this a *retry* step?
				if actions[nextActivity].debitRetry || actions[nextActivity].customerRetry {

					// NOTE: We don't do GIACT calls on retries
					// since we did it the first time (at least for autodebits)

					// setup type of retry
					customerInitiatedFlag := false // autodebit retry
					if actions[nextActivity].customerRetry {
						customerInitiatedFlag = true // customer retry
					}

					// Update item record as paid
					if cfg.Debug {
						log.Println("  - Updating Item Record as paid (again)...")
					}
					_, err = itemPaid.Exec(
						paidTrue,
						productID)
					if err != nil {
						log.Println(err)
						tx.Rollback()
					}

					// Insert payment record
					if cfg.Debug {
						log.Println("  - Writing Collections Payment Record...")
					}
					_, err = paymentInsert.Exec(
						productID,
						amount,
						time.Now(),
						achPayment,
						customerInitiatedFlag,
						achRetryTrue) // set retry flag to true
					if err != nil {
						log.Println(err)
						tx.Rollback()
					}

					report.PaymentsCreated++
				}

				// Is it direct mail time?
				if actions[nextActivity].directMail == true {

					// Insert direct mail record
					if cfg.Debug {
						log.Println("  - Writing Direct Mail Record...")
					}
					_, err = dmInsert.Exec(productID)
					if err != nil {
						log.Println(err)
						tx.Rollback()
					}

					report.DirectMailCreated++
				}

				// Write Activity Record
				actionStep := actions[nextActivity].actionID

				if giactStopActivity {
					// Lookup the right GIACT code by processID
					actionStep = GIACTStopCode[processes[i].processID]
				}

				if cfg.Debug {
					log.Println("  - Writing Activity Record...")
				}
				_, err = activityInsert.Exec(
					productID,
					processes[i].processID,
					actionStep)
				if err != nil {
					log.Println(err)
					tx.Rollback()
				}

				report.Activities++

				// commit database transaction
				err = tx.Commit()
				if err != nil {
					log.Println(err)
				}

			} else {
				if cfg.Debug {
					log.Println("================================================")
					log.Printf("Current process:  %d", processes[i].processID)
					log.Println("Current activity: " + strconv.Itoa(actions[indexPosition].actionID) + " - " + actions[indexPosition].actionDesc)
					log.Println("Next activity:    " + strconv.Itoa(actions[nextActivity].actionID) + " - " + actions[nextActivity].actionDesc)
					log.Printf("Timestamp:        %v\n", timestamp)
					log.Printf("Wait for:         %v\n", wait)
					log.Printf("Time since:       %v\n", time.Since(timestamp))
					log.Println("GO time?          " + strconv.FormatBool(actions[indexPosition].processEnd))
				}
			}

		}
		err = rows.Err()
		u.Check(err)

	}
	return nil
}
