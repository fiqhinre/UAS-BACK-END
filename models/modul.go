package models

type Modul struct {
	ID          uint   `json:"id" gorm:"primaryKey"` // ID modul, primary key
	Name        string `json:"name"`                // Nama modul
	Description string `json:"description"`         // Deskripsi modul
	IsActive    bool   `json:"is_active"`           // Status aktif modul
	CreatedAt   string `json:"created_at"`          // Tanggal dibuat
	UpdatedAt   string `json:"updated_at"`          // Tanggal diperbarui
}
