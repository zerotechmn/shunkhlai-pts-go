package ptsservice

type BaseRequest struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
	Params struct {
		Pump        int     `json:"pump"`
		Nozzle      int     `json:"nozzle"`
		PresetType  string  `json:"preset_type"`
		PresetValue float64 `json:"preset_value"`
	}
}
