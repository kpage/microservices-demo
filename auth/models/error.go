package models

import (
	"github.com/nvellon/hal"
)

type (
	Error struct {
		Message string
	}
)

func (a Error) GetMap() hal.Entry {
	return hal.Entry{
		"message": a.Message,
	}
}
