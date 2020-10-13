package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// JSONData form inputdata
type JSONData struct {
	Tablname   string `json:"tablname"`
	Colonsname string `json:"colonsname"`
	Log        string `json:"log"`
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var buferLog []JSONData
	rezjson := JSONData{}
	databody := ""
	io.WriteString(w, "Hello, Im is loger servis!\n")

	for true {
		bs := make([]byte, 1024)
		n, err := r.Body.Read(bs)
		databody = (databody + string(bs[:n]))
		if n == 0 || err != nil {
			break
		}
	}
	defer r.Body.Close()
	if err := json.Unmarshal([]byte(databody), &rezjson); err != nil {
		log.Fatal(err)
	}

	buferLog = append(buferLog, rezjson)
	//запись на диск
	if realsize() < sizelimit {
		saved(databody)
	}

	for len(buferLog) > 0 {
		err := connectsql(buferLog[0].Tablname, buferLog[0].Colonsname, buferLog[0].Log)
		w.WriteHeader(http.StatusAccepted)
		if err != nil {
			continue
		}
		if len(buferLog) > 1 {
			buferLog = buferLog[1:]

		} else {
			buferLog = buferLog[:0]
		}
	}
}

func main() {
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":4444", nil))
}
