package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	structs "github.com/liukaku/jsonParse/utils"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		return
	}

	apiKey := os.Getenv("API_KEY")

	var inputName string
	var numberOfStops string

	fmt.Println("Enter file name, leave it blank for 'stops'")
	fmt.Scanln(&inputName)
	if inputName == "" {
		inputName = "stops"
	}
	fmt.Println("How many stops? (defaults to all)")
	fmt.Scanln(&numberOfStops)
	if numberOfStops == "" {
		numberOfStops = "1000"
	}

	getUrl := fmt.Sprintf("https://api.tfgm.com/odata/Metrolinks?$top=%s", numberOfStops)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", getUrl, nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)
	res, err := client.Do(req)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fileOpen := res.Body

	readFile, err := io.ReadAll(fileOpen)

	var fileStruct structs.Result
	
	json.Unmarshal(readFile, &fileStruct)


	backToJson, err := json.Marshal(fileStruct)

	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := fmt.Sprintf("%s.json", inputName)

	os.WriteFile(fileName, backToJson, 0644)

}
