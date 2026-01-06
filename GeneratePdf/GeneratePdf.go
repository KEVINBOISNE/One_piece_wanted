package GeneratePdf

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/otaviobaldan/go-pdf-generator/config"
	"github.com/otaviobaldan/go-pdf-generator/constants"
)

type PdfGenerator struct {
	Pdf            *gofpdf.Fpdf
	isPointUnit    bool
	TxtCfgHeader   *config.TextConfig
	TxtCfgFooter   *config.TextConfig
	TxtCfgTitle    *config.TextConfig
	TxtCfgSubtitle *config.TextConfig
	TxtCfgText     *config.TextConfig
}
func getImageByName(name string) string {
	switch strings.ToUpper(name) {
	case "LUFFY":
		return "gear-5-luffy.jpg"
	case "BAGGY":
		return "baggy.jpg"
	case "SHANKS":
		return "shanks.jpg"
	case "MARSHALL D.TEACH":
		return "teach.jpg"
	default:
		return "wantedVierge.jpg"
	}
}


func GeneratePdf(name string, prime string, outputDir string, imgFile string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.Image("wantedVierge.jpg", 20, 20, 170, 0, false, "", 0, "")

	pdf.SetFont("Arial", "B", 22)
	pdf.SetXY(90, 170)
	pdf.Cell(40, 80, name)

	pdf.SetFont("Arial", "", 18)
	pdf.SetXY(60, 190)
	pdf.Cell(40, 80, "Prime : "+prime)

	imagePath := imgFile
	if imagePath == "" {
		imagePath = getImageByName(name)
	}

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		fmt.Println("Image non trouvée, utilisation du fond vierge :", imagePath)
	} else {
		pdf.Image(imagePath, 42, 85, 127, 0, false, "", 0, "")
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("impossible de créer le dossier %s: %v", outputDir, err)
		}
	}

	// Nom du fichier PDF
	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_wanted.pdf", strings.ReplaceAll(name, " ", "_")))

	return pdf.OutputFileAndClose(outputFile)
}

func NewPdfGenerator(
	config *config.PdfConfig,
	TxtCfgHeader *config.TextConfig,
	TxtCfgFooter *config.TextConfig,
	TxtCfgTitle *config.TextConfig,
	TxtCfgSubtitle *config.TextConfig,
	TxtCfgText *config.TextConfig,
) (pdfGenerator *PdfGenerator, err error) {
	pdfGenerator = &PdfGenerator{
		Pdf:            gofpdf.New(config.Orientation, config.Units, config.PaperSize, ""),
		isPointUnit:    config.Units == constants.UnitsPoints,
		TxtCfgHeader:   TxtCfgHeader,
		TxtCfgTitle:    TxtCfgTitle,
		TxtCfgSubtitle: TxtCfgSubtitle,
		TxtCfgText:     TxtCfgText,
		TxtCfgFooter:   TxtCfgFooter,
	}

	if config.RegisterFonts {
		err = config.RegisterExternalFonts(pdfGenerator.Pdf)
		if err != nil {
			return nil, errors.New("error loading fonts")
		}
	}

	margins := config.Margins
	pdfGenerator.Pdf.SetMargins(margins.Left, margins.Top, margins.Right)
	pdfGenerator.Pdf.SetAutoPageBreak(true, pdfGenerator.calculateSize(40))
	pdfGenerator.Pdf.AddPage()

	return pdfGenerator, nil
}

func (pg *PdfGenerator) GenerateDefaultHeader(headerText string) {
	cfg := pg.TxtCfgHeader
	color := cfg.Color
	pg.Pdf.SetHeaderFunc(func() {
		pg.Pdf.SetFont(cfg.FontFamily, cfg.Style, cfg.Size)
		pg.Pdf.SetTextColor(color.R, color.G, color.B)

		stringWidth := pg.calculateSize(pg.Pdf.GetStringWidth(headerText) + 6)
		width, _ := pg.Pdf.GetPageSize()
		pg.Pdf.SetX((width - stringWidth) / 2)

		pg.Pdf.CellFormat(stringWidth, 9, headerText, "", 0, cfg.Align, false, 0, "")
	
		pg.Pdf.Ln(pg.calculateSize(10))
	})
}

func (pg *PdfGenerator) GenerateDefaultFooter(text string, pageNumber bool) {
	cfg := pg.TxtCfgFooter
	color := cfg.Color
	pg.Pdf.SetFooterFunc(func() {

		pg.Pdf.SetY(pg.calculateSize(-15))

		pg.Pdf.SetFont(cfg.FontFamily, cfg.Style, cfg.Size)
		pg.Pdf.SetTextColor(color.R, color.G, color.B)
		pg.Pdf.CellFormat(0, pg.calculateSize(10), text,
			"", 0, cfg.Align, false, 0, "")

		if pageNumber {
			pg.Pdf.SetTextColor(0, 0, 0)
			pg.Pdf.CellFormat(0, pg.calculateSize(10), fmt.Sprintf("Pág. %d", pg.Pdf.PageNo()),
				"", 0, constants.AlignRight, false, 0, "")
		}
	})
}

func (pg *PdfGenerator) GenerateTitle(title string) {
	cfg := pg.TxtCfgTitle
	color := cfg.Color
	pg.Pdf.SetFont(cfg.FontFamily, cfg.Style, cfg.Size)
	pg.Pdf.SetTextColor(color.R, color.G, color.B)
	pg.Pdf.CellFormat(0, pg.calculateSize(constants.SizeTitleHeight), title,
		"", 1, cfg.Align, false, 0, "")

	pg.Pdf.Ln(constants.SizeLineBreak)
}

func (pg *PdfGenerator) GenerateSubtitle(subtitle string) {
	cfg := pg.TxtCfgSubtitle
	color := cfg.Color

	pg.Pdf.SetFont(cfg.FontFamily, cfg.Style, cfg.Size)
	pg.Pdf.SetTextColor(color.R, color.G, color.B)
	pg.Pdf.CellFormat(0, pg.calculateSize(constants.SizeSubTitleHeight), subtitle,
		"", 1, cfg.Align, false, 0, "")

	pg.Pdf.Ln(pg.calculateSize(constants.SizeLineBreak))
}

func (pg *PdfGenerator) GenerateText(text string) {
	cfg := pg.TxtCfgText
	color := cfg.Color

	pg.Pdf.SetFont(cfg.FontFamily, cfg.Style, cfg.Size)
	pg.Pdf.SetTextColor(color.R, color.G, color.B)

	text = strings.ReplaceAll(text, `\n`, "\n")

	pg.Pdf.MultiCell(0, pg.calculateSize(constants.SizeTextHeight), text, "", cfg.Align, false)

	pg.Pdf.Ln(-1)
}

func (pg *PdfGenerator) GenerateSignature(signatureName string) {
	currentY := pg.Pdf.GetY()
	left, _, right, _ := pg.Pdf.GetMargins()
	width, _ := pg.Pdf.GetPageSize()

	lineSize := pg.calculateSize(130)
	availableSpace := (width - left - right - lineSize) / 2
	lineY := currentY + pg.calculateSize(20)
	lineInit := left + availableSpace
	lineEnd := left + availableSpace + lineSize

	pg.Pdf.Line(lineInit, lineY, lineEnd, lineY)
	pg.Pdf.CellFormat(0, pg.calculateSize(50), signatureName, "", 1, "C", false, 0, "")
}

func (pg *PdfGenerator) calculateSize(size float64) float64 {
	if pg.isPointUnit {
		return size * 2.834
	}
	return size
}
