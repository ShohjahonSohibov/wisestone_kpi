package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"kpi/internal/models"
	"kpi/internal/services"
)

type KPICriterionHandler struct {
	kpiCriterionService *services.KPICriterionService
}

func NewKPICriterionHandler(kpiCriterionService *services.KPICriterionService) *KPICriterionHandler {
	return &KPICriterionHandler{
		kpiCriterionService: kpiCriterionService,
	}
}

// @Security ApiKeyAuth
// @Summary Create KPI Criterion
// @Description Create a new KPI Criterion
// @Tags KPI Criterions
// @Accept json
// @Produce json
// @Param request body models.CreateKPICriterion true "KPI Criterion creation request"
// @Success 201 {object} map[string]interface{} "status: 201, message: KPI Criterion created successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-criterions [post]
func (h *KPICriterionHandler) Create(c *gin.Context) {
	var req models.KPICriterion
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err := h.kpiCriterionService.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "kpi criterion created successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Update KPI Criterion
// @Description Update an existing KPI Criterion
// @Tags KPI Criterions
// @Accept json
// @Produce json
// @Param id path string true "KPI Criterion ID"
// @Param request body models.UpdateKPICriterion true "KPI Criterion update request"
// @Success 200 {object} map[string]interface{} "status: 200, message: KPI Criterion updated successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-criterions/{id} [put]
func (h *KPICriterionHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req models.KPICriterion
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req.ID = id
	err := h.kpiCriterionService.Update(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "kpi criterion updated successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Delete KPI Criterion
// @Description Delete a KPI Criterion
// @Tags KPI Criterions
// @Accept json
// @Produce json
// @Param id path string true "KPI Criterion ID"
// @Success 200 {object} map[string]interface{} "status: 200, message: KPI Criterion deleted successfully"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-criterions/{id} [delete]
func (h *KPICriterionHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.kpiCriterionService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "kpi criterion deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Get KPI Criterion by ID
// @Description Get a KPI Criterion by its ID
// @Tags KPI Criterions
// @Accept json
// @Produce json
// @Param id path string true "KPI Criterion ID"
// @Success 200 {object} map[string]interface{} "status: 200, data: KPI Criterion object"
// @Failure 404 {object} map[string]interface{} "status: 404, message: kpi criterion not found"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-criterions/{id} [get]
func (h *KPICriterionHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.kpiCriterionService.GetByID(c.Request.Context(), id)
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
// @Summary List KPI Criterions
// @Description Get a list of KPI Criterions with filtering and pagination
// @Tags KPI Criterions
// @Accept json
// @Produce json
// @Param filter query models.ListKPICriterionRequest false "Filter parameters"
// @Success 200 {object} map[string]interface{} "status: 200, data: ListKPICriterionResponse"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-criterions [get]
func (h *KPICriterionHandler) List(c *gin.Context) {
	var filter models.ListKPICriterionRequest

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
	filter.DivisionID = c.Query("division_id")

	result, err := h.kpiCriterionService.List(c.Request.Context(), &filter)
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
