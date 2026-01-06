package main

import (
	"fmt"
	"one_piece/pdfSaver"
	"one_piece/pirate"
	"one_piece/GeneratePdf"
	"one_piece/csvparser"
)

func main() {
	fmt.Println("Hello Pirate World!")

	course := "Math"
	name := "COOL"
	date := "2026-01-05"

	p, err := pirate.New(course, name, date)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("Wanted pirate créé :", p.Name, "avec une prime de", p.Prime)

	// Génération du PDF
	err = GeneratePdf.GeneratePdf(p.Name, p.Prime)
	if err != nil {
		fmt.Println("Erreur PDF:", err)
		return
	}

	saver := pdfSaver.PdfSaver{OutputDir: "PDFs"}
	fmt.Println("Les PDF seront enregistrés dans :", saver.OutputDir)

		primes, err := csvparser.Parse("wanted.csv")
	if err != nil {
		fmt.Println("Erreur CSV :", err)
		return
	}

	for _, p := range primes {
		fmt.Println("Pirate :", p.Name)
		fmt.Println("Prime :", p.Prime)
		fmt.Println("Image :", p.Img)
		fmt.Println("----")
	}

}

