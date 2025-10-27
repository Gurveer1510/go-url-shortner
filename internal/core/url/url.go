package url

import "time"

type URL struct {
	Id        int       `json:"id"`
	LongURL   string    `json:"longUrl`
	ShortURL  string    `json:shortUrl`
	createdAt time.Time `json:"createdAt"`
}

type URLRequest struct {
	ShortURL	string	`json:"shortUrl"`
}