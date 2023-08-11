package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// Set CORS headers to allow requests from any origin
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// 	fmt.Println("helllo")
	// 	htmlContent := `
	// 		document.getElementById('demo').innerHTML = "<img src='https://fujifilm-x.com/wp-content/uploads/2021/01/gfx100s_sample_01_thum.jpg' />"
	// 	`
	// 	w.Header().Set("Content-Type", "text/html")
	// 	fmt.Fprint(w, htmlContent)
	// })

	// fmt.Println("Server started on :8080")
	// http.ListenAndServe(":8080", nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		// 	// Set CORS headers to allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		http.ServeFile(w, r, "app.js")
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
