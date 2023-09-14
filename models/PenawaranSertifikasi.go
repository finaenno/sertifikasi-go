package models

type PenawaranSertifikasi struct {
	IdPenawaranSertifikasi        int64  `gorm:"primaryKey" json:"id_penawaran_sertifikasi"`
	NamaPenawaranSertifikasi      string `gorm:"type:varchar(300)" json:"nama_penawaran_sertifikasi"`
	DeskripsiPenawaranSertifikasi string `gorm:"type:text" json:"deskripsi_penawaran_sertifikasi"`
}
