package http

import (
	"net/http"
	"nik19ta/backend/projects"
	Projects "nik19ta/backend/projects"

	gin "github.com/gin-gonic/gin"
)

type Handler struct {
	useCase Projects.UseCase
}

func NewHandler(useCase Projects.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) GetProjects(c *gin.Context) {
	lang := c.Query("lang")

	projects, err := h.useCase.GetProjects(lang)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

func (h *Handler) GetProject(c *gin.Context) {
	lang := c.Query("lang")
	id := c.Param("id")

	project, err := h.useCase.GetProject(lang, id)

	if err != nil {
		if err == projects.CouldNotFindProject {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}
