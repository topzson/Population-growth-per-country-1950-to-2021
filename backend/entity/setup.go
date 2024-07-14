package entity

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func stringToInt(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("CSVdata.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Region{}, &CSVdata{},
	)

	db = database

	csvFile2, err := os.Open("D:\\PTNFILE\\workspaces\\Population-growth-per-country-1950-to-2021\\backend\\entity\\data2.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile2.Close()

	reader2 := csv.NewReader(csvFile2)
	reader2.FieldsPerRecord = -1

	csvData2, err := reader2.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, each1 := range csvData2 {
		data2 := Region{
			Name:   each1[0],
			Region: each1[1],
			Url:    each1[2],
		}
		db.Model(&Region{}).Create(&data2)
	}

	csvFile, err := os.Open("D:\\PTNFILE\\workspaces\\Population-growth-per-country-1950-to-2021\\backend\\entity\\data.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, each := range csvData {
		data := CSVdata{
			Name:     each[0],
			Date:     each[1],
			Value:    each[2],
			Category: each[0],
		}
		db.Model(&CSVdata{}).Create(&data)
	}

}
