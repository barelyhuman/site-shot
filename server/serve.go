package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/barelyhuman/site-shot/pkg/shot"
)

func captureHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	urlStr := query.Get("url")
	if urlStr == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	// Parse optional parameters.
	full := false
	if fullStr := query.Get("full"); fullStr != "" {
		if b, err := strconv.ParseBool(fullStr); err == nil {
			full = b
		}
	}

	height := 0
	if h := query.Get("height"); h != "" {
		if hVal, err := strconv.Atoi(h); err == nil {
			height = hVal
		}
	}

	width := 0
	if w := query.Get("width"); w != "" {
		if wVal, err := strconv.Atoi(w); err == nil {
			width = wVal
		}
	}

	// You can also parse quality from the query if needed.
	quality := 90

	// Create a context with a timeout.
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	image, err := shot.Capture(ctx, shot.Options{
		URL:     urlStr,
		Height:  height,
		Width:   width,
		Full:    full,
		Quality: quality,
	})
	if err != nil {
		http.Error(w, "failed to capture screenshot: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}

func main() {
	http.HandleFunc("/screenshot", captureHandler)

	port := ":8080"
	log.Printf("Server is listening on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
