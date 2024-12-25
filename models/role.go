package models

type Role struct {
	ID        string   `bson:"_id,omitempty" json:"id"`
	Name      string   `bson:"name" json:"name"`
	Role      string   `bson:"role" json:"role"` // admin atau civitas
	JenisUser string   `bson:"jenis_user" json:"jenis_user"`
	Modul     []string `bson:"modul" json:"modul"` // List modul yang dimiliki
}
