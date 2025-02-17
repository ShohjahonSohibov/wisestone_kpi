package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"kpi/internal/models"
	"kpi/internal/services"
)

type KPIFactorIndicatorHandler struct {
	kpiFactorIndicatorService *services.KPIFactorIndicatorService
}

func NewKPIFactorIndicatorHandler(kpiFactorIndicatorService *services.KPIFactorIndicatorService) *KPIFactorIndicatorHandler {
	return &KPIFactorIndicatorHandler{
		kpiFactorIndicatorService: kpiFactorIndicatorService,
	}
}

// @Security ApiKeyAuth
// @Summary Create KPI Factor Indicator
// @Description Create a new KPI Factor Indicator
// @Tags KPI Factor Indicators
// @Accept json
// @Produce json
// @Param request body models.CreateKPIFactorIndicator true "KPI Factor Indicator creation request"
// @Success 201 {object} map[string]interface{} "status: 201, message: KPI Factor Indicator created successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-factor-indicators [post]
func (h *KPIFactorIndicatorHandler) Create(c *gin.Context) {
	var req models.KPIFactorIndicator
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err := h.kpiFactorIndicatorService.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "kpi factor indicator created successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Update KPI Factor Indicator
// @Description Update an existing KPI Factor Indicator
// @Tags KPI Factor Indicators
// @Accept json
// @Produce json
// @Param id path string true "KPI Factor Indicator ID"
// @Param request body models.UpdateKPIFactorIndicator true "KPI Factor Indicator update request"
// @Success 200 {object} map[string]interface{} "status: 200, message: KPI Factor Indicator updated successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-factor-indicators/{id} [put]
func (h *KPIFactorIndicatorHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req models.KPIFactorIndicator
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req.ID = id
	err := h.kpiFactorIndicatorService.Update(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "kpi factor indicator updated successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Delete KPI Factor Indicator
// @Description Delete a KPI Factor Indicator
// @Tags KPI Factor Indicators
// @Accept json
// @Produce json
// @Param id path string true "KPI Factor Indicator ID"
// @Success 200 {object} map[string]interface{} "status: 200, message: KPI Factor Indicator deleted successfully"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-factor-indicators/{id} [delete]
func (h *KPIFactorIndicatorHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.kpiFactorIndicatorService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "kpi factor indicator deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Get KPI Factor Indicator by ID
// @Description Get a KPI Factor Indicator by its ID
// @Tags KPI Factor Indicators
// @Accept json
// @Produce json
// @Param id path string true "KPI Factor Indicator ID"
// @Success 200 {object} map[string]interface{} "status: 200, data: KPI Factor Indicator object"
// @Failure 404 {object} map[string]interface{} "status: 404, message: kpi factor indicator not found"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-factor-indicators/{id} [get]
func (h *KPIFactorIndicatorHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.kpiFactorIndicatorService.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "kpi factor indicator not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   result,
	})
}

// @Security ApiKeyAuth
// @Summary List KPI Factor Indicators
// @Description Get a list of KPI Factor Indicators with filtering and pagination
// @Tags KPI Factor Indicators
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param multi_search query string false "Search in name fields"
// @Param sort_order query string false "Sort order (asc/desc)"
// @Param factor_id query string false "Filter by factor ID"
// @Success 200 {object} map[string]interface{} "status: 200, data: ListKPIFactorIndicatorResponse"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-factor-indicators [get]
func (h *KPIFactorIndicatorHandler) List(c *gin.Context) {
	var filter models.ListKPIFactorIndicatorRequest

	offset, limit, err := getPageOffsetLimit(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	filter.Filter.Offset = offset
	filter.Filter.Limit = limit
	filter.MultiSearch = c.Query("multi_search")
	filter.SortOrder = c.Query("sort_order")
	filter.FactorID = c.Query("factor_id")

	result, err := h.kpiFactorIndicatorService.List(c.Request.Context(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   result,
	})
}