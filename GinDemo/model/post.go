package model

type Post struct {
	ID         uint `json:"id"`
	UserID     uint `json:"user_id" gorm:"not null"`
	CategoryID uint `json:"category_id" gorm:"not null"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type:text;not null"`
	CreatedAt  Time   `json:"create_at"`
	UpdatedAt  Time   `json:"update_at"`
}
