```golang
package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {

	var ua string

	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),  // 开启 chrome GUI ,可以看到chrome具体干了什么操作
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()
	// create context
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.whatsmyua.info/?a`),
		chromedp.WaitVisible(`#custom-ua-string`),
		chromedp.Text(`#custom-ua-string`, &ua),
		chromedp.Sleep(10* time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user agent: %s", ua)
}
```