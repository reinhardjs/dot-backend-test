package entity

type Product struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:100"`
	Price      float64
	CategoryID uint
	Category   Category
}
