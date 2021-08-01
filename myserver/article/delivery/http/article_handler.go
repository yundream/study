package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"joinc.co.kr/study/myserver/domain"
)

type ArticleHandler struct {
	AUseCase domain.ArticleUseCase
}

func NewArticleHandler(router *mux.Router, us domain.ArticleUseCase) {
	handler := ArticleHandler{
		AUseCase: us,
	}
	router.HandleFunc("/article", handler.FetchArticle).Methods("GET")
	router.HandleFunc("/article/{id}", handler.GetArticle).Methods("GET")
}

func (a *ArticleHandler) FetchArticle(w http.ResponseWriter, r *http.Request) {
	articles, _ := a.AUseCase.Fetch(0, 0)
	response, _ := json.Marshal(articles)
	fmt.Fprint(w, string(response))
}

func (a *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 16)
	article, _ := a.AUseCase.GetByID(id)
	response, _ := json.Marshal(article)
	fmt.Fprint(w, string(response))
}
