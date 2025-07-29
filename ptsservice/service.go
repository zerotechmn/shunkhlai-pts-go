package ptsservice

import (
	"encoding/json"
	"net/http"
)

type PumpHandler struct {
	Service *PTSService
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
	json.NewEncoder(w).Encode(res)
}

// func (h *PumpHandler) AuthorizeByAmount(w http.ResponseWriter, r *http.Request) {
// 	var body struct {
// 		Pump   int     `json:"pump"`
// 		Nozzle int     `json:"nozzle"`
// 		Amount float64 `json:"amount"`
// 		Price  float64 `json:"price"`
// 	}
// 	json.NewDecoder(r.Body).Decode(&body)
// 	res := h.Service.AuthorizeByAmount(body.Pump, body.Nozzle, body.Amount, body.Price)
// 	json.NewEncoder(w).Encode(res)
// }

// func (h *PumpHandler) TransClose(w http.ResponseWriter, r *http.Request) {
// 	pumpStr := mux.Vars(r)["pump"]
// 	pump, _ := strconv.Atoi(pumpStr)
// 	res := h.Service.TransClose(pump)
// 	json.NewEncoder(w).Encode(res)
// }

// func (h *PumpHandler) TotalRequest(w http.ResponseWriter, r *http.Request) {
// 	var body struct {
// 		Pump   int `json:"pump"`
// 		Nozzle int `json:"nozzle"`
// 	}
// 	json.NewDecoder(r.Body).Decode(&body)
// 	res := h.Service.TotalRequest(body.Pump, body.Nozzle)
// 	json.NewEncoder(w).Encode(res)
// }

// func (h *PumpHandler) GetPumpsTotal(w http.ResponseWriter, r *http.Request) {
// 	res := h.Service.GetPumpsTotal()
// 	json.NewEncoder(w).Encode(res)
// }

// func (h *PumpHandler) PricesRequest(w http.ResponseWriter, r *http.Request) {
// 	pumpStr := mux.Vars(r)["pump"]
// 	pump, _ := strconv.Atoi(pumpStr)
// 	res := h.Service.PricesRequest(pump)
// 	json.NewEncoder(w).Encode(res)
// }

// func (h *PumpHandler) PresetPrices(w http.ResponseWriter, r *http.Request) {
// 	pumpStr := mux.Vars(r)["pump"]
// 	pump, _ := strconv.Atoi(pumpStr)
// 	var body struct {
// 		Prices []float64 `json:"prices"`
// 	}
// 	json.NewDecoder(r.Body).Decode(&body)
// 	res := h.Service.PresetPrices(pump, body.Prices)
// 	json.NewEncoder(w).Encode(res)
// }

// func (h *PumpHandler) Statuses(w http.ResponseWriter, r *http.Request) {
// 	res := h.Service.GetStatuses()
// 	json.NewEncoder(w).Encode(res)
// }

// func (h *PumpHandler) Status(w http.ResponseWriter, r *http.Request) {
// 	pumpStr := mux.Vars(r)["pump"]
// 	pump, _ := strconv.Atoi(pumpStr)
// 	res := h.Service.GetStatus(pump)
// 	json.NewEncoder(w).Encode(res)
// }

// func (h *PumpHandler) Version(w http.ResponseWriter, r *http.Request) {
// 	res := h.Service.VersionGet()
// 	json.NewEncoder(w).Encode(res)
// }

// func (h *PumpHandler) LatestError(w http.ResponseWriter, r *http.Request) {
// 	res := h.Service.LatestError()
// 	json.NewEncoder(w).Encode(res)
// }

// func (h *PumpHandler) SetConfig(w http.ResponseWriter, r *http.Request) {
// 	var config map[string]interface{}
// 	json.NewDecoder(r.Body).Decode(&config)
// 	res := h.Service.SetConfig(config)
// 	json.NewEncoder(w).Encode(res)
// }
