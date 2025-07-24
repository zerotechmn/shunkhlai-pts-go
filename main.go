package main

import (
	"fmt"
	"net/http"

	"github.com/Owner/shunkhlai-pts-go/handlers"
	"github.com/Owner/shunkhlai-pts-go/logger"
	"github.com/Owner/shunkhlai-pts-go/middleware"
	"github.com/Owner/shunkhlai-pts-go/ptsservice"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Амжилттай холбогдлоо! IP: %s", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", handler)
	log := logger.InitLogger()
	service := ptsservice.NewPTSService(log)
	h := handlers.NewPumpHandler(service, log)

	r := mux.NewRouter()
	r.Use(middleware.RecoverMiddleware)
	pump := r.PathPrefix("/pump").Subrouter()

	pump.HandleFunc("/authorize", h.Authorize).Methods("POST")
	pump.HandleFunc("/authorize_by_amount", h.AuthorizeByAmount).Methods("POST")
	pump.HandleFunc("/{pump:[0-9]+}/trans_close", h.TransClose).Methods("GET")
	pump.HandleFunc("/total_request", h.TotalRequest).Methods("POST")
	pump.HandleFunc("/pumps_total", h.GetPumpsTotal).Methods("GET")
	pump.HandleFunc("/{pump:[0-9]+}/prices", h.PricesRequest).Methods("GET")
	pump.HandleFunc("/{pump:[0-9]+}/prices", h.PresetPrices).Methods("POST")
	pump.HandleFunc("/statuses", h.Statuses).Methods("POST")
	pump.HandleFunc("/{pump:[0-9]+}/status", h.Status).Methods("GET")
	pump.HandleFunc("/version", h.Version).Methods("GET")
	pump.HandleFunc("/error", h.LatestError).Methods("GET")
	pump.HandleFunc("/config", h.SetConfig).Methods("POST")
	pump.HandleFunc("/ping", h.Ping).Methods("GET")
	pump.HandleFunc("/shutdown", h.Shutdown).Methods("GET")

	log.Info("Starting server at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
