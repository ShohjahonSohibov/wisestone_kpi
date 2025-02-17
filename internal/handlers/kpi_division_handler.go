package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"kpi/internal/models"
	"kpi/internal/services"
)

type KPIDivisionHandler struct {
	kpiDivisionService *services.KPIDivisionService
}

func NewKPIDivisionHandler(kpiDivisionService *services.KPIDivisionService) *KPIDivisionHandler {
	return &KPIDivisionHandler{
		kpiDivisionService: kpiDivisionService,
	}
}

// @Security ApiKeyAuth
// @Summary Create KPI Division
// @Description Create a new KPI Division
// @Tags KPI Divisions
// @Accept json
// @Produce json
// @Param request body models.CreateKPIDivision true "KPI Division creation request"
// @Success 201 {object} map[string]interface{} "status: 201, message: KPI Division created successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-divisions [post]
func (h *KPIDivisionHandler) Create(c *gin.Context) {
	var req models.KPIDivision
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err := h.kpiDivisionService.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "kpi division created successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Update KPI Division
// @Description Update an existing KPI Division
// @Tags KPI Divisions
// @Accept json
// @Produce json
// @Param id path string true "KPI Division ID"
// @Param request body models.UpdateKPIDivision true "KPI Division update request"
// @Success 200 {object} map[string]interface{} "status: 200, message: KPI Division updated successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-divisions/{id} [put]
func (h *KPIDivisionHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req models.KPIDivision
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req.ID = id
	err := h.kpiDivisionService.Update(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "kpi division updated successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Delete KPI Division
// @Description Delete a KPI Division
// @Tags KPI Divisions
// @Accept json
// @Produce json
// @Param id path string true "KPI Division ID"
// @Success 200 {object} map[string]interface{} "status: 200, message: KPI Division deleted successfully"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-divisions/{id} [delete]
func (h *KPIDivisionHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.kpiDivisionService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "kpi division deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Get KPI Division by ID
// @Description Get a KPI Division by its ID
// @Tags KPI Divisions
// @Accept json
// @Produce json
// @Param id path string true "KPI Division ID"
// @Success 200 {object} map[string]interface{} "status: 200, data: KPI Division object"
// @Failure 404 {object} map[string]interface{} "status: 404, message: kpi division not found"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-divisions/{id} [get]
func (h *KPIDivisionHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.kpiDivisionService.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": err.Error(),
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
// @Summary List KPI Divisions
// @Description Get a list of KPI Divisions with filtering and pagination
// @Tags KPI Divisions
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param multi_search query string false "Search in name fields"
// @Param sort_order query string false "Sort order (asc/desc)"
// @Param parent_id query string false "Filter by parent ID"
// @Success 200 {object} map[string]interface{} "status: 200, data: ListKPIDivisionResponse"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-divisions [get]
func (h *KPIDivisionHandler) List(c *gin.Context) {
	var filter models.ListKPIDivisionRequest

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
	filter.ParentID = c.Query("parent_id")

	result, err := h.kpiDivisionService.List(c.Request.Context(), &filter)
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
