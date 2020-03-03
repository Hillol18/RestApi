package main

import (
	"log"

	_ "github.com/lib/pq"

	"fmt"
	"sync"

	controllers "github.com/nafi/oj-testing/Controllers"
	"github.com/nafi/oj-testing/Infrastructures"
	"github.com/nafi/oj-testing/Repositories"
	"github.com/nafi/oj-testing/Services"
)

type kernel struct{}

const (
	host     = "localhost"
	port     = 5432
	user     = "nafi"
	password = "123"
	dbname   = "mydb"
)

type IServiceContainer interface {
	InjectHomeController() controllers.HomeController
}

func (k *kernel) InjectHomeController() controllers.HomeController {
	fmt.Println("Start of Injecting Home Controller")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Println(psqlInfo)

	dbHandler := &Infrastructures.GORMHandler{}
	dbHandler.InitialMigration(psqlInfo)
	homeRepository := &Repositories.HomeRepository{dbHandler}
	homeServices := &Services.HomeService{homeRepository}
	homeController := controllers.HomeController{homeServices}

	fmt.Println("End of Injecting Home Controller")
	return homeController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
