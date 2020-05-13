package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Wire struct {
	Predia string
	Prelen string
	Fabdia string
}
type FWire struct {
	Predia float64
	Prelen float64
	Fabdia float64
}
type After struct {
	Fablen string
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// jsonファイルを受け取る、そしてWire構造体にパース
	var data Wire
	jsondata, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(jsondata, &data)
	fmt.Printf("%+v\n", data)
	if err != nil {
		log.Fatal(err)
	}
	// contert string data received to float 64 for calculate
	var fdata FWire

	// Convert string to float64 and insert struct
	fdataPredia, _ := strconv.ParseFloat(data.Predia, 64)
	fdata.Predia = fdataPredia

	fdataPrelen, _ := strconv.ParseFloat(data.Prelen, 64)
	fdata.Prelen = fdataPrelen

	fdataFabdia, _ := strconv.ParseFloat(data.Fabdia, 64)
	fdata.Fabdia = fdataFabdia
	fmt.Printf("%+v\n", fdata) // for debug
	fmt.Print(fdata.Fabdia)
	// caluclate the length after fabrication
	var after After
	ffablen := (math.Pow(fdataPredia, 2) / math.Pow(fdataFabdia, 2)) * fdataPrelen
	// Convert float64 to string
	sfablen := strconv.FormatFloat(ffablen, 'f', 4, 64)
	after.Fablen = sfablen
	jsonOUT, _ := json.Marshal(after)
	w.Write(jsonOUT)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", test).Methods("POST")
	//http.ListenAndServe(":8000", router) // for local dev
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}
