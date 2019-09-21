package main;

import ("net/http"; "fmt"; "time"; "encoding/json");

func handleRoot(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html");
	res.Write([]byte("Hello World!"));
}

func handleTime(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json");
	jsonTime, err := json.Marshal(map[string]string{ "time": time.Now().Format(time.RFC3339) });
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError);
		return;
	}
	res.Write(jsonTime);
}

func main() {
	http.HandleFunc("/", handleRoot);
	http.HandleFunc("/time", handleTime);
	defer http.ListenAndServe(":8795", nil);
	fmt.Println("Server is listening on 8795 port...");
}