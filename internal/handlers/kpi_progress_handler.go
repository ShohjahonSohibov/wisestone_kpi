package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kpi/internal/models"
	"kpi/internal/services"
)

type KPIProgressHandler struct {
	kpiProgressService *services.KPIProgressService
}

func NewKPIProgressHandler(kpiProgressService *services.KPIProgressService) *KPIProgressHandler {
	return &KPIProgressHandler{
		kpiProgressService: kpiProgressService,
	}
}

// CreateMany godoc
// @Security ApiKeyAuth
// @Summary Create multiple KPI progresses
// @Description Create multiple KPI progress records at once
// @Tags KPI Progress
// @Accept json
// @Produce json
// @Param request body []models.CreateKPIProgress true "Array of KPI progress creation requests"
// @Success 201 {object} map[string]interface{} "status: 201, message: KPI progresses created successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-progresses/bulk [post]
func (h *KPIProgressHandler) CreateMany(c *gin.Context) {
	var progresses []*models.CreateBulkKPIProgress
	if err := c.ShouldBindJSON(&progresses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if len(progresses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "no progress records provided",
		})
		return
	}

	if err := h.kpiProgressService.CreateMany(c.Request.Context(), progresses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "kpi progresses created successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Create KPI Progress
// @Description Create a new KPI progress record
// @Tags KPI Progress
// @Accept json
// @Produce json
// @Param request body models.CreateKPIProgress true "KPI Progress details"
// @Success 201 {object} models.KPIProgress
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/kpi-progresses [post]
func (h *KPIProgressHandler) Create(c *gin.Context) {
	var req models.KPIProgress
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if err := h.kpiProgressService.Create(c, &req); err != nil {
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
// @Summary Delete KPI Progress
// @Description Delete a KPI progress record
// @Tags KPI Progress
// @Produce json
// @Param date query string false "Progress ID"
// @Param team_id query string false "Progress ID"
// @Param employee_id query string false "Progress ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/kpi-progresses/delete [delete]
func (h *KPIProgressHandler) Delete(c *gin.Context) {
	date := c.Query("date")
	teamId := c.Query("team_id")
	employeeId := c.Query("employee_id")

	if err := h.kpiProgressService.Delete(c, date, teamId, employeeId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusNoContent,
		"message": "Progress deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Summary List KPI Progress
// @Description Get a list of KPI progress records with pagination and filtering
// @Tags KPI Progress
// @Produce json
// @Param date query string false "Filter by date"
// @Param employee_id query string false "Filter by employee ID"
// @Param team_id query string false "Filter by team ID"
// @Success 200 {object} models.ListKPIProgressResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/kpi-progresses [get]
func (h *KPIProgressHandler) List(c *gin.Context) {
	var filter models.ListKPIProgressRequest

	filter.TeamId = c.Query("team_id")
	filter.EmployeeId = c.Query("employee_id")
	filter.Date = c.Query("date")

	response, err := h.kpiProgressService.List(c, &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Progress list retrieved successfully",
		"data":    response,
	})
}

// @Security ApiKeyAuth
// // @Summary Get KPI Progress by ID
// // @Description Get a KPI progress record by its ID
// // @Tags KPI Progress
// // @Produce json
// // @Param id path string true "Progress ID"
// // @Success 200 {object} models.KPIProgress
// // @Failure 500 {object} map[string]string
// // @Router /api/v1/kpi-progresses/{id} [get]
// func (h *KPIProgressHandler) GetByID(c *gin.Context) {
// 	id := c.Param("id")
// 	progress, err := h.kpiProgressService.GetByID(c, id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, progress)
// }

// @Security ApiKeyAuth
// // @Summary Update KPI Progress
// // @Description Update an existing KPI progress record
// // @Tags KPI Progress
// // @Accept json
// // @Produce json
// // @Param id path string true "Progress ID"
// // @Param request body models.UpdateKPIProgress true "KPI Progress details"
// // @Success 200 {object} models.KPIProgress
// // @Failure 400 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /api/v1/kpi-progresses/{id} [put]
// func (h *KPIProgressHandler) Update(c *gin.Context) {
// 	var req models.UpdateKPIProgress
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	req.ID = c.Param("id")
// 	progress := &models.KPIProgress{
// 		ID:                req.ID,
// 		FactorId:          req.FactorId,
// 		FactorIndicatorId: req.FactorIndicatorId,
// 		Ratio:             req.Ratio,
// 		Date:              req.Date,
// 	}

// 	if err := h.kpiProgressService.Update(c, progress); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, progress)
// }
