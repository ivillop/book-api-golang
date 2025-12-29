package models

type Book struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Judul   string `json:"judul" binding:"required"`
	Penulis string `json:"penulis" binding:"required"`
	Tahun   int    `json:"tahun" binding:"required"`
}
