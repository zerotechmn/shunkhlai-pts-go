package ptsservice

import (
	"fmt"
	"net"
	"time"

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
	fmt.Printf("Authorize: pump=%d nozzle=%d volume=%.2f price=%.2f", pump, nozzle, volume, price)

	var req BaseRequest

	req.ID = 2
	req.Method = "authorize_pump"
	req.Params.Nozzle = nozzle
	req.Params.Pump = pump
	req.Params.PresetType = "volume"
	req.Params.PresetValue = volume

	fmt.Printf("Authorize Req : %+v\n", req)

	p.log.Infof("Authorize: pump=%d nozzle=%d volume=%.2f price=%.2f", pump, nozzle, volume, price)
	// TODO: Authorize pts2 request shideh bolon sock uusgej status awah

	return map[string]interface{}{"authorized": true}
}

func (h *PTSService) Ping() map[string]string {
	var localIP string
	ifaces, err := net.Interfaces()
	if err == nil {
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ip != nil && !ip.IsLoopback() && ip.To4() != nil {
					localIP = ip.String()
					break
				}
			}
			if localIP != "" {
				break
			}
		}
	}

	return map[string]string{
		"status":  "ok",
		"version": "v1.0.1",
		"pts_ip":  localIP,
		"data":    time.Now().Format(time.RFC3339),
	}
}
