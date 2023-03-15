// Copyright (c) 2017, Tax Products Group, LLC.
//
// All rights reserved.

package main

import (
	"fmt"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInitialize(t *testing.T) {
	Convey("Initialize", t, func() {
		Convey("Should initialize our configuration", func() {
			fmt.Printf("\n\n")
			err := initialize()
			So(err, ShouldEqual, nil)
		})
	})
}

func TestSetupDatabase(t *testing.T) {
	Convey("Good configuration", t, func() {
		Convey("It can connect to a database", func() {

			// Connect to database
			err1 := setupDatabase()
			So(err1, ShouldEqual, nil)

			// Ping database
			err2 := db.Ping()
			So(err2, ShouldEqual, nil)
		})
	})

}
