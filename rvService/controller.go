package rvService

import (
	"github.com/emicklei/go-restful"
)

type Building struct {
	Address, Borough, Status string
}

func FindBuilding(request *restful.Request, response *restful.Response) {
	Console("fetching results...")
	address := request.PathParameter("address")
	borough := request.PathParameter("borough")

	// here you would fetch user from some persistence system
	GetComplaints(address, borough)

	bldg := &Building{Address: address, Borough: borough, Status: "Ruhl Deece"}
	response.WriteEntity(bldg)
}
