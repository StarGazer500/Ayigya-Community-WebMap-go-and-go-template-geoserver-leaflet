package main

import (
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/inits/db"
	"github.com/joho/godotenv"

	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/models"
	"fmt"

)


func init() {
	db.InitpgDb()
}

func deinit() {

	db.PG.Db.Close()

}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")

	}
	defer deinit()
	// _, err := models.PerformOperation(db.PG.Db, models.BuildingTable.TableName, "shape__len", "Greater than (>)","5")
	// if err != nil {
	// 	fmt.Println("Error occured",err)
	// }


	_,err := models.SearchAllTables(db.PG.Db, "comp")
	if err!=nil{
		fmt.Println("error occured",err)
	}else{
		// fmt.Println(data)
	}

	
}
