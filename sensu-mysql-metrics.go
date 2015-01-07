// Go Sensu Mysql metrics plugin
//
// Copyright 2015 Marek Knappe <marek.knappe (..at..) gmail.com> . All rights reserved.
//
//
// 2015-01-07 1.0 First version
//
//
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package main


import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"flag"
	"os"
	"time"
)

var (
	username = flag.String("username", "root", "MySQL Username")
	password = flag.String("password", "root", "MySQL Password")
	host = flag.String("host", "localhost", "MySQL Server")
	port = flag.Int("port", 3306, "MySQL Port")
	timeout = flag.String("timeout", "10s", "MySQL connection timeout")
	mysqlname = flag.String("mysqlname","localhost", "MySQL name for the graphite output");
)

func init() {
	flag.Parse()
}

func main() {
	var mvar string
	var mvalue string
	var hostname string
	hostname, err := os.Hostname();
	flag.Parse()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/?timeout=%s", *username, *password, *host, *port, *timeout))
	if err != nil {	
		log.Fatal(err)
	}
	defer db.Close()
        
	Rows, err := db.Query("SHOW GLOBAL STATUS"); 	
        for Rows.Next() {
        	err = Rows.Scan(&mvar,&mvalue)
        	if err != nil {
			log.Fatal(err)
            		return
        	}
        	fmt.Printf("%s.mysql.%s.%s %s %d\n",hostname,*mysqlname,mvar,mvalue,time.Now().Unix())
    	}	
}
