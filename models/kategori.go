package models

type Kategori struct {
	ID          uint   `json:"id" gorm:"primaryKey"` // ID kategori, primary key
	Name        string `json:"name"`                // Nama kategori
	Description string `json:"description"`         // Deskripsi kategori
	CreatedAt   string `json:"created_at"`          // Tanggal dibuat
	UpdatedAt   string `json:"updated_at"`          // Tanggal diperbarui
}
