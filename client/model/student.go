package model

type Student struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    uint   `json:"age"`
	Credit int    `json:"credit"`
}
