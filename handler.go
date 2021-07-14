package genchargingstation

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler func(data *ChargingData)

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	datas := strings.Split(string(body), ",")
	if len(datas) < 10 {
		log.Printf("Error reading body")
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	cData := &ChargingData{
		StationID: datas[0],
		BatteryID: datas[1],
	}

	var (
		tmpF  float64
		tmpI  int64
		tmpUi uint64
	)
	tmpF, _ = strconv.ParseFloat(datas[2], 32)
	cData.Volt = float32(tmpF)
	tmpF, _ = strconv.ParseFloat(datas[3], 32)
	cData.Curr = float32(tmpF)
	tmpF, _ = strconv.ParseFloat(datas[4], 32)
	cData.Temp = float32(tmpF)
	tmpF, _ = strconv.ParseFloat(datas[5], 32)
	cData.Capc = float32(tmpF)
	tmpF, _ = strconv.ParseFloat(datas[6], 32)
	cData.SOC = float32(tmpF)
	tmpF, _ = strconv.ParseFloat(datas[7], 32)
	cData.SOH = float32(tmpF)
	tmpI, _ = strconv.ParseInt(datas[8], 10, 32)
	cData.Cycle = int(tmpI)
	tmpUi, _ = strconv.ParseUint(datas[9], 10, 8)
	cData.Error = uint8(tmpUi)

	h(cData)
}
