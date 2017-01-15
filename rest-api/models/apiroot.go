package models

import "github.com/nvellon/hal"

// APIRoot : Placeholder type for the HAL root of the REST API
type APIRoot struct {
}

func (r APIRoot) GetMap() hal.Entry {
	return hal.Entry{}
}
