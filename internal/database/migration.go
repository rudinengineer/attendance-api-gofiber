package database

import (
	"absensi-api/internal/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatal("Error migrating User:", err)
	}
	fmt.Println("✓ User migrated")

	if err := db.AutoMigrate(&domain.Departement{}); err != nil {
		log.Fatal("Error migrating Departement:", err)
	}
	fmt.Println("✓ Departement migrated")

	if err := db.AutoMigrate(&domain.Employee{}); err != nil {
		log.Fatal("Error migrating Employee:", err)
	}
	fmt.Println("✓ Employee migrated")

	if err := db.AutoMigrate(&domain.Attendance{}); err != nil {
		log.Fatal("Error migrating Attendance:", err)
	}
	fmt.Println("✓ Attendance migrated")

	if err := db.AutoMigrate(&domain.AttendanceHistory{}); err != nil {
		log.Fatal("Error migrating AttendanceHistory:", err)
	}
	fmt.Println("✓ AttendanceHistory migrated")

	fmt.Println("All tables migrated successfully!")
}
