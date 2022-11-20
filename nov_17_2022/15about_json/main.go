package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("welcome to the json wideo")
	// EncodeJson()
	DecodeJson()
}

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`
	Status   bool
	Tag      []string `json:"Tags,omitempty"`
}

func EncodeJson() {

	lcoCourse := []course{

		{"ReactJs", 245, "LernCodeOnline.in", "abcd123", true, []string{"web-dev", "js"}},
		{"MERN Bootcamp", 678, "LernCodeOnline.in", "abdg123", false, []string{"web-dev", "js"}},
		{"Angular BootCamp", 256, "LernCodeOnline.in", "pbce123", true, []string{"web-dev", "js"}},
		{"Go BootCamp", 896, "LernCodeOnline.in", "ghcd043", true, nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	decodeDataFromJson := []byte(`
	{
                "coursename": "MERN Bootcamp",
                "Price": 678,
                "website": "LernCodeOnline.in",
                "Status": false,
                "Tags": ["web-dev","js"]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(decodeDataFromJson)

	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(decodeDataFromJson, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}
	fmt.Println("Json was not valid")
}
