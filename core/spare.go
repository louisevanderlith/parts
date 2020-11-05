package core

import "github.com/louisevanderlith/husk/validation"

type Spare struct {
	Number string
	Type   string
	Weight int
}

func (o Spare) Valid() error {
	return validation.Struct(o)
}
