package routes

import (
	"datagps/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handlers.Handler) {
	r.POST("/data", h.Data)
	r.POST("/pendingGroups", h.PendingGroups)
	r.POST("/processPendingGroups", h.ProcessPendingGroups)
	r.POST("/execSetGroup", h.ExecSetGroup)
	r.POST("/groups", h.Groups)
}
