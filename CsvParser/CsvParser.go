package csvparser


import (
	"encoding/csv"
	"errors"
	"os"
)

type CsvPrime struct {
	Name  string
	Prime string
	Img   string
}

// Parse lit un fichier CSV et retourne une liste de primes
func Parse(path string) ([]CsvPrime, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.LazyQuotes = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, errors.New("fichier CSV vide")
	}

	var primes []CsvPrime

	for _, record := range records {

		if len(record) != 3 {
			return nil, errors.New("ligne CSV invalide")
		}

		primes = append(primes, CsvPrime{
			Name:  record[0],
			Prime: record[1],
			Img:   record[2],
		})
	}

	return primes, nil
}
