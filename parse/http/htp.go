package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Field struct {
	Name              string        `json:"name"`
	Description       string        `json:"description"`
	Args              []interface{} `json:"args"`
	Type              interface{}   `json:"type"`
	IsDeprecated      bool          `json:"isDeprecated"`
	DeprecationReason string        `json:"deprecationReason"`
}

type Tipo struct {
	Kind          string        `json:"kind"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Field         []Field       `json:"fields"`
	InputFields   []interface{} `json:"inputFields"`
	Interfaces    []interface{} `json:"interfaces"`
	EnumValues    []interface{} `json:"enumValues"`
	PossibleTypes interface{}   `json:"possibleTypes"`
}

var tipos []Tipo

func main() {
	//shpitems := make(map[string][]string)
	//kinds := make(map[string]string)
	objcts := make(map[string]string)
	jsonFile, err := os.Open("tmQ.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)
	jsonData, err := json.Marshal(result["types"])
	if err != nil {
		log.Fatal(err)
	}

	er := json.Unmarshal(jsonData, &tipos)
	if er != nil {
		log.Fatal(er)
	}
	var output string
	for index := range tipos {

		if tipos[index].Kind != "SCALAR" && tipos[index].Kind != "INPUT_OBJECT" && !strings.Contains(tipos[index].Name, "Filter") && !strings.HasPrefix(tipos[index].Name, "Statisti") {

			objcts[tipos[index].Name] = tipos[index].Name
			fmt.Println(tipos[index].Name)
			//fmt.Println(tipos[index].Name)
			output += tipos[index].Name + "\r\n"
		}
		//	kinds[tipos[index].Kind] = tipos[index].Kind
	}

	err = ioutil.WriteFile("output.txt", []byte(output), 0644)
	if err != nil {
		panic(err)
	}

	//for index := range kinds {
	//	fmt.Println(kinds[index])
	//}

}
