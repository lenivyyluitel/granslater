package granslater

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Translate takes query and its source lanuage
// and converts to the targeted language
// Translate("Hello", "en", "jp")
func Translate(query, source, target string, debugFlag bool) interface{} {
	queryMap := map[string]string{
		"q":      query,
		"source": source,
		"target": target,
		"format": "text",
	}

	jsonData, err := json.Marshal(queryMap)

	resp, err := http.Post("https://libretranslate.de/translate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Panic(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	if debugFlag {
		fmt.Println(res)
	}
	return res["translatedText"]
}
