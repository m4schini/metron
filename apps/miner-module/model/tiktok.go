package model

type Account struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Bio         string `json:"bio"`
	Following   int    `json:"following"`
	Followers   int    `json:"followers"`
	Likes       int    `json:"likes"`
	AvatarURL   string `json:"avatar"`
}

type Video struct {
	URL          string `json:"url"`
	ID           string `json:"id"`
	Username     string `json:"username"`
	VideoURL     string `json:"video"`
	Timestamp    string `json:"timestamp"`
	ThumbnailURL string `json:"thumbnail"`

	Views       int    `json:"views"`
	Likes       int    `json:"likes"`
	Comments    int    `json:"comments"`
	Shares      int    `json:"shares"`
	Audio       string `json:"audio"`
	VideoLength int    `json:"videoLength"`
	Description string `json:"description"`
}
