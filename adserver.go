package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type AdDetail struct {
	ID             string    `json:"id"`
	Price          float32   `json:"price"`
	ExpirationDate time.Time `json:"expiration_date"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the CSV file.")
	}
	filename := os.Args[1]
	data := readCSV(filename)
	handler := func(w http.ResponseWriter, r *http.Request) {
		adID := strings.TrimPrefix(r.URL.Path, "/promotions/")
		if val, ok := data[adID]; ok {
			jsonData, err := json.Marshal(val)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		} else {
			http.Error(w, "Ad not found", http.StatusNotFound)
		}
	}
	http.HandleFunc("/promotions/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readCSV(filename string) map[string]AdDetail {
	data := make(map[string]AdDetail)
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	reader := csv.NewReader(csvFile)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error parsing data", err)
		}
		price, err := strconv.ParseFloat(line[1], 32)
		if err != nil {
			log.Println("Error parsing data", err)
			continue
		}
		expirationDate, err := time.Parse(time.RFC3339, line[2])
		if err != nil {
			log.Println("Error parsing data", err)
			continue
		}
		data[line[0]] = AdDetail{
			ID:             line[0],
			Price:          float32(price),
			ExpirationDate: expirationDate,
		}
	}
	return data
}
