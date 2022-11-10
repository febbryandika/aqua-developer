package entity

type Checkout struct {
	ID      int64 `json:"id" gorm:"collumn:id;type:bigint;primaryKey;autoIncrement"`
	User_ID int64 `json:"user_id"`
	User    User  `gorm:"foreignKey:User_ID"`
}

type CheckoutProduct struct {
	ID          int64    `json:"id" gorm:"collumn:id;type:bigint;primaryKey;autoIncrement"`
	Checkout_ID int64    `json:"checkout_id"`
	Product_ID  int64    `json:"product_id"`
	Quantity    int64    `json:"quantity"`
	Product     Product  `gorm:"foreignKey:Product_ID"`
	Checkout    Checkout `gorm:"foreignKey:Checkout_ID"`
}
