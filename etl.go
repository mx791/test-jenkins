package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
	"net/http"
	"io"
	"strings"

)

var url string = "https://public.opendatasoft.com/explore/dataset/donnees-synop-essentielles-omm/download/?format=csv&refine.nom_epci=M%C3%A9tropole+du+Grand+Paris&timezone=Europe/Berlin&lang=fr&use_labels_for_header=true&csv_separator=%3B"
var path string = "datas.csv"	

type Row struct {
	Name string
	Temp float64
	Month int 
}

var dataset = make([]Row, 0)

func LoadRow(datas []string) {
	tmp, err := strconv.ParseFloat(datas[64], 64)
	name := datas[77]
	date, err1 := strconv.Atoi(datas[81])
	if err == nil && err1 == nil {
		dataset = append(dataset, Row{name, tmp, date})
	}
}

func LoadFile() {
	fmt.Println("loading datas")
	csvFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
    
	csvr := csv.NewReader(csvFile)
	csvr.FieldsPerRecord = -1
	csvr.Comma = ';'
	csvLine, err := csvr.Read()

    for err == nil {
		LoadRow(csvLine)
		csvLine, err = csvr.Read()
	}

}

var stats = make(map[string]float64)
var statsCount = make(map[string]float64)

func Transform() {
	for _, obj := range dataset {
		key := obj.Name + ":" + []string{"0","1","2","3","4","5","6","7","8","9","10","11","12"}[obj.Month]
		val, exist := stats[key]
		if !exist {
			stats[key] = 0.0
			statsCount[key] = 0.0
		}
		stats[key] = val + obj.Temp
		statsCount[key] += 1.0
	}
}

func Load() {
	var sb strings.Builder
	sb.WriteString("<table><tr><td>Nom du departement</td><td>Numéro de mois</td><td>Températue moyenne</td></tr>")
	for key, value := range stats {
		sb.WriteString("<tr><td>" + key + "</td><td></td><td>" + fmt.Sprintf("%f", value / statsCount[key]) + "</td></tr>")
	}
	sb.WriteString("</table>")

	f, _ := os.Create("./datas/data.txt")
    defer f.Close()

    _, err2 := f.WriteString(sb.String())
}

func Download() {

	fmt.Println("downloading datas")

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
}


func main() {
	Download()
	LoadFile()
	fmt.Println(len(dataset))
	Transform()
	Load()
}