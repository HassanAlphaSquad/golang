package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path

	if r.URL.Path == "/" {
		path = "./.html"
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		fmt.Println("Error: File not found -", path)
		return
	}

	content, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		fmt.Println("Error reading file:", err)
		return
	}

	w.Write(content)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
