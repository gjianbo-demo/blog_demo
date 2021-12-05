package model

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id,omitempty"`
	ArticleID uint32 `json:"article_id,omitempty"`
}
