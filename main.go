package main

import (
	"f2_gc1/config"
	"f2_gc1/handler"
	"f2_gc1/migration"
	"f2_gc1/repository"
	"f2_gc1/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Gagal membaca file .env, menggunakan environment default.")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	migration.CreateTableCustomer(db)

	customerRepo := repository.NewCustomerRepository(db)

	customerUserCase := usecase.NewCustomerUseCase(customerRepo)

	customerHandler := handler.NewCustomerHandler(customerUserCase)

	mux := mux.NewRouter()

	mux.HandleFunc("/customer", customerHandler.GetAll).Methods("GET")         //method Get untuk ambil semua
	mux.HandleFunc("/customer", customerHandler.Create).Methods("POST")        //method Post untuk create
	mux.HandleFunc("/customer/{id}", customerHandler.Update).Methods("PUT")    //method Put untuk update
	mux.HandleFunc("/customer/{id}", customerHandler.Delete).Methods("DELETE") //method Delete untuk delete
	mux.HandleFunc("/customer/{id}", customerHandler.GetByID).Methods("GET")   //method Get untuk ambil satuan.

	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
