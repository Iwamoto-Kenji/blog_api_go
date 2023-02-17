package api

import (
	"database/sql"
	"net/http"

	"github.com/Iwamoto-Kenji/blog_api_go/controllers"
	"github.com/Iwamoto-Kenji/blog_api_go/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// NewMyAppService 関数から得られるサービス層を使う
	ser := services.NewMyAppService(db)
	// そのサービス層とつながった状態で
	// NewArticleController 関数・NewCommentController 関数から得られるコントローラ層を使う
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}
