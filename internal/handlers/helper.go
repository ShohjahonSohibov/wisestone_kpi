package handlers

import (
	"fmt"
	"kpi/config"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	defaultPage = "1"
	maxLimit    = "1000000000"
)

type Helper struct {
}

func parseIntParam(val string) (int, error) {
	return strconv.Atoi(val)
}

func getPageOffsetLimit(c *gin.Context) (offset, limit int, err error) {
    // Get limit first since we need it for offset calculation
    cfg := config.Load()

    limitStr := c.Query("limit")
    if limitStr == "*" {
        limitStr = maxLimit
    }
    if limitStr == "" {
        limitStr = cfg.DefaultLimit
    }

    // Parse limit
    if limit, err = parseIntParam(limitStr); err != nil {
        return 0, 0, fmt.Errorf("invalid limit parameter: %w", err)
    }

    // Get page and offset
    pageStr := c.Query("page")
    if pageStr == "" {
        pageStr = defaultPage
    }
    
    offsetStr := c.Query("offset")
    if offsetStr == "" {
        offsetStr = cfg.DefaultOffset
    }

	var page int
	if page, err = parseIntParam(pageStr); err != nil {
		return 0, 0, fmt.Errorf("invalid page parameter: %w", err)
	}

	if offset, err = parseIntParam(offsetStr); err != nil {
		return 0, 0, fmt.Errorf("invalid offset parameter: %w", err)
	}

	// Calculate offset from page if needed
	if page > 0 && offset == 0 {
		offset = (page - 1) * limit
	}

	return offset, limit, nil
}