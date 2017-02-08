package main

import (
	"encoding/json"
	"fmt"
)

type Merchant struct {
	Id         uint
	PayMethods map[string][]PayMethod
	Accounts   []Account
}

type PayMethod string

type AccYa struct {
	ShopId string `json:"shop_id"`
	Scid   string `json:"scid"`
}

type Account struct {
	Id      string `json:"id"`
	Adapter string `json:"adapter"`
	Data    map[string]string `json:"data"`
}

func main() {
	m := Merchant{
		Id: 1,
		PayMethods: map[string][]PayMethod{
			"dol": []PayMethod{
				"ac",
				"ym",
			},
			"pscb": []PayMethod{
				"ac",
				"ym",
			},
		},
		Accounts: []Account{
			Account{
				Id:      "yandex-no-hold",
				Adapter: "yandex",
				Data: map[string]string{
					"ShopId": "123123",
					"scid":   "5555",
				},
			},
			Account{
				Id:      "yandex-hold",
				Adapter: "yandex",
				Data: map[string]string{
					"ShopId": "88888",
					"scid":   "9999",
				},
			},
		},
	}

	payMethodsJson, _ := json.Marshal(m.PayMethods)
	accountJson, _ := json.Marshal(m.Accounts)

	fmt.Println(string(payMethodsJson))
	fmt.Println(string(accountJson))

	a := []Account{}

	json.Unmarshal(accountJson, &a)

	for _, account := range a {
		fmt.Println(account.Adapter)
		//adapter := adapters[account.Adapter]
		//adapter.setAccount(account.Data)
		//url := adapter.BuildUrl()
		//getAcc(account.Data)
	}

	fmt.Println(a)



}

func getAcc(data map[string]string)  {

	/*structValue.FieldByName("ShopId")
	structValue := reflect.ValueOf(AccYa{}).Elem()
	fmt.Println(structValue)*/
}