package dto

type SetCount struct {
	Count int `json:"count" validate:"required,min=0"`
}