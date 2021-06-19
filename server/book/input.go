package book

type CreateBook struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required"`
}

type UpdateBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
