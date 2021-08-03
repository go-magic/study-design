package model

type AddTask struct {
	Args []int `json:"args"`
}

type AddResult struct {
	Sum int `json:"sum"`
}
