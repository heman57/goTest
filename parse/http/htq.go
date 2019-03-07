package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type FieldType struct {
	Kind   string      `json:"kind"`
	Name   string      `json:"name"`
	OfType interface{} `json:"ofType"`
}

type Field struct {
	Name              string        `json:"name"`
	Description       string        `json:"description"`
	Args              []interface{} `json:"args"`
	Type              FieldType     `json:"type"`
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
	//rootF := make(map[string]string)
	objcts := make(map[string][]Field)
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
		//&& tipos[index].Kind == "INPUT_OBJECT"
		if tipos[index].Kind != "SCALAR" && !strings.Contains(tipos[index].Name, "Filter") && !strings.HasPrefix(tipos[index].Name, "Statisti") {

			objcts[tipos[index].Name] = tipos[index].Field
			//fmt.Println(tipos[index].Field)
			//fmt.Println(tipos[index].Name)
			//output += tipos[index].Name + "\r\n"
		}
		//	kinds[tipos[index].Kind] = tipos[index].Kind
	}

	for index := range tipos[0].Field {
		fmt.Println(tipos[0].Field[index].Name)
		output += tipos[0].Field[index].Name + "\r\n"
		for index2 := range objcts[tipos[0].Field[index].Type.Name] {
			//fmt.Println(index)

			fmt.Println(tipos[0].Field[index].Name + " - " + objcts[tipos[0].Field[index].Type.Name][index2].Name)
			output += tipos[0].Field[index].Name + " - " + objcts[tipos[0].Field[index].Type.Name][index2].Name + "\r\n"
		}
		fmt.Println("-------------------------")
		//rootF[tipos[0].Field[index].Name] = tipos[0].Field[index].Name
	}

	err = ioutil.WriteFile("output.txt", []byte(output), 0644)
	if err != nil {
		panic(err)
	}

	//	fmt.Println(objcts["ListMutation"])

	//	fmt.Println(objcts["ApiMutation"][0].Name)

}
