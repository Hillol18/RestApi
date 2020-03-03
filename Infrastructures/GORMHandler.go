package Infrastructures

import (
	"fmt"
	"log"

	"github.com/nafi/oj-testing/Models"

	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type GORMHandler struct {
	db *gorm.DB
}

func (handler *GORMHandler) InitialMigration(connectionString string) error {
	conn, err := gorm.Open("postgres", connectionString)
	fmt.Println("Start AutoMigration")

	if err != nil {
		fmt.Println("Can not connect with DB")
		log.Fatal("This is the error: ", err)
		return err
	}

	handler.db = conn
	handler.db.Debug().AutoMigrate(&Models.UserInfo{})

	fmt.Println("End AutoMigration")
	return err
}
