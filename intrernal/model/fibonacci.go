package model

type Fibonacci struct {
	X uint64 `form:"x" json:"x" binding:"required"`
	Y uint64 `form:"y" json:"y" binding:"required"`
}
