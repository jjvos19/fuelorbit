package handlers

import (
	"datagps/internal/models"
	"datagps/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv *service.AppService
}

func NewHandler(srv *service.AppService) *Handler {
	return &Handler{srv}
}

func (h *Handler) Data(ctx *gin.Context) {
	var req DataRequest
	//log.Println("Se llama a DataRequest")
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("Fecha: %v", req.SendDate)

	sendDate, err := time.Parse("2006-01-02", req.SendDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("Fecha str->Time: %v", sendDate)

	data := models.Data{
		TankerId:      req.CisternaId,
		GpsCoordinate: models.Point{X: req.GpsCoordinate.Lat, Y: req.GpsCoordinate.Lon},
		Volume:        req.Volume,
		StateMotor:    req.StateMotor,
		HashDevice:    req.HashDevice,
		SendDate:      sendDate,
	}
	//log.Printf("Data (db): %v", data)
	err = h.srv.CrearteData(&data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar"})
		return
	}
	ctx.JSON(http.StatusCreated, DataResponse{Id: data.Id, GroupBlck: data.GroupBlck})
}

func (h *Handler) PendingGroups(ctx *gin.Context) {
	var req PendingGroupsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groups, err := h.srv.PendingGroups(req.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al recuperar los grupos"})
		return
	}

	var dtoGrp []Group
	for _, grp := range groups {
		ngrupo := Group{Id: grp.Id, HashGroup: grp.HashGroup, IdStart: grp.IdStart, IdFinish: grp.IdFinish, HashBlck: grp.HashBlkc}
		dtoGrp = append(dtoGrp, ngrupo)
	}

	ctx.JSON(http.StatusOK, dtoGrp)
}

func (h *Handler) ProcessPendingGroups(ctx *gin.Context) {
	var req ProcessPendingGroupsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groups, err := h.srv.ProcessPendingGroups(req.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al recuperar los grupos"})
		return
	}

	var dtoGrp []Group
	for _, grp := range groups {

		ngrupo := Group{Id: grp.Id, HashGroup: grp.HashGroup, IdStart: grp.IdStart, IdFinish: grp.IdFinish, HashBlck: grp.HashBlkc}
		dtoGrp = append(dtoGrp, ngrupo)
	}

	ctx.JSON(http.StatusOK, dtoGrp)
}

func (h *Handler) ExecSetGroup(ctx *gin.Context) {
	var req ExecSetGroupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.srv.ExecSetGroup()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando el grupo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Se genero un grupo"})

}

func (h *Handler) Groups(ctx *gin.Context) {
	var req GroupsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groups, total, err := h.srv.Groups(req.Page, req.TotalPages, req.Records)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo los grupos"})
		return
	}

	var groupsDto []Group
	for _, gr := range groups {
		groupsDto = append(groupsDto, Group{Id: gr.Id, HashGroup: gr.HashGroup, IdStart: gr.IdStart, IdFinish: gr.IdFinish, HashBlck: gr.HashBlkc})
	}

	ctx.JSON(http.StatusOK, GroupResponse{Groups: groupsDto, TotalPages: total})
}
