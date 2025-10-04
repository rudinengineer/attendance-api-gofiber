package connection

import (
	"absensi-api/internal/config"
	"absensi-api/internal/database"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(configuration config.Database) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configuration.User,
		configuration.Pass,
		configuration.Host,
		configuration.Port,
		configuration.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("failed to connect database")
	}

	sqldb, err := db.DB()
	if err != nil {
		log.Fatal("failed to initialized database")
	}

	// Ping DB
	if err := sqldb.Ping(); err != nil {
		log.Fatal("failed to ping database")
	}

	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxLifetime(time.Hour)

	// Drop Table
	sqldb.Exec("SET FOREIGN_KEY_CHECKS = 0")
	sqldb.Exec("DROP TABLE IF EXISTS attendance_histories")
	sqldb.Exec("DROP TABLE IF EXISTS attendances")
	sqldb.Exec("DROP TABLE IF EXISTS employees")
	sqldb.Exec("DROP TABLE IF EXISTS departements")
	sqldb.Exec("DROP TABLE IF EXISTS users")
	sqldb.Exec("SET FOREIGN_KEY_CHECKS = 1")

	fmt.Println("Tables dropped")

	// Database Migrate
	database.RunMigrate(db)

	// Database Seeder
	database.RunSeeder(db)

	return db
}
