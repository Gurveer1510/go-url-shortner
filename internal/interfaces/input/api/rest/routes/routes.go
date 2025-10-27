package routes

import (
	"net/http"

	"github.com/Gurveer1510/url-shortner/internal/interfaces/input/api/rest/handler"
	"github.com/go-chi/chi"
)

func InitRoutes(urlHandler handler.URLHandler) http.Handler {
	router := chi.NewRouter()
	router.Post("/create", urlHandler.CreateShortUrl)
	router.Get("/sendmeto", urlHandler.RedirectToLongURL)
	return router
}
