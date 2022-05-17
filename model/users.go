package model

type Users struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Username string `json:"username" gorm:"Type:varchar(100),unique"`
	Password []byte `json:"password" gorm:"Type:varchar(100)"`
	// gorm.Model
}

type UsersRequest struct {
	Name     string `json:"name" binding:"required,min=3"`
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=7"`
}
