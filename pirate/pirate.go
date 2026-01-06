package pirate

import (
	"errors"
	"strings"
	"unicode"
)

type Pirate struct {
	Name  string
	Prime string
	Img   string
}

func New(course, name, date string) (*Pirate, error) {

	if name == "" {
		return nil, errors.New("nom vide")
	}

	if name != strings.ToUpper(name) {
		return nil, errors.New("le nom doit être en majuscule")
	}

	for _, r := range name {
		if unicode.IsDigit(r) {
			return nil, errors.New("le nom ne doit pas contenir de chiffres")
		}
	}

	return &Pirate{
		Name: name,
	}, nil
}
