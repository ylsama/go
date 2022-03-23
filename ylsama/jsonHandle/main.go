package main

import (
	"encoding/json"
	"fmt"
)

// User model
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social model
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func prettyPrint(arg1 User) {
	byteArray, err := json.MarshalIndent(arg1, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}
func main() {
	social := Social{Facebook: "https://facebook.com", Twitter: "https://twitter.com"}
	user := User{Name: "LanKa", Type: "Author", Age: 25, Social: social}

	JSONbyte, err := toString(user)

	if err != nil {
		panic(err)
	} else {
		toJSON(JSONbyte)
	}

}

func toString(user User) (byteArray []byte, err error) {
	fmt.Println(user)
	byteArray, err = json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
	prettyPrint(user)
	return
}
func toJSON(jsonData []byte) {
	myJSONMap := make(map[string]interface{})
	err := json.Unmarshal(jsonData, &myJSONMap)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", myJSONMap)
}
