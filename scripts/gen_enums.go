// Generate betfair enums.
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

type EnumsInfo struct {
	Type  string
	Enums []string
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
				temp.Enums = append(temp.Enums, record[0])
			} else if len(record) == 1 {
				results = append(results, temp)
				temp = EnumsInfo{Enums: []string{}}
				temp.Type = record[0]
				stage = 2
			} else {
				log.Fatalf("Error. I don't know what happened!")
			}
		default:
			temp = EnumsInfo{Enums: []string{}}
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

	// tmpl := template.Must(template.ParseFiles("assets/templates/enums.go.tmpl").Funcs(template.FuncMap{
	// 	"gfTitle": func(str string) string {
	// 		return strings.Title(str)
	// 	},
	// }))

	// tmpl := template.Must(template.New("assets/templates/enums.go.tmpl").Funcs(template.FuncMap{
	// 	"gfTitle": func(str string) string {
	// 		return strings.Title(str)
	// 	},
	// }).ParseFiles("assets/templates/enums.go.tmpl"))

	// tmpl := template.Must(template.ParseFiles("assets/templates/enums.go.tmpl"))

	tmpl := template.New("enums.go.tmpl")

	tmpl = tmpl.Funcs(template.FuncMap{
		"gfTitle": func(str string) string {
			// return strings.Title(strings.ReplaceAll(str, "_", " "))
			// return strings.ToLower(str)
			return strings.ReplaceAll(strings.Title(strings.ToLower(strings.ReplaceAll(str, "_", " "))), " ", "")
		},
	})

	tmpl, err = tmpl.ParseFiles("assets/templates/enums.go.tmpl")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// var b bytes.Buffer
	buf := bytes.NewBuffer([]byte{})

	err = tmpl.Execute(buf, results)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// result, err := format.Source(buf.Bytes())
	// Write result to file fOut

	_, err = fOut.Write(buf.Bytes())
	// _, err = fOut.Write(result)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fOut.Close()
}
