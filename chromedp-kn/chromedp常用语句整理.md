基础代码
```golang
package main

import (
    "context"
    "log"
    "time"

    "github.com/chromedp/chromedp"
)

func main() {
    log.Printf("自动化助手:")
    dowork()
}

func dowork() {
    //增加选项，允许chrome窗口显示出来
    options := []chromedp.ExecAllocatorOption{
        chromedp.Flag("headless", false),
        chromedp.Flag("hide-scrollbars", false),
        chromedp.Flag("mute-audio", false),
        chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
    }
    options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
    //创建chrome窗口
    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
    defer cancel()
    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()
    
     //可以使用多个chromedp.Run()
    if err := chromedp.Run(ctx,
        chromedp.Navigate(`http://192.168.132.80/login/Login.jsp?logintype=1`),
        chromedp.WaitVisible(`#loginid`, chromedp.ByID),
        chromedp.SendKeys(`input[name=loginid]`, "admin"),
        chromedp.WaitVisible(`#loginid`, chromedp.ByID),
        chromedp.SendKeys(`input[name=userpassword]`, "1234"),
        chromedp.Click(`#login`, chromedp.ByID),
        //在这里加上你需要的后续操作，如Navigate，SendKeys，Click等
        chromedp.Sleep(10*time.Second),
    ); err != nil {
        panic(err)
    }

}
```