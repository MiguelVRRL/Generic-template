package main

import (
	"fmt"
	"log"
    
  "gihub.com/MiguelVRRL/template/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  "github.com/spf13/viper"
)



var DB *gorm.DB

func init() {
  viper.SetConfigFile(".env")
  viper.ReadInConfig()
}

func setupRouter() {
}

func setupDBConfigs() ( host, user, password, dbname, port string ) {
  host = viper.GetString("DB_HOST")
  user = viper.GetString("DB_USER")
  password = viper.GetString("DB_PASSWORD")
  dbname =  viper.GetString("DB_NAME")
  port = viper.GetString("DB_PORT")

  return 
}

func setupDB() {
  host,user, password, dbname, port := setupDBConfigs()
  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",  host,user, password, dbname, port )
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
   log.Fatal("Payment failed")
  }
  DB = db
  DB.AutoMigrate(models.User{},models.DataUser{})

}
