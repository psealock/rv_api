package rvService

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var complaint_tables = map[string]string{
	"bk": "service_request_311_bk",
	"mn": "service_request_311_mn",
	"qn": "service_request_311_qn",
	"bx": "service_request_311_bx",
	"si": "service_request_311_si",
}

func panic(err string) {
	fmt.Println(err)
}

func GetComplaints(address, borough string) []map[string]interface{} {
	db, err := sql.Open("mysql", "rv:12345678@tcp(107.170.173.75:3306)/rv")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Execute the query
	table := complaint_tables[borough]
	key := "IncidentAddress"
	tableColumns := "*"

	query := "SELECT " + tableColumns + " FROM " + table + " WHERE " + key + " = \"" + address + "\""
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var results []map[string]interface{}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		//@TODO return proper types, not just strings.
		record := make(map[string]interface{})
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			record[columns[i]] = value
		}
		results = append(results, record)
	}
	return results
}
