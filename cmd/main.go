package main

import (
	"log"
	"myapp/handler"
	"myapp/infrastructure"
	"myapp/repository"
	"myapp/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := infrastructure.NewDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
