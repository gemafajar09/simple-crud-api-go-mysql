package models

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Nama         string `json:"nama"`
	Telpon       string `json:"telpon"`
	Jeniskelamin string `json:"jeniskelamin"`
	Alamat       string `json:"alamat"`
}
