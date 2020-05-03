package jsfmt

import (
	"encoding/json"
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
	Order      string      `json:"order"`
	Conditions []Condition `json:"conditions"`
}

func ReadQuery(body string) (*Query, *errutils.ZinError) {
	var query = Query{}
	if body != "" {
		if err := json.Unmarshal([]byte(body), &query); err != nil {
			return nil, errutils.JsonFormatError(err.Error())
		}
	}
	return &query, nil
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

func (query *Query) Find(db *gorm.DB, model interface{}, result interface{}) (int, *errutils.ZinError) {
	var err *errutils.ZinError
	var total int

	if db, err = query.GenDB(db.Model(model)); err != nil {
		return 0, err
	}

	db.Count(&total)

	if query.Order != "" {
		db = db.Order(query.Order)
	}

	if err := db.Offset((query.Page - 1) * query.Limit).Limit(query.Limit).Find(result).Error; err != nil {
		return 0, errutils.DBOperationsFailed(err.Error())
	}
	return total, nil
}
