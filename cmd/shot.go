package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/barelyhuman/site-shot/pkg/shot"
)

func main() {
	fullFlag := flag.Bool("full", false, "URL of the page to capture")
	urlFlag := flag.String("url", "", "URL of the page to capture")
	filenameFlag := flag.String("filename", "screenshot.png", "Output filename")
	heightFlag := flag.Int("height", 1920, "Viewport height")
	widthFlag := flag.Int("width", 1080, "Viewport width")
	flag.Parse()

	ctx, cancel := shot.NewTimeoutContext(30 * time.Second)
	defer cancel()
	image, err := shot.Capture(ctx, shot.Options{
		URL:     *urlFlag,
		Height:  *heightFlag,
		Width:   *widthFlag,
		Full:    *fullFlag,
		Quality: 90,
	})
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(
		*filenameFlag, image, os.ModePerm,
	)
}
