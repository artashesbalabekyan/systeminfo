package main

import (
	"encoding/json"
	"fmt"
	"log"

	"system-info/pkg/macos/hwinfo"
)

func main() {
	mem, err := hwinfo.Get()
	if err != nil {
		log.Fatal(err)
	}
	js, _ := json.Marshal(mem)
	fmt.Println(string(js))
}
