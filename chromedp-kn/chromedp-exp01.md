```golang
chromedp.Run(ctx,
        //chromedp.Emulate(device.IPhone7),
        chromedp.Navigate(`http://192.168.132.80/login/Login.jsp?logintype=1`),
        chromedp.WaitVisible(`#loginid`, chromedp.ByID),
        chromedp.Sleep(1*time.Second),
        chromedp.SendKeys(`input[name=loginid]`, "admin"),
        chromedp.WaitVisible(`#loginid`, chromedp.ByID),
        chromedp.SendKeys(`input[name=userpassword]`, "1234"),
        chromedp.Click(`#login`, chromedp.ByID),
        chromedp.WaitVisible(`#_ButtonCancel_0`, chromedp.ByID),
        chromedp.Click(`#_ButtonCancel_0`, chromedp.ByID),
        chromedp.Stop(),
        chromedp.Navigate(`http://192.168.132.80/docs/docs/DocAddForCK.jsp?mainid=15&subid=49&secid=1143&showsubmit=1&coworkid=&prjid=&isExpDiscussion=&crmid=&hrmid=&topage=`),
        chromedp.WaitVisible(`#docsubject`, chromedp.ByID),
        chromedp.Sleep(1*time.Second),
        chromedp.SendKeys(`input[name=docsubject]`, "aa11"),
        //禁止alert弹窗。 防止错误提醒；参考我上篇文章，其实不需window.alert = function(){return false;};这种暴力方法!
        chromedp.EvaluateAsDevTools(`window.alert = function(){return false;};var doc =document.querySelector("#cke_contents_doccontent > iframe").contentWindow.document;
p = doc.createElement("p");
p.innerText="abc";
doc.body.append(p);`, &buf),
        chromedp.Sleep(1*time.Second),
        chromedp.Click("#BUTTONnull", chromedp.ByID),
        chromedp.Sleep(1*time.Second),
        chromedp.Click(`document.querySelector("#BUTTONnull")`, chromedp.ByJSPath),
        chromedp.ActionFunc(func(ctx context.Context) error {
            ioutil.WriteFile("1.txt", buf, 0777)
            return nil
        }),
        chromedp.Sleep(2*time.Second),
        //chromedp.CaptureScreenshot(&buf),
    )
```