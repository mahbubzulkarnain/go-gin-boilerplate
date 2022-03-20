package repository

import (
	"github.com/gomodul/abstraction"
	"gorm.io/gorm"
)

// Pagination ...
func Pagination(p *abstraction.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p != nil {
			p.Init()
			return db.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetOrderBy())
		}
		return db
	}
}
