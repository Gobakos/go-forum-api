package main

import (
    "log"
    "github.com/beego/beego/v2/client/orm"
    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
    _ "test/routers"
    beego "github.com/beego/beego/v2/server/web"
)

func main() { //Entry point for the app

	orm.RegisterDriver("mysql", orm.DRMySQL) //registers mysql db driver with ORM
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DB_SYNC:=true
	DB_PRINT:=true
	dbUser, err := beego.AppConfig.String("mysql::DB_USER")
	dbPassword, err := beego.AppConfig.String("mysql::DB_PASSWORD")
	dbName, err := beego.AppConfig.String("mysql::DB_NAME")
	dbHost, err := beego.AppConfig.String("mysql::DB_HOST")
	dbPort, err := beego.AppConfig.String("mysql::DB_PORT")

	// Register database
	orm.RegisterDataBase("default", "mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8")
	//Synchronizes the db schema with models,2nd argument force synchronization(drop & create),3nd boolean to print or not sql queries
	err = orm.RunSyncdb("default", DB_SYNC, DB_PRINT) 
	
	if err != nil {
		panic(err)
	}

	// Configure Beego settings
	beego.BConfig.WebConfig.DirectoryIndex = false //Eanble or not directory listing
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	// Run the application
	beego.Run()
}
