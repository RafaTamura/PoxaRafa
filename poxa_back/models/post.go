package models

type Post struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Tags       string `json:"tags"`
	CoverImage string `json:"cover_image"`
}
