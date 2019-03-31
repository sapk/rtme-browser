package router

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/sapk/go-genesys/db"

	//"github.com/sapk/rtme-browser/public/swagger"
	//"github.com/sapk/rtme-browser/public/ui"
	v1 "github.com/sapk/rtme-browser/router/api/v1"
)

func generateRouter(rtmeDB, cfgDB *db.DB) *gin.Engine {
	r := gin.New()
	//Configure logger
	r.Use(logger())
	//Recovery
	r.Use(gin.Recovery())
	//Compress responses globally
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	api := r.Group("/api/v1")
	//API v1
	v1.SetupRouter(api, rtmeDB, cfgDB)

	//TODO handler for direct FS package gzip
	//r.StaticFS("/ui", ui.FS(false))
	//r.StaticFS("/swagger", swagger.FS(false))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/ui/")
	})
	return r
}
