```golang
package main

import (
    "context"
    "errors"
    "github.com/chromedp/cdproto/cdp"
    "github.com/chromedp/cdproto/network"
    "time"
    "log"
    "github.com/chromedp/chromedp"
)

var res string  // 定义全局变量，用来保存爬虫的数据

func main() {
    var err error
        
    // 创建链接
    ctxt, cancel := context.WithCancel(context.Background())
    defer cancel()
 
    //创建chrome.New()创建新的chrome实例
    c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf)) 
    if err != nil {
        log.Fatal(err)
    }
    x, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
    if err != nil {
        log.Fatal(err)
    }
    //执行任务
    err = c.Run(ctxt, visitWeb("http://dl.gaggjz.pw:8086/OpRoot/MemberScoreList.aspx?uid=0&op=0&uname=sdafsadsaf"))
    if err != nil {
        log.Fatal(err)
    }
    // 循环翻页
    for i := 1; i < 25000; i++ {
        //执行
        err = x.Run(ctxt, DoCrawler()) //执行爬虫任务
        WirteTXT(res)// res的内容写入文本 
    }

}

// 任务 主要用来设置cookie ，获取登录账号后的页面
func visitWeb(url string) chromedp.Tasks {

    return chromedp.Tasks{  
        chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
            expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
            success, err := network.SetCookie("ASP.NET_SessionId", "这里是值"). //设置cookie
                WithExpires(&expr).
                WithDomain("dl.gaggjz.pw:8086"). //访问网站主体
                WithHTTPOnly(true).
                Do(ctxt, h)
            if err != nil {
                return err
            }
            if !success {
                return errors.New("could not set cookie")
            }

            return nil
        }),
        chromedp.Navigate(url), //页面跳转
    }
}
// 任务 主要执行翻页功能和或者html
func DoCrawler() chromedp.Tasks {

    //sel =fmt.Sprintf(`javascript:__doPostBack('anpDataPager','%s')`,"2")

    return chromedp.Tasks{
        chromedp.Sleep(1*time.Second), // 等待
        chromedp.WaitVisible(`#form1`, chromedp.ByQuery),等待id=from1页面可见  ByQuery是使用DOM选择器查找
        chromedp.Sleep(1*time.Second), 
        chromedp.Click(`.pagination li:nth-last-child(4) a`, chromedp.ByQuery),//点击翻页
        chromedp.OuterHTML(`tbody`, &res, chromedp.ByQuery), //获取改 tbody标签的html
    }
}


func WirteTXT(txt  string ) {
    f, err := os.OpenFile("1.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
    if err != nil {
        fmt.Println("os Create error: ", err)
        return
    }
    defer f.Close()

    bw := bufio.NewWriter(f)
    bw.WriteString(txt+"\n")
    bw.Flush()
}

```

说明
第一个任务主要就是为了设置cookie即模拟登录

第二个任务是主要用来点击下一页按钮和获取指定html内容