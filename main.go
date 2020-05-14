package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
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

func forCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
		return

	})
}

func test(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Methods", "*")
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

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
	router.Use(forCORS)
	router.HandleFunc("/post", test).Methods("POST", "OPTIONS")
	http.ListenAndServe(":8000", router) // for local dev
	//port := os.Getenv("PORT")
	//http.ListenAndServe(":"+port, router)
}
