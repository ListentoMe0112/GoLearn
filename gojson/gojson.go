package gojson

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"`
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func TestMonster() {
	monster := Monster{
		Name:     "Niumowang",
		Age:      500,
		Birthday: "2022-01-01",
		Sal:      8000.0,
		Skill:    "Niumoquan",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("Serilize Fail, err = ", err)
	}

	fmt.Println("After Serilize: ", string(data))
	monster = Monster{}
	err = json.Unmarshal([]byte(data), &monster)
	if err == nil {
		fmt.Println("Unmarshal: ", monster)
	} else {
		fmt.Println("err = ", err)
	}
}
