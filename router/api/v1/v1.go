//package v1 V1 API.
//
// This documentation describes the RTME/CFG API.
//
//     Schemes: http
//     BasePath: /api/v1
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sapk/go-genesys/db"
	"github.com/sapk/go-genesys/db/rtme"
	"github.com/sapk/rtme-browser/router/api/common"
)

//TODO manage db lock in go-genesys

//SetupRouter configure the listener
func SetupRouter(r *gin.RouterGroup, rtmeDB, cfgDB *db.DB) {
	r.GET("/cfg/places/:dbid", func(c *gin.Context) {
		// swagger:operation GET /cfg/places/{dbid} site cfgPlace
		// ---
		// summary: Get info on a place
		// produces:
		// - application/json
		// parameters:
		// - name: dbid
		//   in: path
		//   description: id of place
		//   type: string
		//   required: true
		// responses:
		//   "400":
		//     "$ref": "#/responses/BadRequest"
		//   "500":
		//     "$ref": "#/responses/InternalServerError"
		//   "200":
		//     "description": Place info

		//c.JSON(http.StatusNotImplemented, common.Error{Error: "Not Implemented"})
		//TODO move to go-genesys
		results, err := cfgDB.Engine.Query(fmt.Sprintf("SELECT dbid, name, state, csid, tenant_csid, capacity_dbid, site_dbid, contract_dbid FROM dbo.cfg_place WHERE dbid = '%s'", c.Param("dbid")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}
		if len(results) == 0 {
			c.JSON(http.StatusNotFound, common.Error{Error: "Place not found"})
			return
		}

		c.JSON(200, map[string]string{
			"name": string(results[0]["name"]),
		})
	})

	// RTMEStatusResponse is a RTMEStatusResponse format response
	// swagger:response RTMEStatusResponse
	type RTMEStatusResponse struct {
		// The response
		// in: body
		Body rtme.FormattedStatusResp
	}

	r.GET("rtme/:date/status", func(c *gin.Context) {
		// swagger:operation GET /rtme/{date}/status site rtmeStatus
		// ---
		// summary: Export rtme info of status.
		// produces:
		// - application/json
		// parameters:
		// - name: date
		//   in: path
		//   description: date
		//   type: string
		//   required: true
		// - name: group
		//   in: query
		//   description: filter by group/instance name
		//   type: string
		// responses:
		//   "400":
		//     "$ref": "#/responses/BadRequest"
		//   "500":
		//     "$ref": "#/responses/InternalServerError"
		//   "200":
		//     "$ref": "#/responses/RTMEStatusResponse"
		//     "description": Status info

		//c.JSON(http.StatusNotImplemented, common.Error{Error: "Not Implemented"})
		sessions, err := rtme.FormattedStatusEntriesOfDay(rtmeDB, c.Param("date"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}

		//TODO filter by group

		c.JSON(http.StatusOK, sessions)
	})

	// RTMELoginResponse is a RTMELoginResponse format response
	// swagger:response RTMELoginResponse
	type RTMELoginResponse struct {
		// The response
		// in: body
		Body rtme.FormattedLoginResp
	}

	r.GET("rtme/:date/login", func(c *gin.Context) {
		// swagger:operation GET /rtme/{date}/login site rtmeLogin
		// ---
		// summary: Export rtme info of login.
		// produces:
		// - application/json
		// parameters:
		// - name: date
		//   in: path
		//   description: date
		//   type: string
		//   required: true
		// - name: group
		//   in: query
		//   description: filter by group/instance name
		//   type: string
		// responses:
		//   "400":
		//     "$ref": "#/responses/BadRequest"
		//   "500":
		//     "$ref": "#/responses/InternalServerError"
		//   "200":
		//     "$ref": "#/responses/RTMELoginResponse"
		//     "description": Logins info

		sessions, err := rtme.FormattedLoginEntriesOfDay(rtmeDB, c.Param("date"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}

		//TODO filter by group

		c.JSON(http.StatusOK, sessions)
	})
}
