package model

type Article struct {
	*Model
	Title          string `json:"title"`
	Desc           string `json:"desc"`
	Content        string `json:"content"`
	ConverImageUrl string `json:"cover_image_url"`
	State          uint8  `json:"state"`
}

func (t Article) TableName() string {
	return "blog_article"
}
