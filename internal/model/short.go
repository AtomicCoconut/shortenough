package model

// Short represents a shortened url
type Short struct {
	// The shortened url
	ShortUrl string `json:"short_url"`

	// The original url
	RedirectUrl string `json:"redirect_url"`

	// The amount of uses this short url has
	Uses uint32 `json:"uses"`
}
