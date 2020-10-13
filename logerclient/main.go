package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	json := []string{`{"tablname":"logtable","colonsname":"logs","log":"AlarmTestMessage"}`, `{"tablname":"elogtable","colonsname":"elogs","log":"ErrorMassage"}`}
	for _, zn := range json {
		req, err := http.NewRequest("GET", "http://127.0.0.1:4444/", bytes.NewBuffer([]byte(zn)))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		//fmt.Println("response Status:", resp.Status)
	}
}
