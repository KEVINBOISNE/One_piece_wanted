package main

import (
	"fmt"
	"one_piece/GeneratePdf"
	"one_piece/csvparser"
	"one_piece/pirate"
	"flag"
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

		err = GeneratePdf.GeneratePdf(p.Name, c.Prime, "pdf", "")
		if err != nil {
			fmt.Println("Erreur PDF :", err)
			continue
		}

		fmt.Println("PDF généré pour :", p.Name, "avec une prime de", c.Prime)
	}

		name := flag.String("name", "", "Nom du pirate")
		prime := flag.String("prime", "", "Prime du pirate")
		img := flag.String("img", "", "Image personnalisée du pirate")
		output := flag.String("output", "pdf", "Dossier de sortie des PDFs")
		flag.Parse()

		if *name == "" || *prime == "" {
		fmt.Println("Usage : ./MyGeneratorPrime -name \"ZORO\" -prime \"320,000,000\" -img \"zoro.jpg\" -output \"pdf\"")

			return
		}


		err = GeneratePdf.GeneratePdf(*name, *prime, *output, *img)
		if err != nil {
			fmt.Println("Erreur génération PDF :", err)
			return
		}

		fmt.Println("PDF généré pour", *name, "avec une prime de", *prime)
}

