package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Product
type Product struct {
	gorm.Model
	// ProductID    uint         `gorm:"type: int; PRIMARY_KEY; AUTO_INCREMENT"`
	Name           string  `json:"name"`
	Ingredients    string  `json:"ingredients"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	ManufacturerID uint
	Manufacturer   Manufacturer `gorm:"foreignKey:ManufacturerID"`
}

// Manufacturer
type Manufacturer struct {
	gorm.Model
	// ManufacturerID uint      `gorm:"type: int; PRIMARY_KEY; AUTO_INCREMENT"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Contacts  string `json:"contacts"`
	ProductID uint
	Product   []Product `gorm:"foreignKey:product_id"`
}

// DBMigrate
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.DropTableIfExists(&Product{}, &Manufacturer{})
	db.AutoMigrate(&Product{}, &Manufacturer{})
	// db.Model(&Product{}).AddForeignKey("manufacturer_id", "manufacturers(id)", "CASCADE", "CASCADE")
	// db.Model(&Manufacturer{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	return db
}
