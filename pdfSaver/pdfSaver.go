package pdfSaver

import (
	"fmt"
	"one_piece/cert"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputDir string) (*PdfSaver, error) {

	pirate := &PdfSaver{
		OutputDir: "one_piece/pdfSaver",
	}

	return pirate, nil
}

// func GeneratePdf() {
// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.AddPage()
// 	pdf.SetFont("Arial", "", 14)
// 	pdf.Cell(40, 10, "Hello PDF")
// 	pdf.Image("wantedVierge.jpg", 10, 20, 50, 0, false, "", 0, "")
// 	pdf.OutputFileAndClose("test.pdf")
// }

func (p *PdfSaver) Save(c cert.Cert) error {

	filePath := p.OutputDir + "/" + c.Name + ".pdf"
	fmt.Println("Fichier PDF créé :", filePath)
	return nil
}
