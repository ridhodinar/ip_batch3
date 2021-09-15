package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"review_movie/database"
	"review_movie/model"
	"review_movie/seed"

	//"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgreConfig *database.PostgreConfig
	db            *gorm.DB
)

func initEnv(path string, name string) error {
	/*
		load env file and other environment variable
	*/
	viper.AddConfigPath(path)
	viper.SetConfigName("local")
	// Somehow it's error
	// viper.SetConfigFile("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&postgreConfig); err != nil {
		return err
	}
	return nil
}

func initDb() {
	/*
		db initialization
	*/
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", postgreConfig.DBHost, postgreConfig.DBUser, postgreConfig.DBPassword, postgreConfig.DBName, postgreConfig.DBPort)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	dbSql, _ := db.DB()
	if pingErr := dbSql.Ping(); pingErr != nil {
		fmt.Println(pingErr)
	} else {
		fmt.Println("Connected")
		db.AutoMigrate(&model.User{}, &model.Movie{}, &model.Genre{}, &model.Review{}, &model.MovieGenre{})
	}
}

func main() {
	if err := initEnv(".", "local"); err != nil {
		log.Fatal(err)
	}
	initDb()
	handleArgs()
}

func handleArgs() {
	/*
		Handle argument to seed
	*/
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			seed.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
}
