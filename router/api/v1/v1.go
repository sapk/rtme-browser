package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/sapk/go-genesys/db"
)

//SetupRouter configure the listener
func SetupRouter(r *gin.RouterGroup, rtmeDB, cfgDB *db.DB) {
	r.GET("/cfg/places/:dbid", func(c *gin.Context) {
	})
	r.GET("rtme/:date/status", func(c *gin.Context) {
	})
	r.GET("rtme/:date/login", func(c *gin.Context) {
	})
}
