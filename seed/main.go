package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jaswdr/faker"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type FakeData struct {
	gorm.Model
	ID          string
	SensorId    string
	SensorValue float64
	Company     string
	Timestamp   int64
}

func main() {
	fmt.Println("Seeding database!")
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Truncate the table (drop and auto-migrate)
	if err := db.Migrator().DropTable(&FakeData{}); err != nil {
		log.Fatalf("failed to drop table: %v", err)
	}
	if err := db.AutoMigrate(&FakeData{}); err != nil {
		log.Fatalf("failed to migrate table: %v", err)
	}

	// Initialize the faker generator
	faker := faker.New()

	// Predefined list of companies
	companies := []string{"Dickens PLC", "Goldner PLC", "Brown-Brown", "Herzog Group", "Kuffenbach Group", "Kuhic Group", "Kuhlman Group", "Kuhn Group", "Kulas Group"}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Loop to create 1000 entries
	for i := 0; i < 10000; i++ {
		fakeData := FakeData{
			ID:          faker.UUID().V4(),
			SensorId:    faker.UUID().V4(),
			SensorValue: faker.Float64(0, 0, 100),
			Company:     companies[rand.Intn(len(companies))], // Randomly select a company
			Timestamp:   faker.Time().Unix(time.Now()),
		}

		db.Create(&fakeData)
	}

	fmt.Println("Seeding database successfull!")

}
