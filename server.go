package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Server is listening on 8795 port...")
	http.HandleFunc("/time", HelloServer)
	http.ListenAndServe(":8795", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	jsonTime, err := json.Marshal(map[string]string{"time": time.Now().Format(time.RFC3339)})
	if err != nil {
		return
	}
	w.Write(jsonTime)
}
