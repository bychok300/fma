package dtos

type Result struct {
	Result string `json:"result"`
}

type LimitOffsetDto struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

type RatePostDto struct {
	Id int `json:"id"`
	Rate int `json:"rate"`
}
