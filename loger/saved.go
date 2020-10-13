package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func formattime(t time.Time) string {
	return t.Format("2006.01.02-15.04.05")
}

func saved(data string) {
	filename := fmt.Sprintf("savedlogs/%s", formattime(time.Now()))
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(data)
}

func realsize() int64 {
	var size int64
	files, err := ioutil.ReadDir("savedlogs")
	if err != nil {
		fmt.Println(err)
	}
	for _, zn := range files {
		size += zn.Size()
	}
	return size
}
