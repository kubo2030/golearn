```
配置UA
```

```golang
package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {

	var ua string

	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)

	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()
	// create context
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.whatsmyua.info/?a`),
		chromedp.WaitVisible(`#custom-ua-string`),
		chromedp.Text(`#custom-ua-string`, &ua),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user agent: %s", ua)
}

```