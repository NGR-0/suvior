package config

import (
	"fmt"
	"log"
	"os"
	"suvio/pkg/models"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}

	DB_Host := os.Getenv("DATABASE_HOST")
	DB_Port := os.Getenv("DATABASE_PORT")
	DB_Name := os.Getenv("DATABASE_NAME")
	DB_Username := os.Getenv("DATABASE_USER")
	DB_Password := os.Getenv("DATABASE_PASS")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_Host, DB_Username, DB_Password, DB_Name, DB_Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database")
	}

	conn, err := db.DB()
	if err != nil {
		log.Fatalf("error connect database %s", err)
	}

	conn.SetMaxIdleConns(100)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Second * 10)

	return db
}

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.RatingProd{},
		&models.CategoryProd{},
		&models.Order{},
		&models.ProductOrder{},
		&models.ImageBlog{},
		&models.ImageProd{},
		&models.Enq{},
		&models.Coupon{},
		&models.Cart{},
		&models.CartItem{},
		&models.BlogCat{},
		&models.Blog{},
	)

	if err != nil {
		log.Fatalf("failed to migrate database")
	}
	log.Println("sukses Migrate DB")
}

func Initdb() {
	conn := newDatabase()

	Migration(conn)
}
