package service

import (
	"encoding/json"
	"go-extensive-client-server-api/server/dao"
	exchange "go-extensive-client-server-api/server/models"
	"io"
	"net/http"
	"time"
)

func FindExchange() (*exchange.Quotation, error) {

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	c := http.Client{Timeout: time.Duration(2000) * time.Millisecond}
	resp, err := c.Get(url)

	if err != nil {
		print(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		print(err.Error())
		return nil, err
	}

	var usdBrl exchange.Quotation
	err = json.Unmarshal(res, &usdBrl)

	if err != nil {
		return nil, err
	}

	go dao.SaveRequest(usdBrl)

	return &usdBrl, nil

}
