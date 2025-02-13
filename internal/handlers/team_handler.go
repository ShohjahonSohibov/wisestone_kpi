package handlers

import (
	"kpi/internal/models"
	"kpi/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	teamService *services.TeamService
}

func NewTeamHandler(teamService *services.TeamService) *TeamHandler {
	return &TeamHandler{teamService: teamService}
}

// GetTeam godoc
// @Security ApiKeyAuth
// @Summary Get team by id
// @Description Get team details by id
// @Tags Teams
// @Accept json
// @Produce json
// @Param id path string true "Team Id"
// @Success 200 {object} models.Team
// @Failure 404 {object} map[string]string "error: Team not found"
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /api/v1/teams/{id} [get]
func (h *TeamHandler) GetTeam(c *gin.Context) {
	id := c.Param("id")
	team, err := h.teamService.GetById(c.Request.Context(), id)
	if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
					"status":  http.StatusNotFound,
					"message": err.Error(),
			})
			return
	}
	c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   team,
	})
}

// CreateTeam godoc
// @Security ApiKeyAuth
// @Summary Create a new team
// @Description Create a new team with the provided details
// @Tags Teams
// @Accept json
// @Produce json
// @Param team body models.CreateTeam true "Team Details"
// @Success 201 {object} map[string]string "message: Team created successfully"
// @Failure 400 {object} map[string]string "error: Validation error"
// @Router /api/v1/teams [post]
func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": err.Error(),
			})
			return
	}

	if err := h.teamService.Create(c.Request.Context(), &team); err != nil {
			c.JSON(http.StatusConflict, gin.H{
					"status":  http.StatusConflict,
					"message": err.Error(),
			})
			return
	}

	c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "team created successfully",
	})
}

// UpdateTeam godoc
// @Security ApiKeyAuth
// @Summary Update existing team
// @Description Update team details by ID
// @Tags Teams
// @Accept json
// @Produce json
// @Param id path string true "Team ID"
// @Param team body models.Team true "Team Details"
// @Success 200 {object} map[string]string "message: Team updated successfully"
// @Failure 400 {object} map[string]string "error: Validation error"
// @Failure 404 {object} map[string]string "error: Team not found"
// @Router /api/v1/teams/{id} [put]
func (h *TeamHandler) UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": err.Error(),
			})
			return
	}

	team.ID = id
	if err := h.teamService.Update(c.Request.Context(), &team); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
					"status":  http.StatusNotFound,
					"message": err.Error(),
			})
			return
	}

	c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "team updated successfully",
	})
}

// DeleteTeam godoc
// @Security ApiKeyAuth
// @Summary Delete team
// @Description Delete team by ID
// @Tags Teams
// @Accept json
// @Produce json
// @Param id path string true "Team Id"
// @Success 200 {object} map[string]string "message: Team deleted successfully"
// @Failure 404 {object} map[string]string "error: Team not found"
// @Router /api/v1/teams/{id} [delete]
func (h *TeamHandler) DeleteTeam(c *gin.Context) {
	id := c.Param("id")
	if len(id) != 24 {
			c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": "invalid ID format",
			})
			return
	}

	if err := h.teamService.Delete(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
					"status":  http.StatusNotFound,
					"message": err.Error(),
			})
			return
	}
	c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "team deleted successfully",
	})
}

// ListTeams godoc
// @Security ApiKeyAuth
// @Summary List all teams
// @Description Get a list of all teams
// @Tags Teams
// @Accept json
// @Produce json
// @Param filter query models.ListTeamsRequest false "Filter parameters"
// @Success 200 {object} models.ListTeamsResponse
// @Failure 400 {object} map[string]string "error: Invalid parameters"
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /api/v1/teams [get]
func (h *TeamHandler) ListTeams(c *gin.Context) {
	filter := &models.ListTeamsRequest{}
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
	filter.MultiSearch = c.Query("multi_search")
	filter.SortOrder = c.Query("sort_order")

	teams, err := h.teamService.List(c.Request.Context(), filter)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
					"status":  http.StatusInternalServerError,
					"message": err.Error(),
			})
			return
	}
	c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   teams,
	})
}