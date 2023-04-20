package main

import (
	"encoding/json"
	"go-extensive-client-server-api/server/dao"
	models "go-extensive-client-server-api/server/models"
	"io"
	"net/http"
	"time"
)

func main() {
	//criando as rotas
	http.HandleFunc("/cotacao", findCambioHandler)

	// inicializando o servidor
	http.ListenAndServe(":8080", nil)
}

func findCambioHandler(w http.ResponseWriter, r *http.Request) {
	result, err := findCambio()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//adicionando o resultado no response
	json.NewEncoder(w).Encode(result)

}

func findCambio() (*models.CambioQuotation, error) {

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	c := http.Client{Timeout: time.Duration(200) * time.Millisecond}
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

	var usdBrl models.CambioQuotation
	err = json.Unmarshal(res, &usdBrl)
	if err != nil {
		print(err.Error())
	}

	dao.SaveRequest(usdBrl)
	return &usdBrl, nil

}
