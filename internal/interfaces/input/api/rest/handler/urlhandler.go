package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Gurveer1510/url-shortner/internal/core/url"
	"github.com/Gurveer1510/url-shortner/internal/usecase"
)

type URLHandler struct {
	URLService usecase.URLService
}

func NewURLHandler(urlService usecase.URLService) *URLHandler {
	return &URLHandler{URLService: urlService}
}

func (h *URLHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var newUrl url.URL

	if err := json.NewDecoder(r.Body).Decode(&newUrl); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := h.URLService.CreateShortUrl(newUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	newUrl.ShortURL = res.ShortURL

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUrl)

}

func (h *URLHandler) RedirectToLongURL(w http.ResponseWriter, r *http.Request) {
	var urlReq url.URLRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&urlReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	longUrl, err := h.URLService.GetLongURL(urlReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}


	http.Redirect(w, r, longUrl, http.StatusFound)
	// w.Write([]byte(longUrl))
}
