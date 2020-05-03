package jsfmt

import (
	"github.com/jinzhu/gorm"
	"github.com/zpdev/zins/common/errutils"
)

type Condition struct {
	Field    string      `json:"field"`
	Value    interface{} `json:"value"`
	Operator int         `json:"operator"`
	//Logic    int         `json:"logic"`
}

type Query struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	Total      int         `json:"total"`
	Conditions []Condition `json:"conditions"`
}

func (query *Query) GenDB(db *gorm.DB) (*gorm.DB, *errutils.ZinError) {
	for i := 0; i < len(query.Conditions); i++ {
		condition := query.Conditions[i]

		switch condition.Operator {
		case 1: // equal
			db = db.Where(condition.Field+" = ?", condition.Value)
		case 2: // like
			db = db.Where(condition.Field+" LIKE ?", "%"+condition.Value.(string)+"%")
		case 3: // not equal
			db = db.Not(condition.Field+" = ?", condition.Value)
		default:
			return nil, errutils.ConditionError()
		}
	}
	return db, nil
}
