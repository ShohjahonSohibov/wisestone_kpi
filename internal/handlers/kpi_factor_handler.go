package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"kpi/internal/models"
	"kpi/internal/services"
)

type KPIFactorHandler struct {
	kpiFactorService *services.KPIFactorService
}

func NewKPIFactorHandler(kpiFactorService *services.KPIFactorService) *KPIFactorHandler {
	return &KPIFactorHandler{
		kpiFactorService: kpiFactorService,
	}
}

// @Security ApiKeyAuth
// @Summary Create KPI Factor
// @Description Create a new KPI Factor
// @Tags KPI Factors
// @Accept json
// @Produce json
// @Param request body models.CreateKPIFactor true "KPI Factor creation request"
// @Success 201 {object} map[string]interface{} "status: 201, message: KPI Factor created successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-factors [post]
func (h *KPIFactorHandler) Create(c *gin.Context) {
	var req models.KPIFactor
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err := h.kpiFactorService.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "kpi factor created successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Update KPI Factor
// @Description Update an existing KPI Factor
// @Tags KPI Factors
// @Accept json
// @Produce json
// @Param id path string true "KPI Factor ID"
// @Param request body models.UpdateKPIFactor true "KPI Factor update request"
// @Success 200 {object} map[string]interface{} "status: 200, message: KPI Factor updated successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-factors/{id} [put]
func (h *KPIFactorHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req models.KPIFactor
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req.ID = id
	err := h.kpiFactorService.Update(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "kpi factor updated successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Delete KPI Factor
// @Description Delete a KPI Factor
// @Tags KPI Factors
// @Accept json
// @Produce json
// @Param id path string true "KPI Factor ID"
// @Success 200 {object} map[string]interface{} "status: 200, message: KPI Factor deleted successfully"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-factors/{id} [delete]
func (h *KPIFactorHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.kpiFactorService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "kpi factor deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Get KPI Factor by ID
// @Description Get a KPI Factor by its ID
// @Tags KPI Factors
// @Accept json
// @Produce json
// @Param id path string true "KPI Factor ID"
// @Success 200 {object} map[string]interface{} "status: 200, data: KPI Factor object"
// @Failure 404 {object} map[string]interface{} "status: 404, message: kpi factor not found"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-factors/{id} [get]
func (h *KPIFactorHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.kpiFactorService.GetByID(c.Request.Context(), id)
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
// @Summary List KPI Factors
// @Description Get a list of KPI Factors with filtering and pagination
// @Tags KPI Factors
// @Accept json
// @Produce json
// @Param filter query models.ListKPIFactorRequest false "Filter parameters"
// @Success 200 {object} map[string]interface{} "status: 200, data: ListKPIFactorResponse"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-factors [get]
func (h *KPIFactorHandler) List(c *gin.Context) {
	var filter models.ListKPIFactorRequest

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
	filter.CriterionID = c.Query("criterion_id")

	result, err := h.kpiFactorService.List(c.Request.Context(), &filter)
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