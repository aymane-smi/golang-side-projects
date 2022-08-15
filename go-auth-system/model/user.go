package model

type User struct {
	ID       uint   `gorm:"PrimaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password []byte `json:"password"`
}
