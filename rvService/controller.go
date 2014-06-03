package rvService

import (
	"github.com/emicklei/go-restful"
)

func HandleServiceRequest(request *restful.Request, response *restful.Response) {
	Console("fetching 311...")
	address := request.PathParameter("address")
	borough := request.PathParameter("borough")

	//grab array of service requests.
	cmplnt := GetServiceRequest(address, borough)
	response.WriteEntity(cmplnt)
}

func HandleDobComplaint(request *restful.Request, response *restful.Response) {
	Console("fetching DoB...")
	bin := request.PathParameter("bin")

	//grab array of dob complaints
	dob := GetDobComplaint(bin)
	response.WriteEntity(dob)
}

func HandlePluto(request *restful.Request, response *restful.Response) {
	Console("fetching Pluto...")

	bbl := request.PathParameter("bbl")
	borough := request.PathParameter("borough")

	//grab pluto
	pluto := GetPluto(bbl, borough)
	response.WriteEntity(pluto)
}

func HandleAll(request *restful.Request, response *restful.Response) {
	Console("fetching All...")

	address := request.PathParameter("address")
	borough := request.PathParameter("borough")
	bin := request.PathParameter("bin")
	bbl := request.PathParameter("bbl")

	//grab all
	all := GetAll(address, borough, bin, bbl)
	response.WriteEntity(all)
}
