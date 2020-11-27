package table

import "time"

type Produk struct {
	ProdukID   uint `gorm:"primaryKey;autoIncrement"`
	NamaProduk string
	BrandID    int
	Stock      int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
