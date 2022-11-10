package entity

type User struct {
	ID       int64  `json:"id" gorm:"collumn:id;type:bigint;primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone" gorm:"unique"`
	Address  string `json:"address"`
}
