package entity

type Cart struct {
	ID         int64   `json:"id" gorm:"collumn:id;type:bigint;primaryKey;autoIncrement"`
	User_ID    int64   `json:"user_id"`
	Product_ID int64   `json:"product_id"`
	Quantity   int64   `json:"quantity"`
	Product    Product `gorm:"foreignKey:Product_ID"`
	User       User    `gorm:"foreignKey:User_ID"`
}
