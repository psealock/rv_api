package rvService

import (
	"github.com/emicklei/go-restful"
)

func FindBuilding(request *restful.Request, response *restful.Response) {
	Console("fetching results...")
	address := request.PathParameter("address")
	borough := request.PathParameter("borough")

	//grab array of complaints.
	cmplnt := GetComplaints(address, borough)
	response.WriteEntity(cmplnt)
}
