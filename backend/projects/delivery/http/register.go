package http

import (
	Projects "nik19ta/backend/projects"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc Projects.UseCase) {
	h := NewHandler(uc)

	ProjectsEndpoints := router.Group("/api/projects")
	{
		ProjectsEndpoints.GET("", h.GetProjects)
		ProjectsEndpoints.GET("/:id", h.GetProject)
	}
}
