package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_articleDelivery "joinc.co.kr/study/myserver/article/delivery/http"
	"joinc.co.kr/study/myserver/article/repository/mysql"
	_articleUsecase "joinc.co.kr/study/myserver/article/usecase"
)

func main() {
	router := mux.NewRouter()
	db, err := mysql.NewMySQLRepository()
	if err != nil {
		os.Exit(1)
	}
	au := _articleUsecase.NewArticleUseCase(db)
	_articleDelivery.NewArticleHandler(router, au)
	log.Fatal(http.ListenAndServe(":8080", router))
}
