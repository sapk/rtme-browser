package router

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/sapk/rtme-browser/public/swagger"
	"github.com/sapk/rtme-browser/public/ui"
	v1 "github.com/sapk/rtme-browser/router/api/v1"
)

func generateRouter(allowCORS bool) *gin.Engine {
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
	if allowCORS {
		//CORS
		log.Info().Msg("Allow CORS request")
		api.Use(cors)
		api.Use(options)
	}
	//API v1
	v1.SetupRouter(api)

	//TODO handler for direct FS package gzip
	r.StaticFS("/ui", ui.FS(false))
	r.StaticFS("/swagger", swagger.FS(false))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/ui/")
	})
	return r
}
func cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin")) //TODO Filter
	c.Header("Vary", "Origin")
}

func options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin")) //TODO Filter
		c.Header("Vary", "Origin")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}
