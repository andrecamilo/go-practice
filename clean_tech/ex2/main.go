package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"

	"github.com/andrecamilo/go-practice/clean_tech/ex2/config"
	"github.com/andrecamilo/go-practice/clean_tech/ex2/infrastructure/datastore"
	"github.com/andrecamilo/go-practice/clean_tech/ex2/infrastructure/router"
	"github.com/andrecamilo/go-practice/clean_tech/ex2/registry"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
