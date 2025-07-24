package handlers

import (
	"encoding/json"
	"net/http"

	"shunkhlai-pts-go/ptsservice"

	log "github.com/sirupsen/logrus"
)

type PumpHandler struct {
	Service *ptsservice.PTSService
	Log     *log.Logger
}

func NewPumpHandler(service *ptsservice.PTSService, logger *log.Logger) *PumpHandler {
	return &PumpHandler{Service: service, Log: logger}
}

func (h *PumpHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Pump   int     `json:"pump"`
		Nozzle int     `json:"nozzle"`
		Volume float64 `json:"volume"`
		Price  float64 `json:"price"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	res := h.Service.Authorize(body.Pump, body.Nozzle, body.Volume, body.Price)
	respondJSON(w, res)
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}
