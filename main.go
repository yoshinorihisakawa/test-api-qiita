package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func main() {
	db := NewMySqlDB()
	e := echo.New()

	// ルーティング
	e.GET("/users", func(c echo.Context) error {
		users := []*User{}
		err := db.Find(&users).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, users)
	})
	e.POST("/users", func(c echo.Context) error {
		user := &User{}
		// bind
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err := db.Save(user).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, user)
	})

	// サーバー起動
	e.Start(":" + os.Getenv("PORT"))
}

func NewMySqlDB() *gorm.DB {
	username := "rhltbiaqtrtetz"
	password := "4388a60afa9be97a56f952769d8ea617fc3c11e4525fa3cb514d0b8b19d4a124"
	dbName := "de3bn2nk12740q"
	dbHost := "ec2-23-23-173-30.compute-1.amazonaws.com"
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=require password=%s",
		dbHost, username, dbName, password)
	fmt.Println(dbUri)
	conn, err := gorm.Open("postgres", dbUri)
	if nil != err {
		panic(err)
	}

	// DBのエンジンを設定
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn
}

type User struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
