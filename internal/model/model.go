package model

type Model struct {
	ID         uint32 `json:"id"`
	CreatedBy  string `json:"created_by,omitempty"`
	ModifiedBy string `json:"modifie_by,omitempty"`
	CreatedOn  uint32 `json:"created_on,omitempty"`
	ModifiedOn uint32 `json:"modified_on,omitempty"`
	DeletedOn  uint32 `json:"deleted_on,omitempty"`
	IsDel      uint8  `json:"is_del,omitempty"`
}
