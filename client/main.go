package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-extensive-client-server-api/server/models"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)

	//cancelando o contexto
	defer cancel()
	url := "http://localhost:8080/quotation"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var cambioQuotation models.Quotation
	err = json.Unmarshal(body, &cambioQuotation)
	if err != nil {
		panic(err)
	}
	fmt.Print(cambioQuotation)
	writeFile(cambioQuotation.USDBRL.Bid)

	defer res.Body.Close()

}

func writeFile(bid string) {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	length, err := f.Write([]byte("Valor dólar: " + bid))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso, possui %v bytes", length)
	f.Close()
}
