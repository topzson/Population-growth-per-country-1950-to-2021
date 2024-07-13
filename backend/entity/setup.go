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
		&CSVdata{},
	)

	db = database

	csvFile, err := os.Open("D:\\PTNFILE\\workspaces\\Population-growth-per-country-1950-to-2021\\backend\\entity\\data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ตำแหน่งงาน --------------------------------------------------------------
	for _, each := range csvData {
		data := CSVdata{
			Name:     each[0],
			Date:     each[1],
			Value:    each[2],
			Category: each[0],
			// 	Population_1:        stringToInt(each[3]),
			// 	Population_5:        stringToInt(each[4]),
			// 	Population_15:       stringToInt(each[5]),
			// 	Population_25:       stringToInt(each[6]),
			// 	Population_15_to_64: stringToInt(each[7]),
			// 	Population_older_15: stringToInt(each[8]),
			// 	Population_older_18: stringToInt(each[9]),
			// 	Population_at_1:     stringToInt(each[10]),
			// 	Population_1_to_4:   stringToInt(each[11]),
			// 	Population_5_to_9:   stringToInt(each[12]),
			// 	Population_10_to_14: stringToInt(each[13]),
			// 	Population_15_to_19: stringToInt(each[14]),
			// 	Population_20_to_29: stringToInt(each[15]),
			// 	Population_30_to_39: stringToInt(each[16]),
			// 	Population_40_to_49: stringToInt(each[17]),
			// 	Population_50_to_59: stringToInt(each[18]),
			// 	Population_60_to_69: stringToInt(each[19]),
			// 	Population_70_to_79: stringToInt(each[20]),
			// 	Population_80_to_89: stringToInt(each[21]),
			// 	Population_90_to_99: stringToInt(each[22]),
			// 	Population_100:      stringToInt(each[23]),
		}
		db.Model(&CSVdata{}).Create(&data)
	}

}
