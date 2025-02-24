package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kpi/internal/models"
	"kpi/internal/services"
)

type KPIProgressStatusHandler struct {
	service *services.KPIProgressStatusService
}

func NewKPIProgressStatusHandler(service *services.KPIProgressStatusService) *KPIProgressStatusHandler {
	return &KPIProgressStatusHandler{service: service}
}

// @Security ApiKeyAuth
// Create godoc
// @Summary Create KPI progress status
// @Description Create a new KPI progress status
// @Tags KPI Progress Status
// @Accept json
// @Produce json
// @Param request body models.CreateKPIProgressStatus true "Create KPI progress status request"
// @Success 201 {object} models.KPIProgressStatus
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/kpi-progress-status [post]
func (h *KPIProgressStatusHandler) Create(c *gin.Context) {
	var req models.CreateKPIProgressStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Progress created successfully",
	})
}

// @Security ApiKeyAuth
// Update godoc
// @Summary Update KPI progress status
// @Description Update an existing KPI progress status
// @Tags KPI Progress Status
// @Accept json
// @Produce json
// @Param id path string true "KPI Progress Status ID"
// @Param request body models.UpdateKPIProgressStatus true "Update KPI progress status request"
// @Success 200 {object} models.KPIProgressStatus
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/kpi-progress-status/{id} [put]
func (h *KPIProgressStatusHandler) Update(c *gin.Context) {
	var req models.UpdateKPIProgressStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req.ID = c.Param("id")
	if req.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID is required",
		})
		return
	}

	err := h.service.Update(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   "status updated successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Delete KPI Progress Status
// @Description Delete a KPI progress status record
// @Tags KPI Progress Status
// @Produce json
// @Param id path string true "Progress ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/kpi-progress-status/{id} [delete]
func (h *KPIProgressStatusHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"status":  http.StatusNoContent,
		"message": "Status deleted successfully",
	})
}

// @Security ApiKeyAuth
// List godoc
// @Summary List KPI progress statuses
// @Description Get a list of KPI progress statuses with optional filtering
// @Tags KPI Progress Status
// @Accept json
// @Produce json
// @Param team_id query string false "Team ID"
// @Param employee_id query string false "Employee ID"
// @Param date query string false "Date (YYYY-MM)"
// @Param status query string false "Status"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} models.ListKPIProgressStatusResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/kpi-progress-status [get]
func (h *KPIProgressStatusHandler) List(c *gin.Context) {
	var filter models.ListKPIProgressStatusRequest

	offset, limit, err := getPageOffsetLimit(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	filter.Offset = offset
	filter.Limit = limit
	filter.TeamId = c.Query("team_id")
	filter.EmployeeId = c.Query("employee_id")
	filter.Date = c.Query("date")
	filter.Status = c.Query("status")

	response, err := h.service.List(c.Request.Context(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   response,
	})
}
