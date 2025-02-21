package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"kpi/config"
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

// @Summary Create KPI Progress
// @Description Create a new KPI progress record
// @Tags KPI Progress
// @Accept json
// @Produce json
// @Param request body models.CreateKPIProgress true "KPI Progress details"
// @Success 201 {object} models.KPIProgress
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /kpi-progresses [post]
func (h *KPIProgressHandler) Create(c *gin.Context) {
	var req models.CreateKPIProgress
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	progress := &models.KPIProgress{
		FactorId:          req.FactorId,
		FactorIndicatorId: req.FactorIndicatorId,
		Ratio:             req.Ratio,
		Date:              req.Date,
	}

	if err := h.kpiProgressService.Create(c, progress); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Progress created successfully",
		"data":    progress,
	})
}

// @Summary Delete KPI Progress
// @Description Delete a KPI progress record
// @Tags KPI Progress
// @Produce json
// @Param id path string true "Progress ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /kpi-progresses/{id} [delete]
func (h *KPIProgressHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.kpiProgressService.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Progress deleted successfully"})
}

// @Summary List KPI Progress
// @Description Get a list of KPI progress records with pagination and filtering
// @Tags KPI Progress
// @Produce json
// @Param date query string false "Filter by date"
// @Param user_id query string false "Filter by user ID"
// @Param team_id query string false "Filter by team ID"
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {object} models.ListKPIProgressResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /kpi-progresses [get]
func (h *KPIProgressHandler) List(c *gin.Context) {
	var filter models.ListKPIProgressRequest

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

	// Parse date if provided
	dateStr := c.Query("date")
	if dateStr != "" {
		parsedDate, err := time.Parse(config.TimeFormat, dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid date format. Expected YYYY-MM-DD",
			})
			return
		}
		filter.Date = parsedDate
	}

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

// // @Summary Get KPI Progress by ID
// // @Description Get a KPI progress record by its ID
// // @Tags KPI Progress
// // @Produce json
// // @Param id path string true "Progress ID"
// // @Success 200 {object} models.KPIProgress
// // @Failure 500 {object} map[string]string
// // @Router /kpi-progresses/{id} [get]
// func (h *KPIProgressHandler) GetByID(c *gin.Context) {
// 	id := c.Param("id")
// 	progress, err := h.kpiProgressService.GetByID(c, id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, progress)
// }

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
// // @Router /kpi-progresses/{id} [put]
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
