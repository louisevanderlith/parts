package core

import "github.com/louisevanderlith/husk/validation"

type Spare struct {
	Number    string
	Type      string
	Weight    int
	AppliesTo []string `hsk:"null"`
}

func (o Spare) Valid() error {
	return validation.Struct(o)
}
