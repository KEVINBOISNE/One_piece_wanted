package main

import (
	"fmt"
	"one_piece/GeneratePdf"
	"one_piece/csvparser"
	"one_piece/pirate"
)

func main() {
	fmt.Println("Hello Pirate World!")

	primes, err := csvparser.Parse("wanted.csv")
	if err != nil {
		fmt.Println("Erreur CSV :", err)
		return
	}

	for _, c := range primes {

		p, err := pirate.New("", c.Name, "")
		if err != nil {
			fmt.Println("Pirate invalide :", err)
			continue
		}

		err = GeneratePdf.GeneratePdf(p.Name, c.Prime)
		if err != nil {
			fmt.Println("Erreur PDF :", err)
			continue
		}

		fmt.Println("PDF généré pour :", p.Name, "avec une prime de", c.Prime)
	}
}
