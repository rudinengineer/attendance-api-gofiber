package database

import (
	"absensi-api/internal/domain"
	"absensi-api/pkg/utils"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var ctx = context.Background()

func RunSeeder(db *gorm.DB) {
	// Create User
	_, err := gorm.G[domain.User](db).First(ctx)
	if err != nil {
		password, err := utils.EncryptPassword("123456")
		if err == nil {
			if err := gorm.G[domain.User](db).Create(ctx, &domain.User{
				Name:     "Superadmin",
				Username: "superadmin",
				Password: password,
			}); err != nil {
				fmt.Println("✓ User seeded")
			}
		}
	}

	// Create Departement
	_, err = gorm.G[domain.Departement](db).First(ctx)
	if err != nil {
		maxClockInTime, err := time.Parse("2006-01-02 15:04:05", "2006-01-02 "+"08:00:00")
		if err == nil {
			maxClockInOut, err := time.Parse("2006-01-02 15:04:05", "2006-01-02 "+"16:00:00")
			if err == nil {
				if err := gorm.G[domain.Departement](db).Create(ctx, &domain.Departement{
					DepartementName: "North Departement",
					MaxClockInTime:  maxClockInTime,
					MaxClockInOut:   maxClockInOut,
				}); err != nil {
					fmt.Println("✓ Departement seeded")
				}
			}
		}
	}

	// Create Employee
	_, err = gorm.G[domain.Employee](db).First(ctx)
	if err != nil {
		if err := gorm.G[domain.Employee](db).Create(ctx, &domain.Employee{
			DepartementID: 1,
			Name:          "Erick Setyawan",
			EmployeeID:    "12345678",
			Address:       "jln. pramuka No.83",
		}); err == nil {
			fmt.Println("✓ Employee seeded")
		}
	}
}
