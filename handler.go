package genchargingstation

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler func(data *ChargingData)

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var grad byte = 2

	encBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	body := make([]byte, len(encBody))
	for i, v := range encBody {
		body[i] = v - grad
	}
	// trueData := "1,7456,54.2,10.2,58.92,12.2,77,83,423,1,-7.433639308019364,112.69351752859586"
	datas := strings.Split(string(body), ",")
	if len(datas) < 12 {
		log.Printf("Error reading body")
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	cData := &ChargingData{}

	var (
		tmpF  float64
		tmpI  int64
		tmpUi uint64
	)
	tmpI, _ = strconv.ParseInt(datas[0], 10, 32)
	cData.Hole = uint8(tmpI)
	tmpI, _ = strconv.ParseInt(datas[1], 10, 32)
	cData.BatteryID = uint32(tmpI)
	tmpF, _ = strconv.ParseFloat(datas[2], 32)
	cData.Volt = float32(tmpF)
	tmpF, _ = strconv.ParseFloat(datas[3], 32)
	cData.Curr = float32(tmpF)
	tmpF, _ = strconv.ParseFloat(datas[4], 32)
	cData.Temp = float32(tmpF)
	tmpF, _ = strconv.ParseFloat(datas[5], 32)
	cData.Capc = float32(tmpF)
	tmpI, _ = strconv.ParseInt(datas[6], 10, 32)
	cData.SOC = int(tmpI)
	tmpI, _ = strconv.ParseInt(datas[7], 10, 32)
	cData.SOH = int(tmpI)
	tmpI, _ = strconv.ParseInt(datas[8], 10, 32)
	cData.Cycle = int(tmpI)
	tmpUi, _ = strconv.ParseUint(datas[9], 10, 8)
	cData.Error = uint8(tmpUi)
	tmpF, _ = strconv.ParseFloat(datas[10], 32)
	cData.LocLong = tmpF
	tmpF, _ = strconv.ParseFloat(datas[11], 32)
	cData.LocLat = tmpF
	cData.StationID = fmt.Sprintf("%f, %f", cData.LocLong, cData.LocLat)

	fmt.Fprintf(w, "SUCCESS")
	h(cData)
}
