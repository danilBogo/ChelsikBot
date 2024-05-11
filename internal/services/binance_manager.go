package services

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const binancePriceUrl = "https://fapi.binance.com/fapi/v1/ticker/price?symbol=TONUSDT"

type Data struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type BinanceManager struct{}

func NewBinanceManager() *BinanceManager {
	return &BinanceManager{}
}

func (bm *BinanceManager) GetTon() Data {
	response, err := http.Get(binancePriceUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bytesBuffer := bytes.Buffer{}
	buffer := make([]byte, 1024)
	for {
		bytesRead, err := response.Body.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytesBuffer.Write(buffer[:bytesRead])

		if err == io.EOF {
			break
		}
	}

	var data Data
	err = json.Unmarshal(bytesBuffer.Bytes(), &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
