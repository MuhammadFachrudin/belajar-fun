package table

type User struct {
	UserID       uint `gorm:"primaryKey;autoIncrement"`
	NamaDepan    string
	NamaBelakang string
	AlamatID     int
}
