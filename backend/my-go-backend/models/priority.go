package models

import (
	"errors"
)

type Priority struct {
	value int
}

func NewPriority(v int) Priority {
	return Priority{value: v}
}

func (p Priority) Value() int {
	return p.value
}

func (p *Priority) Set(v int) error {
	if v < 1 || v > 5{
		return errors.New("priority must be between 1 and 5")
	}
	p.value = v
	return nil
}