package db

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var DB *gorm.DB

func SqlServerDb() {
	// for windows authentication
	//connectionString := "sqlserver://@127.0.0.1:1433?database=GoLangDB"
	// for SQL Server authentication
	connectionString := "sqlserver://sonarsqlauth:Eqs2024.@5CD114LJ30:1433?database=GoLangDB"

	var err error
	DB, err = gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Create
	DB.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	DB.First(&product, 1)                 // find product with integer primary key
	DB.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	DB.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	DB.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	DB.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	DB.Delete(&product, 1)
}

/**********************************************************************************************************/

/*
package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func SqliteDb() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Create
	DB.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	DB.First(&product, 1)                 // find product with integer primary key
	DB.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	DB.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	DB.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	DB.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	DB.Delete(&product, 1)
}

*/

/*
func SqliteDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)

}
*/
