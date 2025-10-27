package usecase

import (
	"github.com/Gurveer1510/url-shortner/internal/adaptors/persistance"
	"github.com/Gurveer1510/url-shortner/internal/core/url"
	"github.com/Gurveer1510/url-shortner/pkg/utils"
)

type URLService struct {
	URLRepo *persistance.URLRepo
}

func NewURLService(urlRepo persistance.URLRepo) *URLService {
	return &URLService{URLRepo: &urlRepo}
}

func (u *URLService) CreateShortUrl(urlEntry url.URL) (url.URL, error) {
	randString := utils.GenerateRandomChar()
	urlEntry.ShortURL = randString
	return u.URLRepo.CreatShortUrl(urlEntry)
}

func (u *URLService) GetLongURL(urlReq url.URLRequest) (string, error){
	return u.URLRepo.GetLongUrl(urlReq)
} 