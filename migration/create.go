package migration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Initialization() {

	var path string = "migration"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}

	fmt.Println("okey")
}

func CallAPI() {

	type Linkedaccounts struct {
		Id   string `json:"id"`
		Icon string `json:"icon"`
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	type Resbody struct {
		Code    string `json:"code"`
		Message struct {
			LinkedAccount []Linkedaccounts `json:"linked_accounts"`
		} `json:"message"`
	}

	resp, err := http.Get("https://private-0205a-partnerwallet.apiary-mock.com/v3/partner-wallet/ex/linked")
	if err != nil {
		fmt.Printf("error : ")
	} else {

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error : ", err)
		} else {
			// bd := string(body)
			// fmt.Println(bd)
			_ = json.Unmarshal([]byte(body), &Resbody)
			err != nil{}
		}
	}
}
