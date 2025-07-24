package pts

import (
	log "github.com/sirupsen/logrus"
)

type PTSService struct {
	isPTSOpen bool
	log       *log.Logger
}

func NewPTSService(logger *log.Logger) *PTSService {
	return &PTSService{log: logger}
}

func (p *PTSService) Authorize(pump, nozzle int, volume, price float64) map[string]interface{} {
	p.log.Infof("Authorize: pump=%d nozzle=%d volume=%.2f price=%.2f", pump, nozzle, volume, price)
	return map[string]interface{}{"authorized": true}
}
