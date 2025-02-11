package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// func handle_php(w http.ResponseWriter, r *http.Request) {
// 	cmd := exec.Command("php", "."+r.URL.Path) // Execute PHP script
// 	output, err := cmd.Output()
// 	if err != nil {
// 		http.Error(w, "Error executing PHP script", http.StatusInternalServerError)
// 		fmt.Println("PHP Execution Error:", err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "text/html")
// 	w.Write(output)
// }

func handle_php(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("php", "."+r.URL.Path) // Execute PHP script -> php filename.php
	output, err := cmd.CombinedOutput()        // Captures both stdout & stderr
	if err != nil {
		http.Error(w, string(output), http.StatusInternalServerError) // Show actual PHP error
		fmt.Println("PHP Execution Error:", err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(output) // Return PHP output (including errors if any)
}

// lists all files in a directory, excluding specified extensions
func list_files(directory string, excludeExts []string) ([]string, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() { // only add files, not directories
			ext := filepath.Ext(entry.Name())  // Get file extension (Name)
			ext = strings.TrimPrefix(ext, ".") // Remove the dot for comparison

			// Skip files with excluded extensions
			skip := false
			for _, ex := range excludeExts {
				if ext == ex {
					skip = true
					break
				}
			}
			if !skip {
				files = append(files, entry.Name())
			}
		}
	}
	return files, nil
}

// list file links into homepage
func gen_file_links(excludeExts []string) (string, error) {
	content, err := os.ReadFile(".html")
	if err != nil {
		return "", err
	}

	files, err := list_files("./", excludeExts) // filters out excluded extension files
	if err != nil {
		return "", err
	}

	html := string(content)

	var listItems []string
	for _, file := range files {
		if file != ".html" && file != "404.html" { // avoid self-referencing
			listItems = append(listItems, fmt.Sprintf(`<li><a href="%s">%s</a></li>`, file, file))
		}
	}

	// Replace <ol class="links"> dynamically
	re := regexp.MustCompile(`<ol class="links">.*?</ol>`)
	newList := fmt.Sprintf(`<ol class="links">%s</ol>`, strings.Join(listItems, "\n"))
	html = re.ReplaceAllString(html, newList)

	return html, nil
}

func handle_req(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		excluded := []string{"go"} // excluded extensions
		html, err := gen_file_links(excluded)
		if err != nil {
			http.Error(w, "Error generating HTML", http.StatusInternalServerError)
			fmt.Println("Error generating HTML:", err)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
		return
	}

	// Handle PHP files separately
	if filepath.Ext(r.URL.Path) == ".php" {
		handle_php(w, r)
		return
	}

	path := "." + r.URL.Path

	if _, err := os.Stat(path); os.IsNotExist(err) {
		path = "./responses/404.html" // Set path to the 404 page

		content, err := os.ReadFile(path)
		if err != nil {
			http.Error(w, "404 Not Found", http.StatusNotFound) // Fallback if 404.html is missing
			fmt.Println("Error: 404 page not found -", err)
			return
		}

		w.WriteHeader(http.StatusNotFound) // Set HTTP 404 status
		w.Write(content)                   // Serve 404 page
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.HandleFunc("/", handle_req)

	fmt.Printf("Server is running on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
