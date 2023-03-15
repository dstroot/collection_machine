// Copyright (c) 2017, Tax Products Group, LLC.
//
// All rights reserved.

package main

import (
	"fmt"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
)

var giactTests = []struct {
	rtn      string
	dan      string
	expected int
	code     string
	debit    bool
}{
	{"122105278", "0000000001", 19, "GN01", true},
	{"122105278", "0000000002", 20, "GN05", false},
	{"122105278", "0000000003", 5, "GP01", true},
	{"122105278", "0000000004", 1, "GS01", false},
	{"122105278", "0000000005", 2, "GS02", false},
	{"122105278", "0000000006", 3, "GS03", true},
	{"122105278", "0000000007", 4, "GS04", true},
	{"122105278", "0000000008", 21, "ND00", true},
	{"122105278", "0000000009", 22, "ND01", false},
	{"122105278", "0000000010", 6, "RT00", false},
	{"122105278", "0000000011", 7, "RT01", false},
	{"122105278", "0000000012", 8, "RT02", false},
	{"122105278", "0000000013", 9, "RT03", true},
	{"122105278", "0000000014", 10, "RT04", false},
	{"122105278", "0000000015", 11, "RT05", false},
	{"122105278", "0000000016", 12, "1111", true},
	{"122105278", "0000000017", 13, "2222", true},
	{"122105278", "0000000018", 14, "3333", true},
	{"122105278", "0000000019", 15, "5555", true},
	{"122105278", "0000000020", 16, "7777", true},
	{"122105278", "0000000021", 17, "8888", true},
	{"122105278", "0000000022", 18, "9999", true},
}

func TestVerify(t *testing.T) {
	Convey("GIACT test data", t, func() {
		Convey("should return expected codes\n", func() {

			fmt.Printf("\n\n")
			err := initialize()
			So(err, ShouldEqual, nil)
			So(cfg.GiactURL, ShouldNotEqual, "")
			cfg.Debug = true // turn debug off
			fmt.Printf("\nInitialization Complete\n\n")

			// Get GIACT codes
			err = getGiactActions()
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			for _, test := range giactTests {

				// Test Data
				account := &giactReq{
					UniqueID: uuid.NewV4().String(),
					Check: Check{
						RoutingNumber: test.rtn,
						AccountNumber: test.dan,
						CheckAmount:   99.99,
						AccountType:   1,
					},
					GVerifyEnabled: true,
				}

				// Call GIACT API
				err1 := verify(account, 10)
				So(err1, ShouldEqual, nil)
				So(response.AccountResponseCode, ShouldNotBeNil)
				So(response.AccountResponseCode, ShouldEqual, test.expected)

				// Test the GIACT response code to see if it "passed"
				autodebit, ok := giactList[response.AccountResponseCode]
				log.Printf("Response: %d\n", response.AccountResponseCode)
				So(ok, ShouldEqual, true) // should be in map
				So(autodebit, ShouldEqual, test.debit)

				if autodebit != test.debit {
					t.Errorf("GIACT(%s, %s): expected %d (%s) pass = %t, actual %d pass = %t\n", test.rtn, test.dan, test.expected, test.code, test.debit, response.AccountResponseCode, autodebit)
				} else {
					fmt.Printf("GIACT(%s, %s): expected %d (%s) pass = %t, actual %d pass = %t\n", test.rtn, test.dan, test.expected, test.code, test.debit, response.AccountResponseCode, autodebit)
				}

			}
		})
	})
}

// For sandbox gVerify and gAuthenticate inquiries, you must use test bank account data. The table below
// contains the default set of test accounts that return a set AccountResponseCode value.
// NOTE: If you submit an inquiry with data that is not found in the test bank account list, the
// AccountResponseCode value will be “ND00”. If you need to add your own test accounts, please
// contact the integration support team.
//
// NOTE: CD = AccountResponseCode (what the REST API actually returns)
//
// RTN      , DAN       , CD, Code, Description
// -------------------------------------------------------------------------------------------------
// 122105278, 0000000001, 19, GN01, "Negative information was found in the account’s history"
// 122105278, 0000000002, 20, GN05, "Routing number is not assigned to a financial institution"
// 122105278, 0000000003,  5, GP01, "Account found in your API user’s Private Bad Checks list"
// 122105278, 0000000004,  1, GS01, "Invalid/Unassigned routing number"
// 122105278, 0000000005,  2, GS02, "Invalid account number"
// 122105278, 0000000006,  3, GS03, "Invalid check number"
// 122105278, 0000000007,  4, GS04, "Invalid check amount"
// 122105278, 0000000008, 21, ND00, "No positive or negative information available for the account information"
// 122105278, 0000000009, 22, ND01, "Routing number can only be valid for a US Government institution"
// 122105278, 0000000010,  6, RT00, "Routing number is participating bank, but account number not located"
// 122105278, 0000000011,  7, RT01, "Account should be declined based on the risk factor reported"
// 122105278, 0000000012,  8, RT02, "Item (Check Number) should be declined based on the risk factor reported"
// 122105278, 0000000013,  9, RT03, "Current negative data exists on the account. Ex: NSF or recent returns"
// 122105278, 0000000014, 10, RT04, "Non-Demand Deposit Account (Post No Debits)"
// 122105278, 0000000015, 11, RT05, "Recent negative data exists on the account. Ex: NSF or recent returns"
// 122105278, 0000000016, 12, 1111, "Account Verified – Open and valid checking or savings account"
// 122105278, 0000000017, 13, 2222, "AMEX – The account is an American Express Travelers Cheque account"
// 122105278, 0000000018, 14, 3333, "Non-Participant Provider – Account reported as having positive data"
// 122105278, 0000000019, 15, 5555, "Savings Account Verified – Open and valid savings account"
// 122105278, 0000000020, 16, 7777, "Checking or savings account was found to have positive historical data"
// 122105278, 0000000021, 17, 8888, "Savings account was found to have positive historical data"
// 122105278, 0000000022, 18, 9999, "Account reported as having positive historical data"

// AccountResponseCode
// ------------------------------------
// Debit	Index	Value	Description
// ------------------------------------
// n/a		0		Null	There is no AccountResponseCode value for this result.
// NO		1		GS01	Invalid Routing Number - The routing number supplied fails the validation test.
// NO		2		GS02	Invalid Account Number - The account number supplied fails the validation test.
// YES		3		GS03	Invalid Check Number - The check number supplied fails the validation test.
// YES		4		GS04	Invalid Amount - The amount supplied fails the validation test.
// YES		5		GP01	The account was found as active in your Private Bad Checks List.
// NO		6		RT00	The routing number belongs to a reporting bank; however, no positive nor negative information has been reported on the account number.
// NO		7		RT01	This account should be declined based on the risk factor being reported.
// NO		8		RT02	This item should be rejected based on the risk factor being reported.
// YES		9		RT03	Current negative data exists on this account. Accept transaction with risk.
// NO		10		RT04	Non-Demand Deposit Account (post no debits), Credit Card Check, Line of Credit, Home Equity, or a Brokerage check.
// n/a		11		RT05	N/A
// YES		12		_1111	Account Verified – The account was found to be an open and valid checking account.
// YES		13		_2222	AMEX – The account was found to be an American Express Travelers Cheque account.
// YES		14		_3333	Non-Participant Provider – This account was reported with acceptable, positive data found in current or recent transactions.
// YES		15		_5555	Savings Account Verified – The account was found to be an open and valid savings account.
// YES		16		_7777	N/A
// YES		17		_8888	N/A
// YES		18		_9999	N/A
// YES		19		GN01	Negative information was found in this account's history.
// NO		20		GN05	The routing number is reported as not currently assigned to a financial institution.
// YES		21		ND00	No positive or negative information has been reported on the account.
// ???		22		ND01	This routing number can only be valid for US Government financial institutions.

// NOTE: AccountResponseCode values _7777, _8888, _9999, and RT05 are only applicable to API users that are restricted to the Real-Time Extended process. If your API user does not have to use the Real-Time Extended process, then you will never see these values returned.
