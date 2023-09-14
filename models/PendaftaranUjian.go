package models

type PendaftaranUjian struct {
	IdPendaftaranUjian     int64 `gorm:"primaryKey" json:"id_pendaftaran_ujian"`
	IdPenawaranSertifikasi int64
	PenawaranSertifikasi   PenawaranSertifikasi `gorm:"foreignKey:IdPenawaranSertifikasi"`
}
