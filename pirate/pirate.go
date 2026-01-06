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

type Saver interface {
	Save(c Pirate) error
}

func New(course, name, date string) (*Pirate, error) {
	if name != strings.ToUpper(name) {
		return nil, errors.New("le nom du pirate doit être en majuscule")
	}

	for _, r := range name { // on parcourt chaque rune (caractère Unicode)
        if unicode.IsDigit(r) {
            return nil,  errors.New("le nom ne doit pas contenir de chiffres")
        }
    }
	
	prime := "1000€"

	pirate := &Pirate{
		Name:  name,
		Prime: prime,
		Img:   "wantedVierge.png",
	}

	return pirate, nil
}
