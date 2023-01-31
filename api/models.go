package api

type Member struct {
	ID         int    `json:"user_id"`
	Name       string `json:"user_name"`
	Birthday   string `json:"birthday"`
	Height     string `json:"height"`
	Blood      string `json:"blood_type"`
	Generation string `json:"generation"`
	BlogURL    string `json:"blog_url"`
	ImgURL     string `json:"img_url"`
}

// Group represents group information.
type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
