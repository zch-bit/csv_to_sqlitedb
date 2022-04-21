package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	path := "./x.csv"
	file := ReadFile(path)
	objects := LoadData(file)
	db := connectDB()

	// Batch insert data
	db.CreateInBatches(objects[1:], 1000)
}

func connectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintln("failed to connect with database", err.Error()))
	}
	err = db.AutoMigrate(&Object{})
	if err != nil {
		panic(fmt.Sprintln("failed to migrate database", err.Error()))
	}
	return db
}

func ReadFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("read file error")
		return nil
	}
	return file
}

func LoadData(file *os.File) []*Object {
	defer file.Close()
	reader := csv.NewReader(file)

	objects := make([]*Object, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			fmt.Println("read all data")
			break
		}
		if err != nil {
			fmt.Println("csv reader error: ", err.Error())
			break
		}
		objects = append(objects, ParseObject(record))
	}

	return objects
}

func ParseObject(record []string) *Object {
	return &Object{
		//...
	}
}

type Object struct {
	// Define the object here
	_ struct{}
	// ...
}
