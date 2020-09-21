package wantMore

import "net/http"

func getHeaders() [][]string {
	return [][]string{
		{"Content-Type", "text/plain"},
	}
}

func addHeaders(w http.ResponseWriter) {
	for _, h := range getHeaders() {
		w.Header().Add(h[0], h[1])
	}
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	//addHeaders(w)
	w.Write([]byte("Hello world"))
}

func WantMore_Solution() {
	http.HandleFunc("/", serverHandler)
	http.ListenAndServe(":80", nil)
}
