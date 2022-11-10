package entity

type Product struct {
	ID           int64  `json:"id" gorm:"collumn:id;type:bigint;primaryKey;autoIncrement"`
	Product_Name string `json:"product_name"`
	Picture      string `json:"picture"`
	Price        int64  `json:"price"`
	Category     string `json:"category"`
	Description  string `json:"description"`
}
