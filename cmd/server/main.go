package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nguyenhoaibao/seatsvc"
)

var (
	rows        = flag.Uint("rows", 60, "Number of rows")
	seatsPerRow = flag.Uint("cols", 10, "Number of seats per row")
)

func main() {
	flag.Parse()

	svc := seatsvc.New(*rows, *seatsPerRow)

	r := mux.NewRouter()
	r.HandleFunc("/name", SeatNameHandler(svc)).Queries("number", "{number:[0-9]*?}").Methods("GET").Methods("GET")
	r.HandleFunc("/dimensions", SeatDimensionHandler(svc)).Queries("rows", "{rows}", "cols", "{cols}").Methods("POST")
	http.Handle("/", r)
	log.Println("Server started on port 8888...")
	if err := http.ListenAndServe(":8888", r); err != nil {
		log.Fatal(err)
	}
}

// SeatNameHandler returns the seat name.
func SeatNameHandler(svc *seatsvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		number := mux.Vars(r)["number"]
		seat, _ := strconv.Atoi(number)
		name := svc.SeatName(uint(seat))

		type response struct {
			Name string `json:"name"`
		}
		resp := response{name}

		_ = json.NewEncoder(w).Encode(resp)
	}
}

// SeatDimensionHandler sets the seat dimensions.
func SeatDimensionHandler(svc *seatsvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			vars = mux.Vars(r)

			rows = vars["rows"]
			cols = vars["cols"]

			nRows = strings.Split(rows, ",")
			nCols = strings.Split(cols, ",")
		)

		if err := svc.SetDimensions(nRows, nCols); err != nil {
			w.WriteHeader(400)
			rep := reply{err.Error()}
			_ = json.NewEncoder(w).Encode(rep)
			return
		}

		w.WriteHeader(200)
		rep := reply{"set dimensions successfully"}
		_ = json.NewEncoder(w).Encode(rep)
	}
}

type reply struct {
	Message string `json:"message"`
}
