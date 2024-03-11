package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price int
}

/*
TableNameメソッドを追加しその戻り値に文字列を指定することで、
Gormが構造体名から勝手にテーブル名を推測することがなくなる。
上記の場合「products」だと認識される。
*/
func (Product) TableName() string {
	return "product"
}
