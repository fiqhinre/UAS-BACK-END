package models

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"` // ID user, primary key
	Username     string `json:"username" gorm:"unique"` // Username (unik)
	Password     string `json:"password"`            // Password (hashed)
	Email        string `json:"email" gorm:"unique"` // Email (unik)
	JenisUserID  uint   `json:"jenis_user_id"`       // ID Jenis User (foreign key)
	Photo        string `json:"photo"`              // Foto profil
	CreatedAt    string `json:"created_at"`         // Tanggal dibuat
	UpdatedAt    string `json:"updated_at"`         // Tanggal diperbarui
}
