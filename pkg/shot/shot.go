package shot

import (
	"C"
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

// Options defines configuration for taking a screenshot.
type Options struct {
	URL     string // URL to capture
	Full    bool   // If true, capture full page
	Width   int    // Viewport width
	Height  int    // Viewport height
	Quality int    // Quality for full screenshot (default: 90)
}

// NewTimeoutContext creates a new chromedp context with the given timeout.
func NewTimeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx := context.Background()
	return context.WithTimeout(ctx, timeout)
}

// Capture takes a screenshot based on the provided options.
func Capture(ctx context.Context, opts Options) ([]byte, error) {
	browserCtx, cancel := chromedp.NewContext(ctx)
	defer cancel()
	var actions []chromedp.Action
	actions = append(actions,
		chromedp.EmulateViewport(int64(opts.Width), int64(opts.Height)),
		chromedp.Navigate(opts.URL),
		chromedp.WaitVisible("body", chromedp.ByQuery),
	)

	var buf []byte
	if !opts.Full {
		actions = append(actions, chromedp.CaptureScreenshot(&buf))
	} else {
		quality := opts.Quality
		if quality == 0 {
			quality = 90
		}
		actions = append(actions, chromedp.FullScreenshot(&buf, quality))
	}

	if err := chromedp.Run(browserCtx, actions...); err != nil {
		return nil, err
	}
	return buf, nil
}
