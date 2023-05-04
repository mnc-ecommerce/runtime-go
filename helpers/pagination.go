package helpers

import (
	"math"

	"github.com/zackyymughnii/runtime-go/web"
)

func CalculatePage(current int, limit int, total int) web.Pagination {
	var prev, next int
	totalPage := int(math.Ceil(float64(total) / float64(limit)))
	if current > 1 {
		prev = current - 1
	}
	if current < totalPage {
		next = current + 1
	}

	return web.Pagination{
		Page: web.Page{
			Current:  current,
			Previous: prev,
			Next:     next,
			Total:    totalPage,
		},
		Total: total,
		Limit: limit,
	}
}
