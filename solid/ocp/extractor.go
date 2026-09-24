package ocp

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

type Data struct {
	Name  string
	Age   int
	Email string
}

func ExtractContent(data []Data, format string) error {
	switch format {
	case "csv":
		file, err := os.Create("users.csv")
		if err != nil {
			log.Fatalf("eror creating file : %v", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.Write([]string{"Name", "Age", "Email"})
		for _, v := range data {
			writer.Write([]string{v.Name, strconv.Itoa(v.Age), v.Email})
		}
	case "pdf":
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
	case "json":
		jsonData, err := json.MarshalIndent(data, "", " ")
		if err != nil {
			log.Fatalf("error converting to json : %v", err)
		}
		err = os.WriteFile("users.json", jsonData, 0644)
		if err != nil {
			log.Fatalf("error writing jsonfile: %v", err)
		}
	}
	return nil
}
