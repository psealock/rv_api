package rvService

import (
	"github.com/emicklei/go-restful"
)

func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/rv").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	service.Route(service.GET("311/{address}/{borough}").To(HandleServiceRequest))

	service.Route(service.GET("dob/{bin}").To(HandleDobComplaint))

	service.Route(service.GET("pluto/{bbl}/{borough}").To(HandlePluto))

	service.Route(service.GET("all/{address}/{borough}/{bin}/{bbl}").To(HandleAll))

	return service
}
