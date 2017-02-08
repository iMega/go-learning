package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type AccYa struct {
	ShopId string `json:"shop_id"`
	Scid   string `json:"scid"`
}

type Account struct {
	Id      string `json:"id"`
	Adapter string `json:"adapter"`
	Data    interface{} `json:"data"`
}

func main() {
	//jsStr := []byte("[{\"id\":\"yandex-no-hold\",\"adapter\":\"yandex\",\"data\":{\"ShopId\":\"shopfromjson\",\"scid\":\"5555\"}},{\"id\":\"yandex-hold\",\"adapter\":\"yandex\",\"data\":{\"ShopId\":\"88888\",\"scid\":\"9999\"}}]")
	jsStr := []byte("[{\"id\":\"yandex-no-hold\",\"adapter\":\"yandex\",\"data\":{\"shop_id\":\"shopfromjson\",\"scid\":\"5555\"}}]")
	a := []Account{}

	json.Unmarshal(jsStr, &a)

	for _, v := range a {
		fmt.Println(v.Data)
		getAccount(v.Data)
	}
}

func getAccount(data interface{}) {
	a := AccYa{}

	value := reflect.ValueOf(&a).Elem()
	typ := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		valueField := getValue(typ.Field(i).Tag.Get("json"), data)
		field.SetString(valueField.Interface().(string))
	}

	fmt.Println(a)
}

func getValue(keyName string, data interface{}) reflect.Value {
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Map {
		for _, key := range value.MapKeys() {
			if keyName == key.String() {
				return value.MapIndex(key)
			}
		}
	}
	return reflect.Value{}
}


