package main

import ( 
	"os"
	"encoding/json"
)

type Account struct {
	Number int `json:"n"`
	Currency int `json:"ccr"`
}

func main() {
	account := Account{Number: 1, Currency: 1}
	res, err := json.Marshal(account)

	if err != nil {
		panic(err)
	}	

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	
	if err != nil {
		panic(err)
	}
	
	rawJson := []byte(`{"n": 2, "ccr": 200}`)
	var accountX Account

	err = json.Unmarshal(rawJson, &accountX)
	
	if err != nil {
		panic(err)
	}
	
	println(accountX.Currency)
}