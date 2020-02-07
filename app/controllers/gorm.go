package controllers

import (
        _ "github.com/go-sql-driver/mysql"
        "github.com/jinzhu/gorm"
        "github.com/revel/revel"
        "github.com/validyong/go-first/app/models"
        "log"
)

var DB *gorm.DB

func InitDB() {
        db, err := gorm.Open("mysql", dbInfoString())
        if err != nil {
                log.Panicf("Failed to connect to database: %v\n", err)
        }

        db.DB()
        db.AutoMigrate(&models.Comment{}) # ここで table の作成を行っている
        DB = &db
}

func dbInfoString() string {
        s, b := revel.Config.String("db.info")
        if !b {
                log.Panicf("database info not found")
        }

        return s
}
