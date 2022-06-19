package tools

import (
	"fmt"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
)

func PaginateIndex(c *gin.Context, page int, pageLimit int, count int64, handler string, queries ...string) gin.H {
	var queriesString string
	for _, query := range queries {
		queriesString += query
	}
	totalPages := math.Ceil(float64(count) / float64(pageLimit))
	var next string
	var previous string
	if page+1 <= int(totalPages) {
		next = fmt.Sprintf("%s/%s?page=%d%s", c.Request.Host, handler, page+1, queriesString)
	}
	if page-1 > 0 {
		previous = fmt.Sprintf("%s/%s?page=%d%s", c.Request.Host, handler, page-1, queriesString)
	}
	links := gin.H{
		"count": count,
		"first": gin.H{
			"href": fmt.Sprintf("%s/%s?page=%d%s", c.Request.Host, handler, 1, queriesString),
		},
		"last": gin.H{
			"href": fmt.Sprintf("%s/%s?page=%.f%s", c.Request.Host, handler, totalPages, queriesString),
		},
		"next": gin.H{
			"href": next,
		},
		"previous": gin.H{
			"href": previous,
		},
		"self": gin.H{
			"href": fmt.Sprintf("%s/%s?page=%d%s", c.Request.Host, handler, page, queriesString),
		},
	}
	return links
}

func FormatErr(err string) string {
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Email") {
		return "Email is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "email") && strings.Contains(err, "Email") {
		return "Invalid email format"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Password") {
		return "Password is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Email") {
		return "Email is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "min") && strings.Contains(err, "Password") {
		return "Password must have more than 3 characters"
	}
	if strings.Contains(err, "SQLSTATE") && strings.Contains(err, "unique") && strings.Contains(err, "email") {
		return "Email is already used"
	}
	if strings.Contains(err, "SQLSTATE") && strings.Contains(err, "varying(255)") {
		return "Fields must be under 255 characters"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Name") {
		return "Name is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Age") {
		return "Age is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Gender") {
		return "Gender is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "max") && strings.Contains(err, "Name") {
		return "Name must be under 50 characters"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "lte") && strings.Contains(err, "Age") {
		return "Age must be less than 90"
	}
	if strings.Contains(err, "SQLSTATE") && strings.Contains(err, "gender") {
		return "Gender must be male, female or x"
	}
	if strings.Contains(err, "cannot unmarshal") {
		return "Invalid field type"
	}
	return err
}
