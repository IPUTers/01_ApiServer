package domain

import "gorm.io/gorm"

type TestApi struct {
	gorm.Model
	Name string
	Age  int
}

func (TestApi) TableName() string {
	return "test_api"
}
