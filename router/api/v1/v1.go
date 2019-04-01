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
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sapk/go-genesys/db/rtme"
	"github.com/sapk/rtme-browser/module/config"
	"github.com/sapk/rtme-browser/router/api/common"
)

//TODO manage db lock in go-genesys

//SetupRouter configure the listener
func SetupRouter(r *gin.RouterGroup) {
	r.GET("/app/status", func(c *gin.Context) {
		// swagger:operation GET /app/status app appStatus
		// ---
		// summary: Get app status
		// produces:
		// - application/json
		// responses:
		//   "200":
		//     "description": status response

		if config.IsNotInit() {
			c.JSON(200, map[string]string{
				"status": "init",
			})
			return
		}
		c.JSON(200, map[string]string{
			"status": "ready",
		})
	})

	r.POST("/app/config", func(c *gin.Context) {
		// swagger:operation POST /app/config app appConfig
		// ---
		// summary: Setup app config
		// produces:
		// - application/json
		// parameters:
		// - name: DBTYPE
		//   in: body
		//   description: db type
		//   type: string
		//   required: true
		// - name: DBRTMEURL
		//   in: body
		//   description: db rtme url
		//   type: string
		//   required: true
		// - name: DBCFGURL
		//   in: body
		//   description: db cfg url
		//   type: string
		//   required: true
		// responses:
		//   "400":
		//     "$ref": "#/responses/BadRequest"
		//   "500":
		//     "$ref": "#/responses/InternalServerError"
		//   "204":
		//     "description": No Content

		var configObj struct {
			DBTYPE    string `form:"DBTYPE" json:"DBTYPE"`
			DBRTMEURL string `form:"DBRTMEURL" json:"DBRTMEURL"`
			DBCFGURL  string `form:"DBCFGURL" json:"DBCFGURL"`
		}
		err := c.Bind(&configObj)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}
		err = config.Setup(configObj.DBTYPE, configObj.DBRTMEURL, configObj.DBCFGURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}
		if config.IsNotInit() {
			c.JSON(200, map[string]string{
				"status": "init",
			})
			return
		}
		c.JSON(200, map[string]string{
			"status": "ready",
		})
	})

	r.GET("/cfg/groups", func(c *gin.Context) {
		// swagger:operation GET /cfg/groups config cfgGroups
		// ---
		// summary: Get group name list
		// produces:
		// - application/json
		// responses:
		//   "400":
		//     "$ref": "#/responses/BadRequest"
		//   "500":
		//     "$ref": "#/responses/InternalServerError"
		//   "200":
		//     "description": Groups names

		//c.JSON(http.StatusNotImplemented, common.Error{Error: "Not Implemented"})
		//TODO move to go-genesys
		cfg, err := config.Get()
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}
		results, err := cfg.CFG.Engine.Query("SELECT DISTINCT name FROM dbo.cfg_group")
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}
		list := make([]string, len(results))
		for i, r := range results {
			list[i] = strings.TrimSpace(string(r["name"]))
		}
		c.JSON(200, list)
	})

	r.GET("/cfg/groups/:name/users", func(c *gin.Context) {
		// swagger:operation GET /cfg/groups/{name}/users config cfgGroups
		// ---
		// summary: Get group agent list
		// produces:
		// - application/json
		// parameters:
		// - name: name
		//   in: path
		//   description: name of group
		//   type: string
		//   required: true
		// responses:
		//   "400":
		//     "$ref": "#/responses/BadRequest"
		//   "500":
		//     "$ref": "#/responses/InternalServerError"
		//   "200":
		//     "description": Agent dbid list

		//c.JSON(http.StatusNotImplemented, common.Error{Error: "Not Implemented"})
		//TODO move to go-genesys
		cfg, err := config.Get()
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}
		results, err := cfg.CFG.Engine.Query("SELECT agent_dbid FROM dbo.cfg_agent_group INNER JOIN dbo.cfg_group ON cfg_group.dbid = cfg_agent_group.group_dbid WHERE cfg_group.name = ?", c.Param("name"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}
		list := make([]string, len(results))
		for i, r := range results {
			list[i] = strings.TrimSpace(string(r["agent_dbid"]))
		}
		c.JSON(200, list)
	})

	r.GET("/cfg/places/:dbid", func(c *gin.Context) {
		// swagger:operation GET /cfg/places/{dbid} config cfgPlace
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
		cfg, err := config.Get()
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}
		results, err := cfg.CFG.Engine.Query("SELECT dbid, name, state, csid, tenant_csid, capacity_dbid, site_dbid, contract_dbid FROM dbo.cfg_place WHERE dbid = ?", c.Param("dbid"))
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
		// swagger:operation GET /rtme/{date}/status rtme rtmeStatus
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
		cfg, err := config.Get()
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}
		sessions, err := rtme.FormattedStatusEntriesOfDay(cfg.RTME, c.Param("date"))
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
		// swagger:operation GET /rtme/{date}/login rtme rtmeLogin
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

		cfg, err := config.Get()
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}
		sessions, err := rtme.FormattedLoginEntriesOfDay(cfg.RTME, c.Param("date"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
			return
		}

		//TODO filter by group

		c.JSON(http.StatusOK, sessions)
	})
}
