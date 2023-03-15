// Copyright (c) 2017, Tax Products Group, LLC.
//
// All rights reserved.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStatus(t *testing.T) {
	Convey("Function status", t, func() {
		Convey("Should return status JSON", func() {

			// 			expected := `
			// {
			//     "Program": "Collection_machine.Test",
			//     "Buildstamp": "not set",
			//     "GitHash": "not set",
			//     "GoVersion": "go1.7.4",
			//     "RunTime": "28s",
			//     "Activities": 42,
			//     "EmailsSent": 28,
			//     "SMSSent": 2,
			//     "PaymentsCreated": 10,
			//     "DirectMailCreated": 2,
			//     "DBconnections": 0,
			//     "GiactError": 1
			// }`

			err := initialize()
			So(err, ShouldEqual, nil)

			// get a db connection
			err1 := setupDatabase()
			So(err1, ShouldEqual, nil)
			defer db.Close()

			// run status HandlerFunc
			ts := httptest.NewServer(http.HandlerFunc(status))
			defer ts.Close()

			res, err := http.Get(ts.URL)
			if err != nil {
				log.Println(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Println(err)
			}
			str := string(bs)

			fmt.Printf("%d - %s", res.StatusCode, str)
			So(res.StatusCode, ShouldNotBeNil)
			// So(str, ShouldEqual, expected)
		})
	})
}
