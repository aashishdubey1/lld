package ocp

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

type Exporter interface {
	Export(data []Data) error
}

type CSVExporter struct{}
type JSONExporter struct{}
type PDFExporter struct{}

func (c *CSVExporter) Export(data []Data) error {
	file, err := os.Create("users.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"Name", "Age", "Email"}); err != nil {
		return err
	}
	for _, v := range data {
		if err := writer.Write([]string{v.Name, strconv.Itoa(v.Age), v.Email}); err != nil {
			return err
		}
	}
	return nil
}

func (p *PDFExporter) Export(data []Data) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Users")
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(50, 10, "Name")
	pdf.Cell(30, 10, "Age")
	pdf.Cell(80, 10, "Email")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	for _, user := range data {
		pdf.Cell(50, 10, user.Name)
		pdf.Cell(30, 10, strconv.Itoa(user.Age))
		pdf.Cell(80, 10, user.Email)
		pdf.Ln(10)
	}
	return pdf.OutputFileAndClose("users.pdf")
}

func (j *JSONExporter) Export(data []Data) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("users.json", jsonData, 0644)
}

func ProcessExport(exporter Exporter, data []Data) error {
	return exporter.Export(data)
}
