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
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sapk/go-genesys/db"
	"github.com/sapk/rtme-browser/router/api/common"
)

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
		//   "200":
		//     "description": Place info
		c.JSON(http.StatusNotImplemented, common.Error{Error: "Not Implemented"})
	})
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
		//   "200":
		//     "description": Status info
		c.JSON(http.StatusNotImplemented, common.Error{Error: "Not Implemented"})
	})
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
		//   "200":
		//     "description": Logins info
		c.JSON(http.StatusNotImplemented, common.Error{Error: "Not Implemented"})
	})
}
