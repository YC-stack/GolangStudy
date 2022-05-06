package main

import (
	"fmt"
	"encoding/json"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"Zhouxingchi", "Zhangbozhi"}}

	// 编码的过程，结构体--->json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return 
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码的过程，json--->结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json Unmarshal error", err)
		return 
	}

	fmt.Printf("%v\n", myMovie)

}