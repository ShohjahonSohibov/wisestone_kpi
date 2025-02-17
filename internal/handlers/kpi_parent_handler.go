package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kpi/internal/models"
	"kpi/internal/services"
)

type KpiParentHandler struct {
	kpiParentService *services.KpiParentService
}

func NewKPIParentHandler(kpiParentService *services.KpiParentService) *KpiParentHandler {
	return &KpiParentHandler{
		kpiParentService: kpiParentService,
	}
}

// @Summary Create KPI Parent
// @Description Create a new KPI Parent
// @Tags KPI Parents
// @Accept json
// @Produce json
// @Param request body models.CreateKPIParent true "KPI Parent creation request"
// @Success 201 {object} models.KPIParent
// @Failure 400 {object} object{error=string}
// @Failure 500 {object} object{error=string}
// @Router /api/v1/kpi-parents [post]
func (h *KpiParentHandler) Create(c *gin.Context) {
	var req models.KPIParent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.kpiParentService.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, "success")
}

// @Summary Update KPI Parent
// @Description Update an existing KPI Parent by ID
// @Tags KPI Parents
// @Accept json
// @Produce json
// @Param id path string true "KPI Parent ID"
// @Param request body models.UpdateKPIParent true "KPI Parent update request"
// @Success 200 {object} models.KPIParent
// @Failure 400 {object} object{error=string}
// @Failure 500 {object} object{error=string}
// @Router /api/v1/kpi-parents/{id} [put]
func (h *KpiParentHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req models.KPIParent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.ID = id
	err := h.kpiParentService.Update(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "success")
}

// @Summary Delete KPI Parent
// @Description Delete a KPI Parent by ID
// @Tags KPI Parents
// @Produce json
// @Param id path string true "KPI Parent ID"
// @Success 204 "No Content"
// @Failure 500 {object} object{error=string}
// @Router /api/v1/kpi-parents/{id} [delete]
func (h *KpiParentHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.kpiParentService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Get KPI Parent by ID
// @Description Get a KPI Parent by its ID
// @Tags KPI Parents
// @Produce json
// @Param id path string true "KPI Parent ID"
// @Success 200 {object} models.KPIParent
// @Failure 500 {object} object{error=string}
// @Router /api/v1/kpi-parents/{id} [get]
func (h *KpiParentHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.kpiParentService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

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

	result, err := h.kpiParentService.List(c.Request.Context(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}