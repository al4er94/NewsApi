// API приложения GoNews.
package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"NewsApi/pkg/db"

	"github.com/gorilla/mux"
)

type API struct {
	db *db.DB
	r  *mux.Router
}

type News struct {
	Id    int
	Title string
}

// Конструктор API.
func New(db *db.DB) *API {
	a := API{db: db, r: mux.NewRouter()}
	a.endpoints()
	return &a
}

// Router возвращает маршрутизатор для использования
// в качестве аргумента HTTP-сервера.
func (api *API) Router() *mux.Router {
	return api.r
}

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	// Получение новостей
	api.r.HandleFunc("/news/{n}", api.newsShortDetailed).Methods(http.MethodGet, http.MethodOptions)
	api.r.HandleFunc("/details/{id}", api.newsFullDetailed).Methods(http.MethodGet, http.MethodOptions)
}

func (api *API) newsShortDetailed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	s := mux.Vars(r)["n"]
	n, _ := strconv.Atoi(s)
	news, err := api.db.News(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prepareNews := []News{}

	for _, ne := range news {
		prepareNews = append(prepareNews, News{Id: ne.ID, Title: ne.Title})
	}

	json.NewEncoder(w).Encode(prepareNews)
}

func (api *API) newsFullDetailed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	s := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(s)
	news, err := api.db.NewsDetail(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(news[0])
}
