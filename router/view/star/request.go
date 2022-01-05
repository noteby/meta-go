package star

type CreateStarRequest struct {
	Title   string `validate:"required,max=36"`
	Content string `validate:"required"`
}

type UpdateStarRequest struct {
	ID      uint   `validate:"required"`
	Title   string `validate:"required,max=36"`
	Content string `validate:"required"`
}

type StarIDRequest struct {
	ID uint `validate:"required"`
}

type StarListRequest struct {
	Offset int `validate:"gte=0"`
	Limit  int `validate:"gte=0,max=100"`
}
