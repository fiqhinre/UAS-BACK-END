package models

type JenisUser struct {
	ID        uint   `json:"id" gorm:"primaryKey"` // ID jenis user, primary key
	Name      string `json:"name"`                // Nama jenis user
	CreatedAt string `json:"created_at"`          // Tanggal dibuat
	UpdatedAt string `json:"updated_at"`          // Tanggal diperbarui
}
