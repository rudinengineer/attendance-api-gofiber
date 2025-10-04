package domain

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Username string `gorm:"type:varchar(255);uniqueIndex" json:"username"`
	Password string `gorm:"type:text" json:"password"`
}
