package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	httpv1 "github.com/pprishchepa/go-invitecoder-example/internal/controller/http/v1"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(invites *httpv1.InvitesRoutes) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	e := gin.New()
	e.ContextWithFallback = true
	e.HandleMethodNotAllowed = true

	e.Use(gin.Recovery())

	e.GET("/", func(c *gin.Context) { c.Status(http.StatusNoContent) })

	// Go metrics.
	e.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// For k8s readiness and liveness probes.
	e.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	// For k8s lifecycle.preStop.
	// For clusters with slow ingress resource propagation.
	e.GET("/sleep", func(c *gin.Context) {
		time.Sleep(20 * time.Second)
		c.Status(http.StatusNoContent)
	})

	v1 := e.Group("/v1")
	{
		invites.RegisterRoutes(v1)
	}

	return e
}
