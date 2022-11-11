package client

import (
	"encoding/json"
	"log"
	"net/http"
)

func FetchCrypto(fiat string, crypto string) (map[string]any , error) {

	URL := "https://rest.coinapi.io/v1/exchangerate/"+crypto+"/"+fiat+"?apikey=5DD84F59-0100-470F-BE2B-86CC6BB5A12A"

	resp, err := http.Get(URL)

	if err != nil {
		log.Fatal(err,1)
	}

	
	defer resp.Body.Close()

	var myResponse map[string]any
    	
	if err:= json.NewDecoder(resp.Body).Decode(&myResponse); err!=nil {
		log.Fatal(err,2)
	}
	
	return myResponse, err
}