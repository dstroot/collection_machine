// Copyright (c) 2017, Tax Products Group, LLC.
//
// All rights reserved.

package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type item struct {
	productID     int
	orderID       string
	firstName     string
	middleInitial string
	lastName      string
	address1      string
	city          string
	state         string
	postalCode    string
	email         string
	phone         string
	amount        float64
	rtn           string
	dan           string
	taxAmount     float64
	process       int
}

type step struct {
	processID  int
	activityID int
	duration   int
	processEnd int
	emailFrom  string
	email      string
	subject    string
	template   string
}

type activity struct {
	productID int
	process   int
	activity  int
}

//       *** GIACT TEST ACCOUNTS ***
// {"122105278", "0000000001", 19, "GN01", true},
// {"122105278", "0000000002", 20, "GN05", false},
// {"122105278", "0000000003", 5, "GP01", true},
// {"122105278", "0000000004", 1, "GS01", false},
// {"122105278", "0000000005", 2, "GS02", false},
// {"122105278", "0000000006", 3, "GS03", true},
// {"122105278", "0000000007", 4, "GS04", true},
// {"122105278", "0000000008", 21, "ND00", true},
// {"122105278", "0000000009", 22, "ND01", false},
// {"122105278", "0000000010", 6, "RT00", false},
// {"122105278", "0000000011", 7, "RT01", false},
// {"122105278", "0000000012", 8, "RT02", false},
// {"122105278", "0000000013", 9, "RT03", true},
// {"122105278", "0000000014", 10, "RT04", false},
// {"122105278", "0000000015", 11, "RT05", false},
// {"122105278", "0000000016", 12, "1111", true},
// {"122105278", "0000000017", 13, "2222", true},
// {"122105278", "0000000018", 14, "3333", true},
// {"122105278", "0000000019", 15, "5555", true},
// {"122105278", "0000000020", 16, "7777", true},
// {"122105278", "0000000021", 17, "8888", true},
// {"122105278", "0000000022", 18, "9999", true},

//ToNullString invalidates a sql.NullString if empty, validates if not empty
func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func insertTestData() {

	steps := []step{
		{10, 0, 0, 0, "", "", "", ""},
		{10, 1, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Your TurboTax Payment is Due", "1"},
		{10, 2, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Reminder: Your TurboTax Payment is Due", "2"},
		{10, 3, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Important Notice of Action: Bank Account Debit", "3"},
		{10, 4, 0, 1, "", "", "", ""},
		{10, 5, 0, 0, "", "", "", ""},
		{10, 6, 0, 1, "Intuit", "dan.stroot@sbtpg.com", "Important Request for Action: Balance Past Due", "5"},
		{10, 7, 0, 0, "", "", "", ""},
		{10, 8, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Important Notice of Action: Bank Account Retry Pending", "4"},
		{10, 9, 0, 1, "", "", "", ""},
		{10, 10, 0, 0, "", "", "", ""},
		{10, 11, 0, 1, "Intuit", "dan.stroot@sbtpg.com", "Important Request for Action: Balance Past Due", "5"},
		{10, 21, 0, 1, "", "", "", ""},
		{10, 22, 0, 1, "", "", "", ""},
		{10, 23, 0, 0, "", "", "", ""},
		{10, 24, 0, 1, "Intuit", "dan.stroot@sbtpg.com", "Important Request for Action: Balance Past Due", "5"},
		{10, 25, 0, 0, "", "", "", ""},
		{10, 26, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Important Notice of Action: Bank Account Retry Pending", "4"},
		{10, 27, 0, 1, "", "", "", ""},
		{10, 28, 0, 0, "", "", "", ""},
		{10, 29, 0, 1, "Intuit", "dan.stroot@sbtpg.com", "Important Request for Action: Balance Past Due", "5"},

		{10, 30, 0, 0, "", "", "", ""},
		{10, 31, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Autodebit Deferral Completed", "3"},
		{10, 32, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Autodebit Deferral Completed", "3"},
		{10, 33, 0, 1, "", "", "", ""},

		{10, 35, 0, 0, "", "", "", ""},
		{10, 36, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Autodebit Deferral Completed", "3"},
		{10, 37, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Autodebit Deferral Completed", "3"},
		{10, 38, 0, 1, "", "", "", ""},

		{10, 40, 0, 0, "", "", "", ""},
		{10, 41, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Autodebit Deferral Completed", "3"},
		{10, 42, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Autodebit Deferral Completed", "3"},
		{10, 43, 0, 1, "", "", "", ""},

		{10, 70, 0, 1, "", "", "", ""},
		{10, 71, 0, 1, "Intuit", "dan.stroot@sbtpg.com", "Important Request for Action: Balance Past Due", "5"},

		{10, 75, 0, 0, "", "", "", ""},
		{10, 76, 0, 1, "Intuit", "dan.stroot@sbtpg.com", "Important Request for Action: Balance Past Due", "5"},

		{10, 79, 0, 1, "", "", "", ""},
		{10, 89, 0, 1, "", "", "", ""},
		{10, 90, 0, 1, "", "", "", ""},

		{10, 99, 0, 1, "", "", "", ""},

		{20, 0, 0, 0, "", "", "", ""},
		{20, 1, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Your TurboTax Payment is Due", "1"},
		{20, 2, 0, 0, "Intuit", "dan.stroot@sbtpg.com", "Reminder: Your TurboTax Payment is Due", "2"},
		{20, 15, 0, 1, "", "", "", ""},
		{20, 21, 0, 1, "", "", "", ""},
		{20, 79, 0, 1, "", "", "", ""},
		{20, 89, 0, 1, "", "", "", ""},
		{20, 99, 0, 1, "", "", "", ""},

		{30, 0, 0, 0, "", "", "", ""},
		{30, 1, 0, 0, "Taxslayer", "dan.stroot@sbtpg.com", "Your Taxslayer Payment is Due", "1"},
		{30, 2, 0, 0, "Taxslayer", "dan.stroot@sbtpg.com", "Reminder: Your Taxslayer Payment is Due", "2"},
		{30, 3, 0, 0, "Taxslayer", "dan.stroot@sbtpg.com", "Important Notice of Action: Bank Account Debit", "3"},
		{30, 4, 0, 1, "", "", "", ""},
		{30, 5, 0, 0, "", "", "", ""},
		{30, 6, 0, 1, "Taxslayer", "dan.stroot@sbtpg.com", "Important Request for Action: Balance Past Due", "5"},
		{30, 7, 0, 0, "", "", "", ""},
		{30, 8, 0, 0, "Taxslayer", "dan.stroot@sbtpg.com", "Important Notice of Action: Bank Account Retry Pending", "4"},
		{30, 9, 0, 1, "", "", "", ""},
		{30, 10, 0, 0, "", "", "", ""},
		{30, 11, 0, 1, "Taxslayer", "dan.stroot@sbtpg.com", "Important Request for Action: Balance Past Due", "5"},
		{30, 21, 0, 1, "", "", "", ""},
		{30, 22, 0, 1, "", "", "", ""},
		{30, 23, 0, 0, "", "", "", ""},
		{30, 24, 0, 1, "Taxslayer", "dan.stroot@sbtpg.com", "Important Request for Action: Balance Past Due", "5"},
		{30, 25, 0, 0, "", "", "", ""},
		{30, 26, 0, 0, "Taxslayer", "dan.stroot@sbtpg.com", "Important Notice of Action: Bank Account Retry Pending", "4"},
		{30, 27, 0, 1, "", "", "", ""},
		{30, 28, 0, 0, "", "", "", ""},
		{30, 29, 0, 1, "Taxslayer", "dan.stroot@sbtpg.com", "Important Request for Action: Balance Past Due", "5"},
		{30, 79, 0, 1, "", "", "", ""},
		{30, 89, 0, 1, "", "", "", ""},
		{30, 99, 0, 1, "", "", "", ""},

		{40, 0, 0, 0, "", "", "", ""},
		{40, 1, 0, 0, "Taxslayer", "dan.stroot@sbtpg.com", "Your Taxslayer Payment is Due", "1"},
		{40, 2, 0, 0, "Taxslayer", "dan.stroot@sbtpg.com", "Reminder: Your Taxslayer Payment is Due", "2"},
		{40, 15, 0, 1, "", "", "", ""},
		{40, 21, 0, 1, "", "", "", ""},
		{40, 79, 0, 1, "", "", "", ""},
		{40, 89, 0, 1, "", "", "", ""},
		{40, 99, 0, 1, "", "", "", ""},
	}

	query1 := `
	IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='PROCESS_DEFINITION')
	BEGIN
	  DROP TABLE dbo.PROCESS_DEFINITION;
	END;

	CREATE TABLE dbo.PROCESS_DEFINITION
	(
	  PROCESS_DEFINITION_ID   INT IDENTITY(1,1) ,
	  PROCESS_TYPE_ID         INT           NOT NULL ,
	  ACTIVITY_TYPE_ID        INT           NOT NULL ,
	  DURATION_MINUTES        BIGINT        NOT NULL ,
	  PROCESS_END             BIT           NOT NULL ,
	  EMAIL_FROM_NAME         VARCHAR(56)   NULL ,
	  EMAIL_FROM_EMAIL        VARCHAR(56)   NULL ,
	  EMAIL_SUBJECT           VARCHAR(100)  NULL ,
	  EMAIL_TEMPLATE          VARCHAR(25)   NULL ,
	  SMS_TEMPLATE            VARCHAR(25)   NULL ,
	  ACTIVE_START_DATE       DATETIME      NULL ,
	  ACTIVE_END_DATE         DATETIME      NULL ,
	  CREATION_DATE           DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
	  MODIFY_DATE             DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
	  CREATED_BY              VARCHAR(56)   NULL DEFAULT SUSER_SNAME() ,
	  MODIFIED_BY             VARCHAR(56)   NULL DEFAULT SUSER_SNAME()
	);

	ALTER TABLE dbo.PROCESS_DEFINITION
	  ADD CONSTRAINT [PK_PROCESS_DEFINITION_ID] PRIMARY KEY  CLUSTERED ([PROCESS_DEFINITION_ID] ASC);

	CREATE NONCLUSTERED INDEX [IDX_PROCESS_TYPE_ID] ON [PROCESS_DEFINITION]
	(
	  [PROCESS_TYPE_ID]         ASC
	);

	CREATE NONCLUSTERED INDEX [IDX_ACTIVITY_TYPE_ID] ON [PROCESS_DEFINITION]
	(
	  [ACTIVITY_TYPE_ID]         ASC
	);`

	_, err := db.Exec(query1)
	So(err, ShouldBeNil)

	query2 := `INSERT INTO dbo.PROCESS_DEFINITION (
		PROCESS_TYPE_ID,
		ACTIVITY_TYPE_ID,
		DURATION_MINUTES,
		PROCESS_END,
		EMAIL_FROM_NAME,
		EMAIL_FROM_EMAIL,
		EMAIL_SUBJECT,
		EMAIL_TEMPLATE
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// Insert process
	for _, i := range steps {
		_, err := db.Exec(query2, i.processID, i.activityID, i.duration, i.processEnd, ToNullString(i.emailFrom), ToNullString(i.email), ToNullString(i.subject), ToNullString(i.template))
		So(err, ShouldBeNil)
	}

	// Collection items - there is at least one item for each step in the process definition table
	// each step must be tested for correct behavior.  There are extra steps to validate GIACT behavior.
	items := []item{
		// Process 10
		{1110, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1111, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1112, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1113, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1114, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1115, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1116, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1117, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1118, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1119, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1120, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1121, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1122, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1123, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1124, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1125, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1126, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1127, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1128, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1129, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1130, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1131, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1131, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1132, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1133, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1134, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1135, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1136, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1137, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1138, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1139, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1140, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1141, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1141, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1142, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1143, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1144, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1145, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1146, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1147, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1148, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1149, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1150, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},

		// Process 20
		{1151, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 20},
		{1152, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 20},
		{1153, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 20},
		{1154, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 20},
		{1155, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 20},
		{1156, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 20},
		{1157, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 20},
		{1158, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 20},

		// Process 30
		{1160, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1161, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1162, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1163, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1164, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1165, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1166, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1167, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1168, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1169, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1170, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1171, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1172, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1173, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1174, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1175, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1176, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1177, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1178, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1179, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1180, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1182, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1182, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},
		{1183, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 30},

		// Process 40
		{1190, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 40},
		{1191, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 40},
		{1192, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 40},
		{1193, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 40},
		{1194, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 40},
		{1195, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 40},
		{1196, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 40},
		{1197, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 40},

		// GIACT tests (DAN with characters, Good DAN, Bad DAN)
		{1201, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "sally", 0, 10},
		{1202, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000001", 0, 10},
		{1203, "EFE432XQX14JDR59", "Bill", "M", "Madison", "1223 Main Street", "New Brunswick", "CA", "92101", "dan.stroot@gmail.com", "19494634044", 39.95, "122105278", "0000000002", 0, 10},
	}

	query3 := `
	INSERT INTO COLLECTION_ITEM (
		REFUND_TRANSACTION_ID,
		PROVIDER_ORDER_ID,
		FIRST_NAME,
		MIDDLE_INITIAL,
		LAST_NAME,
		ADDRESS1,
		CITY,
		STATE,
		ZIPCODE,
		EMAIL,
		MOBILE_PHONE,
		AMOUNT_OWED,
		RTN,
		DAN,
		TAX_AMOUNT_OWED,
		PROCESS_TYPE_ID
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Insert items
	for _, i := range items {
		_, err := db.Exec(query3, i.productID, i.orderID, i.firstName, i.middleInitial, i.lastName, i.address1, i.city, i.state, i.postalCode, i.email, i.phone, i.amount, i.rtn, i.dan, i.taxAmount, i.process)
		So(err, ShouldBeNil)
	}

	// Activity steps
	activities := []activity{

		// process 10
		{1110, 10, 0},  // Start (email and SMS)
		{1111, 10, 1},  // Email 1
		{1112, 10, 2},  // Email 2
		{1113, 10, 3},  // Email 3
		{1114, 10, 4},  // Auto Debit (stop)
		{1115, 10, 5},  // Permanent Failure
		{1116, 10, 6},  // Final Email (stop)
		{1117, 10, 7},  // Failure to retry
		{1118, 10, 8},  // Retry email
		{1119, 10, 9},  // Autodebit retry (stop)
		{1120, 10, 10}, // Second auto debit failed
		{1121, 10, 11}, // Final Email (stop)

		{1122, 10, 21}, // CC payment (stop)
		{1123, 10, 22}, // Bank Payment (stop)
		{1124, 10, 23}, // Bank auth failure
		{1125, 10, 24}, // Final Email (stop)
		{1126, 10, 25}, // Bank auth failure retry
		{1127, 10, 26}, // Retry email
		{1128, 10, 27}, // Retry bank auth payment (stop)
		{1129, 10, 28}, // Second bank auth failed
		{1130, 10, 29}, // Final email (stop)

		{1131, 10, 30}, // Autodebit deferred (email 1)
		{1132, 10, 31}, // Email reminder 1
		{1133, 10, 32}, // Email reminder 2
		{1134, 10, 33}, // Autodebit (stop)

		{1135, 10, 35}, // Autodebit deferred (email 2)
		{1136, 10, 36}, // Email reminder 1
		{1137, 10, 37}, // Email reminder 2
		{1138, 10, 38}, // Autodebit (stop)

		{1139, 10, 40}, // Autodebit deferred (email 3)
		{1140, 10, 41}, // Email reminder 1
		{1141, 10, 42}, // Email reminder 2
		{1142, 10, 43}, // Autodebit (stop)

		{1143, 10, 70}, // Stopped Autodebit (agent)
		{1144, 10, 71}, // Final email (stop)

		{1145, 10, 75}, // GIACT stop
		{1146, 10, 76}, // Final email (stop)

		{1147, 10, 79}, // Fees Refunded (stop)
		{1148, 10, 89}, // Fees suppressed (stop)
		{1149, 10, 90}, // Fees waived (Agent) (stop)

		{1150, 10, 99}, // Funded (stop)

		// Process 20
		{1151, 20, 0},  // Start (email and SMS)
		{1152, 20, 1},  // Email 1
		{1153, 20, 2},  // Email 2
		{1154, 20, 15}, // Direct mail step (stop)
		{1155, 20, 21}, // CC payment (stop)
		{1156, 20, 79}, // Fees Refunded (stop)
		{1157, 20, 89}, // Fees suppressed (stop)
		{1158, 20, 99}, // Funded (stop)

		// Process 30
		{1160, 30, 0},  // Start (email and SMS)
		{1161, 30, 1},  // Email 1
		{1162, 30, 2},  // Email 2
		{1163, 30, 3},  // Email 3
		{1164, 30, 4},  // Auto Debit (stop)
		{1165, 30, 5},  // Permanent Failure
		{1166, 30, 6},  // Final Email (stop)
		{1167, 30, 7},  // Failure to retry
		{1168, 30, 8},  // Retry email
		{1169, 30, 9},  // Autodebit retry (stop)
		{1170, 30, 10}, // Second auto debit failed
		{1171, 30, 11}, // Final Email (stop)

		{1172, 30, 21}, // CC payment (stop)
		{1173, 30, 22}, // Bank Payment (stop)
		{1174, 30, 23}, // Bank auth failure
		{1175, 30, 24}, // Final Email (stop)
		{1176, 30, 25}, // Bank auth failure retry
		{1177, 30, 26}, // Retry email
		{1178, 30, 27}, // Retry bank auth payment (stop)
		{1179, 30, 28}, // Second bank auth failed
		{1180, 30, 29}, // Final email (stop)

		{1181, 30, 79}, // Fees Refunded (stop)
		{1182, 30, 89}, // Fees suppressed (stop)
		{1183, 30, 99}, // Funded (stop)

		// Process 40
		{1190, 40, 0},  // Start (email and SMS)
		{1191, 40, 1},  // Email 1
		{1192, 40, 2},  // Email 2
		{1193, 40, 15}, // Direct mail step (stop)
		{1194, 40, 21}, // CC payment (stop)
		{1195, 40, 79}, // Fees Refunded (stop)
		{1196, 40, 89}, // Fees suppressed (stop)
		{1197, 40, 99}, // Funded (stop)

		// GIACT
		{1201, 10, 3}, // Bad DAN
		{1202, 10, 3}, // Good DAN
		{1203, 10, 3}, // Bad DAN
	}

	query4 := `INSERT INTO COLLECTION_ACTIVITY (
		REFUND_TRANSACTION_ID,
		PROCESS_TYPE_ID,
		ACTIVITY_TYPE_ID
	) VALUES (?, ?, ?)`

	// Insert steps
	for _, i := range activities {
		_, err1 := db.Exec(query4, i.productID, i.process, i.activity)
		So(err1, ShouldBeNil)
	}
}

func checkResults(t *testing.T) {

	expected := make(map[int]int)

	// define the activity step each record should now be in:

	// process 10
	expected[1110] = 1  // {10, 0}  Start (email and SMS)
	expected[1111] = 2  // {10, 1}  Email 1
	expected[1112] = 3  // {10, 2}  Email 2
	expected[1113] = 4  // {10, 3}  Email 3
	expected[1114] = 4  // {10, 4}  Auto Debit (stop)
	expected[1115] = 6  // {10, 5}  Permanent Failure
	expected[1116] = 6  // {10, 6}  Final Email (stop)
	expected[1117] = 8  // {10, 7}  Failure to retry
	expected[1118] = 9  // {10, 8}  Retry email
	expected[1119] = 9  // {10, 9}  Autodebit retry (stop)
	expected[1120] = 11 // {10, 10} Second auto debit failed
	expected[1121] = 11 // {10, 11} Final Email (stop)

	expected[1122] = 21 // {10, 21} CC payment (stop)
	expected[1123] = 22 // {10, 22} Bank Payment (stop)
	expected[1124] = 24 // {10, 23} Bank auth failure
	expected[1125] = 24 // {10, 24} Final Email (stop)
	expected[1126] = 26 // {10, 25} Bank auth failure retry
	expected[1127] = 27 // {10, 26} Retry email
	expected[1128] = 27 // {10, 27} Retry bank auth payment (stop)
	expected[1129] = 29 // {10, 28} Second bank auth failed
	expected[1130] = 29 // {10, 29} Final email (stop)

	expected[1131] = 31 // {10, 30} Autodebit deferred (email 1)
	expected[1132] = 32 // {10, 31} Email reminder 1
	expected[1133] = 33 // {10, 32} Email reminder 2
	expected[1134] = 33 // {10, 33} Autodebit (stop)

	expected[1135] = 36 // {10, 35} Autodebit deferred (email 2)
	expected[1136] = 37 // {10, 36} Email reminder 2
	expected[1137] = 38 // {10, 37} Email reminder 2
	expected[1138] = 38 // {10, 38} Autodebit (stop)

	expected[1139] = 41 // {10, 40} Autodebit deferred (email 3)
	expected[1140] = 42 // {10, 41} Email reminder 3
	expected[1141] = 43 // {10, 42} Email reminder 3
	expected[1142] = 43 // {10, 43} Autodebit (stop)

	expected[1143] = 70 // {10, 70} Stopped Autodebit (stop)
	expected[1144] = 71 // {10, 71} Final email (stop)

	expected[1145] = 76 // {10, 75} GIACT stop
	expected[1146] = 76 // {10, 76} Final email (stop)

	expected[1147] = 79 // {10, 79} Fees Refunded (stop)
	expected[1148] = 89 // {10, 89} Fees suppressed (stop)
	expected[1149] = 90 // {10, 90} Fees waived (Agent) (stop)

	expected[1150] = 99 // {10, 99} Funded (stop)

	// Process 20
	expected[1151] = 1  // {20, 0}  Start (email and SMS)
	expected[1152] = 2  // {20, 1}  Email 1
	expected[1153] = 15 // {20, 2}  Email 2
	expected[1154] = 15 // {20, 15} Direct mail step (stop)
	expected[1155] = 21 // {20, 21} CC payment (stop)
	expected[1156] = 79 // {20, 79} Fees Refunded (stop)
	expected[1157] = 89 // {20, 89} Fees suppressed (stop)
	expected[1158] = 99 // {20, 99} Funded (stop)

	// Process 30
	expected[1160] = 1  // {30, 0}  Start (email and SMS)
	expected[1161] = 2  // {30, 1}  Email 1
	expected[1162] = 3  // {30, 2}  Email 2
	expected[1163] = 4  // {30, 3}  Email 3
	expected[1164] = 4  // {30, 4}  Auto Debit (stop)
	expected[1165] = 6  // {30, 5}  Permanent Failure
	expected[1166] = 6  // {30, 6}  Final Email (stop)
	expected[1167] = 8  // {30, 7}  Failure to retry
	expected[1168] = 9  // {30, 8}  Retry email
	expected[1169] = 9  // {30, 9}  Autodebit retry (stop)
	expected[1170] = 11 // {30, 10} Second auto debit failed
	expected[1171] = 11 // {30, 11} Final Email (stop)

	expected[1172] = 21 // {30, 21} CC payment (stop)
	expected[1173] = 22 // {30, 22} Bank Payment (stop)
	expected[1174] = 24 // {30, 23} Bank auth failure
	expected[1175] = 24 // {30, 24} Final Email (stop)
	expected[1176] = 26 // {30, 25} Bank auth failure retry
	expected[1177] = 27 // {30, 26} Retry email
	expected[1178] = 27 // {30, 27} Retry bank auth payment (stop)
	expected[1179] = 29 // {30, 28} Second bank auth failed
	expected[1180] = 29 // {30, 29} Final email (stop)

	expected[1181] = 79 // {30, 79} Fees Refunded (stop)
	expected[1182] = 89 // {30, 89} Fees suppressed (stop)
	expected[1183] = 99 // {30, 99} Funded (stop)

	// Process 40
	expected[1190] = 1  // {40, 0}  Start (email and SMS)
	expected[1191] = 2  // {40, 1}  Email 1
	expected[1192] = 15 // {40, 2}  Email 2
	expected[1193] = 15 // {40, 15} Direct mail step (stop)
	expected[1194] = 21 // {40, 21} CC payment (stop)
	expected[1195] = 79 // {40, 79} Fees Refunded (stop)
	expected[1196] = 89 // {40, 89} Fees suppressed (stop)
	expected[1197] = 99 // {40, 99} Funded (stop)

	// GIACT
	expected[1201] = 75 // Bad DAN
	expected[1202] = 4  // Good DAN
	expected[1203] = 75 // Bad DAN

	fmt.Printf("\n\n")

	query := `
	SELECT
	  I.REFUND_TRANSACTION_ID,
	  tmp.PROCESS_TYPE_ID,
	  tmp.ACTIVITY_TYPE_ID
	FROM
	(
		SELECT C.REFUND_TRANSACTION_ID, C.PROCESS_TYPE_ID, C.ACTIVITY_TYPE_ID, C.TIMESTAMP_UTC,
		Rank() over (Partition BY C.REFUND_TRANSACTION_ID ORDER BY C.TIMESTAMP_UTC DESC, C.PROCESS_TYPE_ID DESC,
		C.ACTIVITY_TYPE_ID DESC) AS LAST_ACTIVITY
		FROM OLTP_SYS.dbo.COLLECTION_ACTIVITY C
	) tmp
	INNER JOIN OLTP_SYS.dbo.PROCESS_DEFINITION D ON tmp.PROCESS_TYPE_ID = D.PROCESS_TYPE_ID
		AND tmp.ACTIVITY_TYPE_ID = D.ACTIVITY_TYPE_ID
	INNER JOIN OLTP_SYS.dbo.COLLECTION_ITEM I ON tmp.REFUND_TRANSACTION_ID = I.REFUND_TRANSACTION_ID
	WHERE LAST_ACTIVITY = 1
		ORDER BY I.REFUND_TRANSACTION_ID ASC`

	rows, err := db.Query(query)
	So(err, ShouldBeNil)
	defer rows.Close()

	var (
		productID int
		process   int
		activity  int
	)

	fmt.Printf("\n")

	// iterate the rows
	for rows.Next() {
		// scan the row into vars
		err1 := rows.Scan(
			&productID,
			&process,
			&activity)
		So(err1, ShouldBeNil)

		// Test the record agianst the map of expected results
		log.Printf("Product ID: %d, Process: %d, Activity ID: %d", productID, process, activity)
		step, ok := expected[productID]

		So(ok, ShouldEqual, true)
		So(step, ShouldEqual, activity)

		fmt.Printf("Process: %d, record (%d): expected activity step: %d, actual: %d\n", process, productID, step, activity)
	}

	err3 := rows.Err()
	So(err3, ShouldBeNil)
}

func deleteTestData() {
	_, err := db.Exec("DELETE FROM COLLECTION_ITEM WHERE REFUND_TRANSACTION_ID >= 1110")
	So(err, ShouldBeNil)

	_, err = db.Exec("DELETE FROM COLLECTION_ACTIVITY WHERE REFUND_TRANSACTION_ID >= 1110")
	So(err, ShouldBeNil)

	_, err = db.Exec("DELETE FROM EMAIL_QUEUE WHERE REFUND_TRANSACTION_ID >= 1110")
	So(err, ShouldBeNil)
}

func TestProcessCollections(t *testing.T) {
	Convey("Process collections", t, func() {
		Convey("should run cleanly\n\n", func() {

			err := initialize()
			So(err, ShouldBeNil)
			cfg.Debug = true

			// get a db connection
			err1 := setupDatabase()
			So(err1, ShouldBeNil)
			defer db.Close()

			fmt.Printf("Load test data: \n\n")

			insertTestData()
			defer deleteTestData()

			fmt.Printf("\n\n")

			err2 := processCollections()
			So(err2, ShouldBeNil)

			checkResults(t)
		})
	})
}

func TestGenerateToken(t *testing.T) {
	Convey("Function generateToken", t, func() {
		Convey("Should generate tokens", func() {
			token, hash, err := generateToken()
			So(err, ShouldBeNil)
			So(token, ShouldNotBeEmpty)
			So(hash, ShouldNotBeEmpty)
		})
	})
}

func TestGetProcesses(t *testing.T) {
	Convey("Function getProcesses", t, func() {
		Convey("Should load up the processes map", func() {

			err := initialize()
			So(err, ShouldBeNil)

			// get a db connection
			err1 := setupDatabase()
			So(err1, ShouldBeNil)
			defer db.Close()

			processes, err3 := getProcesses()
			So(err3, ShouldBeNil)
			So(processes, ShouldNotBeEmpty)
		})
	})
}

func TestGetActions(t *testing.T) {
	Convey("Function getActions", t, func() {
		Convey("Should load up the action map", func() {

			err := initialize()
			So(err, ShouldBeNil)

			// get a db connection
			err1 := setupDatabase()
			So(err1, ShouldBeNil)
			defer db.Close()

			actions, err3 := getActions(10)
			So(err3, ShouldBeNil)
			So(actions, ShouldNotBeEmpty)
		})
	})
}
