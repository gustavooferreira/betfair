// Generate betfair enums.
package main

import (
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"os"
)

type EnumsInfo struct {
	Type  string
	Enums map[string]string
}

type EnumsInfoArray []EnumsInfo

func main() {
	file, err := os.Open("assets/enums.csv")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	stage := -1
	results := EnumsInfoArray{}
	var temp EnumsInfo

	for {
		record, err := r.Read()
		if err == io.EOF {
			results = append(results, temp)
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		switch stage {
		case 1:
			if len(record) != 1 {
				log.Fatalf("Error. I don't know what happened!")
			}
			temp.Type = record[0]
			stage = 2
		case 2:
			if len(record) != 2 || record[0] != "Value" || record[1] != "Description" {
				log.Fatalf("Error. I don't know what happened!")
			}
			stage = 3
		case 3:
			if len(record) == 2 {
				temp.Enums[record[0]] = record[1]
			} else if len(record) == 1 {
				results = append(results, temp)
				temp = EnumsInfo{Enums: map[string]string{}}
				temp.Type = record[0]
				stage = 2
			} else {
				log.Fatalf("Error. I don't know what happened!")
			}
		default:
			temp = EnumsInfo{Enums: map[string]string{}}
			if len(record) != 1 {
				log.Fatalf("Error. I don't know what happened!")
			}
			temp.Type = record[0]
			stage = 2
		}
	}

	// fmt.Printf("%+v\n", results)

	fOut, err := os.Create("enums.go")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	tmpl := template.Must(template.ParseFiles("assets/templates/enums.go.tmpl"))

	data := struct{ Results EnumsInfoArray }{Results: results}
	tmpl.Execute(fOut, data)

	fOut.Close()
}
