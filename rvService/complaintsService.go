package rvService

var complaint_tables = map[string]string{
	"bk": "service_request_311_bk",
	"mn": "service_request_311_mn",
	"qn": "service_request_311_qn",
	"bx": "service_request_311_bx",
	"si": "service_request_311_si",
}

var pluto_tables = map[string]string{
	"bk": "pluto_bk",
	"mn": "pluto_mn",
	"qn": "pluto_qn",
	"bx": "pluto_bx",
	"si": "pluto_si",
}

func GetServiceRequest(address, borough string) []map[string]interface{} {

	// Execute the query
	table := complaint_tables[borough]
	key := "IncidentAddress"
	tableColumns := "*"

	query := "SELECT " + tableColumns + " FROM " + table + " WHERE " + key + " = \"" + address + "\""

	return DB(query)
}

func GetDobComplaint(bin string) []map[string]interface{} {
	table := "dob_complaints_all"
	key := "BIN"
	tableColumns := "*"

	query := "SELECT " + tableColumns + " FROM " + table + " WHERE " + key + " = \"" + bin + "\""

	return DB(query)
}

func GetPluto(bbl, borough string) []map[string]interface{} {
	table := pluto_tables[borough]
	key := "BBL"
	tableColumns := "*"

	query := "SELECT " + tableColumns + " FROM " + table + " WHERE " + key + " = \"" + bbl + "\""

	return DB(query)
}

func GetAll(address, borough, bin, bbl string) map[string]interface{} {
	//@TODO: Make these queries run concurrently with goroutines
	record := make(map[string]interface{})

	record["serviceRequests"] = GetServiceRequest(address, borough)
	record["dobComplaints"] = GetDobComplaint(bin)
	record["pluto"] = GetPluto(bbl, borough)

	return record
}
