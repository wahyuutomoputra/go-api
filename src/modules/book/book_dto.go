package book

type CreateBook struct {
	Name   string `json:"name" validate:"required"`
	Author string `json:"name validate:"required""`
}
