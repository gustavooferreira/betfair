// Generate betfair enums.
package main

import (
	"bytes"
	"encoding/csv"
	"go/format"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"
)

type EnumsInfo struct {
	Type      string // example: MarketProjection
	TypeCamel string // example: marketProjection
	VarName   string // example: mp

	Enums           []string // example: EVENT_TYPE
	EnumsPascalCase []string // example: EventType
}

type EnumsInfoArray []EnumsInfo

func main() {
	// Revise this
	// filePath := os.Args[1]
	filePath := "assets/enums.csv"

	results := readCSV(filePath)

	addExtraTransformations(&results)

	// fmt.Printf("%+v\n", results)

	// Generate from template
	buf := genCode(results)

	fOut, err := os.Create("enums.go")
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	result, err := format.Source(buf.Bytes())

	// Write result to file fOut
	// _, err = fOut.Write(buf.Bytes())
	_, err = fOut.Write(result)
	if err != nil {
		log.Fatalf("error: %s\n", err)
		return
	}

	fOut.Close()
}

func readCSV(filePath string) EnumsInfoArray {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	stage := 1
	results := EnumsInfoArray{}
	var temp EnumsInfo

	for {
		record, err := r.Read()
		if err == io.EOF {
			results = append(results, temp)
			break
		} else if err != nil {
			log.Fatal(err)
		}

		switch stage {
		case 2:
			if len(record) != 2 || record[0] != "Value" || record[1] != "Description" {
				log.Fatalf("error: expecting 'Value,Description', got instead: %s\n", record)
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
				log.Fatalf("error: expecting 'enumValue,enumDescription', got instead: %s\n", record)
			}
		default:
			temp = EnumsInfo{Enums: []string{}}
			if len(record) != 1 {
				log.Fatalf("error: Expecting first line of the file to start with the first Enum type!\n")
			}
			temp.Type = record[0]
			stage = 2
		}
	}

	return results
}

func addExtraTransformations(data *EnumsInfoArray) {

	for i, elem := range *data {
		typeRune := []rune(elem.Type)
		temp := append([]rune{unicode.ToLower(typeRune[0])}, typeRune[1:]...)
		(*data)[i].TypeCamel = string(temp)

		varName := []rune{}
		for _, elemType := range []rune(elem.Type) {
			if unicode.IsUpper(elemType) {
				varName = append(varName, elemType)
			}
		}
		(*data)[i].VarName = strings.ToLower(string(varName))

		(*data)[i].EnumsPascalCase = []string{}
		for _, elemEnums := range elem.Enums {
			// convert _ to space and apply string.title then remove spaces
			temp := strings.Replace(elemEnums, "_", " ", -1)
			temp = strings.ToLower(temp)
			temp = strings.Title(temp)
			temp = strings.Replace(temp, " ", "", -1)

			(*data)[i].EnumsPascalCase = append((*data)[i].EnumsPascalCase, temp)
		}
	}
}

func genCode(data EnumsInfoArray) *bytes.Buffer {
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

	// tmpl = tmpl.Funcs(template.FuncMap{
	// 	"gfTitle": func(str string) string {
	// 		// return strings.Title(strings.ReplaceAll(str, "_", " "))
	// 		// return strings.ToLower(str)
	// 		return strings.ReplaceAll(strings.Title(strings.ToLower(strings.ReplaceAll(str, "_", " "))), " ", "")
	// 	},
	// })

	tmpl, err := tmpl.ParseFiles("assets/templates/enums.go.tmpl")
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	// var b bytes.Buffer
	buf := bytes.NewBuffer([]byte{})

	err = tmpl.Execute(buf, data)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	return buf
}
