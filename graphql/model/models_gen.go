// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthorFilter struct {
	Q         *string   `json:"q"`
	Ids       []*string `json:"ids"`
	ID        *string   `json:"id"`
	CreatedAt *string   `json:"createdAt"`
	UpdatedAt *string   `json:"updatedAt"`
	FullName  string    `json:"fullName"`
	Age       int       `json:"age"`
}

type FilmFilter struct {
	Q           *string   `json:"q"`
	Ids         []*string `json:"ids"`
	ID          *string   `json:"id"`
	CreatedAt   *string   `json:"createdAt"`
	UpdatedAt   *string   `json:"updatedAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type ListMetadata struct {
	Count *int `json:"count"`
}

type NewAuthor struct {
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
}

type NewFilm struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateAuthor struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
}

type UpdateFilm struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
