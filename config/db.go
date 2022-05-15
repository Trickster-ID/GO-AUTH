package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB{
	// errEnv := godotenv.Load()
	// if errEnv != nil {
	// 	panic("fail to load .env")
	// }
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")
	sslmode := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",host,user,pass,dbname,port,sslmode)
	
	fmt.Println(dsn)
	db, errDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("--------fail connect to DB-------")
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		panic(errDB)
	}
	fmt.Println("====================================")
	fmt.Println("--------success connect to DB-------")
	fmt.Println("====================================")
	return db
}

func CloseDatabaseConnection(db *gorm.DB){
	dbcon, err := db.DB()
	if err != nil {
		panic("fail to close con database")
	}
	dbcon.Close()
}