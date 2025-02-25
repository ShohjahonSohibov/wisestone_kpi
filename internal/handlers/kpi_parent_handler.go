package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"kpi/internal/models"
	"kpi/internal/services"
)

type KpiParentHandler struct {
	kpiParentService *services.KPIParentService
}

func NewKPIParentHandler(kpiParentService *services.KPIParentService) *KpiParentHandler {
	return &KpiParentHandler{
		kpiParentService: kpiParentService,
	}
}

// @Security ApiKeyAuth
// @Summary Create KPI Parent
// @Description Create a new KPI Parent
// @Tags KPI Parents
// @Accept json
// @Produce json
// @Param request body models.CreateKPIParent true "KPI Parent creation request"
// @Success 201 {object} map[string]interface{} "status: 201, message: KPI Parent created successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-parents [post]
func (h *KpiParentHandler) Create(c *gin.Context) {
	var req models.KPIParent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err := h.kpiParentService.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "kpi parent created successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Update KPI Parent
// @Description Update an existing KPI Parent by ID
// @Tags KPI Parents
// @Accept json
// @Produce json
// @Param id path string true "KPI Parent ID"
// @Param request body models.UpdateKPIParent true "KPI Parent update request"
// @Success 200 {object} map[string]interface{} "status: 200, message: KPI Parent updated successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-parents/{id} [put]
func (h *KpiParentHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req models.KPIParent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req.ID = id
	err := h.kpiParentService.Update(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "kpi parent updated successfully",
	})
}


// @Security ApiKeyAuth
// @Summary Update KPI Parent Status
// @Description Update an existing KPI Parent Status by ID
// @Tags KPI Parents
// @Accept json
// @Produce json
// @Param id path string true "KPI Parent ID"
// @Param request body models.UpdateKPIParentStatus true "KPI Parent status update request"
// @Success 200 {object} map[string]interface{} "status: 200, message: KPI Parent status updated successfully"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-parents/status/{id} [put]
func (h *KpiParentHandler) UpdateStatus(c *gin.Context) {
	var req models.UpdateKPIParentStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = c.Param("id")
	if req.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if err := h.kpiParentService.UpdateStatus(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "status updated successfully"})
}

// @Security ApiKeyAuth
// @Summary Delete KPI Parent
// @Description Delete a KPI Parent by ID
// @Tags KPI Parents
// @Produce json
// @Param id path string true "KPI Parent ID"
// @Success 200 {object} map[string]interface{} "status: 200, message: KPI Parent deleted successfully"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-parents/{id} [delete]
func (h *KpiParentHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.kpiParentService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "kpi parent deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Summary Get KPI Parent by ID
// @Description Get a KPI Parent by its ID
// @Tags KPI Parents
// @Produce json
// @Param id query string false "KPI Parent ID"
// @Param type query string false "KPI Parent type"
// @Success 200 {object} map[string]interface{} "status: 200, data: KPI Parent object"
// @Failure 500 {object} map[string]interface{} "status: 500, message: error message"
// @Router /api/v1/kpi-parents/single [get]
func (h *KpiParentHandler) GetByID(c *gin.Context) {
	id := c.Query("id")
	kpiType := c.Query("type")

	result, err := h.kpiParentService.GetByID(c.Request.Context(), id, kpiType)
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
// @Summary List KPI Parents
// @Description Get a list of KPI Parents with filtering and pagination
// @Tags KPI Parents
// @Produce json
// @Param filter query models.ListKPIParentRequest false "Filter parameters"
// @Success 200 {object} models.ListKPIParentResponse
// @Failure 400 {object} object{error=string}
// @Failure 500 {object} object{error=string}
// @Router /api/v1/kpi-parents [get]
func (h *KpiParentHandler) List(c *gin.Context) {
	var filter models.ListKPIParentRequest

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
	filter.Year = c.Query("year")
	filter.Status = c.Query("status")
	filter.Status = c.Query("type")

	result, err := h.kpiParentService.List(c.Request.Context(), &filter)
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
