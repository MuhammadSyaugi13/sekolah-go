package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/MuhammadSyaugi13/go-project-1/app"
	"github.com/MuhammadSyaugi13/go-project-1/controller"
	"github.com/MuhammadSyaugi13/go-project-1/helper"
	"github.com/MuhammadSyaugi13/go-project-1/repository"
	"github.com/MuhammadSyaugi13/go-project-1/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validateService := validator.New()

	kelasRepository := repository.NewKelasRepository()
	KelasService := service.NewKelasService(kelasRepository, db, validateService)
	kelasController := controller.NewKelasController(KelasService)

	router := httprouter.New()

	router.POST("/api/kelas", kelasController.Create)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Menjalankan server......")
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
