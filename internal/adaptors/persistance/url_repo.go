package persistance

import (
	"log"

	"github.com/Gurveer1510/url-shortner/internal/core/url"
)

type URLRepo struct {
	db *Database		
}

func NewURLRepo(db *Database) *URLRepo {
	return &URLRepo{db: db}
}

func (u *URLRepo) CreatShortUrl(urlEntry url.URL) (url.URL, error) {
	newURL := url.URL{}
	query := `INSERT INTO urls(url, short_url) VALUES($1, $2) RETURNING id, short_url`
	err := u.db.db.QueryRow(query, urlEntry.LongURL, urlEntry.ShortURL).Scan(&newURL.Id, &newURL.ShortURL)
	if err != nil {
		log.Printf("REPO LAYER ERROR: %s", err)
		return newURL, err
	}
	return newURL, nil
}

func (u *URLRepo) GetLongUrl(urlReq url.URLRequest) (string, error) {
	query := "SELECT url from urls where short_url=$1"
	var longURL string
	err := u.db.db.QueryRow(query, urlReq.ShortURL).Scan(&longURL)
	if err != nil {
		log.Printf("REPO LAYER ERROR: %s", err)
		return "", err
	}
	return longURL, nil
}