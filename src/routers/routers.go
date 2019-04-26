package routers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// CreateRouter defines the router and sets the basic URL mappings.
func CreateRouter() *gin.Engine {
	// create the router
	router := gin.Default()
	var nrMiddleware gin.HandlerFunc
	if os.Getenv("GO_ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
		// nrMiddleware = metric.NewRelic(true)
	} else {
		// nrMiddleware = metric.NewRelic(false)
	}
	// router.Use(context.PropagateHeaders)
	router.Use(gin.Recovery())
	//router.Use(datadog.Handler)

	router.NoRoute(noRouteHandler)
	router.NoMethod(methodNotAllowedHandler)

	router.HandleMethodNotAllowed = true
	router.RedirectFixedPath = false
	router.RedirectTrailingSlash = false

	// ping
	router.GET("/ping", pingHandler)

	//router
	parentGroup := router.Group("/")
	if nrMiddleware != nil {
		parentGroup.Use(nrMiddleware)
	}
	// configure the whole set of URL/Handlers mappings
	configureMappings(parentGroup)

	return router
}

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func noRouteHandler(c *gin.Context) {
	err := struct {
		ErrorMessage string `json:"message"`
		ErrorCode    string `json:"error"`
		ErrorStatus  int    `json:"status"`
	}{
		ErrorMessage: "resource " + c.Request.URL.Path,
		ErrorCode:    "not_found",
		ErrorStatus:  http.StatusNotFound,
	}
	c.JSON(err.ErrorStatus, err)
}

func methodNotAllowedHandler(c *gin.Context) {
	err := struct {
		ErrorMessage string `json:"message"`
		ErrorCode    string `json:"error"`
		ErrorStatus  int    `json:"status"`
	}{
		ErrorMessage: "Method not allowed",
		ErrorCode:    "Method not allowed",
		ErrorStatus:  http.StatusMethodNotAllowed,
	}
	c.JSON(err.ErrorStatus, err)
}
