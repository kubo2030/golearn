常用功能：

```golang

1、给input设置值（还可以SendKeys）

chromedp.SetValue(`#loginid`, `aa`, chromedp.ByID),
chromedp.SendKeys(`input[name=userpassword]`, "123"),

2、选择元素，除chromedp.ByID，还可用 chromedp.ByJSPath

chromedp.SetValue(`document.querySelector("#loginid")`, `bb`, chromedp.ByJSPath),

3、设置值：

chromedp.SetValue(`#loginid`, `cc`, chromedp.ByQuery),

4、延时几秒：

chromedp.Sleep(10*time.Second),

5、输出OuterHTML(难点在iframe的选择）

chromedp.OuterHTML(`document.querySelectorAll("iframe")[3]`, &text1, chromedp.ByJSPath),

6、在页面上执行javascript

chromedp.EvaluateAsDevTools(`alert("test eval");`, &text1),

7、运行自定义函数

chromedp.ActionFunc(func(ctx context.Context) error {
ioutil.WriteFile("1.txt", []byte(text1), 0777)
return nil
}),

8、获取iframe内容，页面有个id=#cke的td,其中有个iframe，用：

document.querySelector("#cke_contents_doccontent > iframe").contentWindow

上面的语句是在chrome console中测试出来的，
在console中$和document.getElementById返回值类型不一样，一个是数组，可以在console中看出来。
用类似以下语句，获取和设置iframe中的内容:
document.querySelector("#cke_contents_doccontent > iframe").contentWindow.document.querySelector('p').innerText="aaaa"

9、停止网页加载（不停止的话，有时会长时间加载）

chromedp.Stop(),

10、等元素出现时

chromedp.WaitVisible(`#docsubject`, chromedp.ByID),

11、等元素消失时

chromedp.WaitNotVisible(`#docsubject`, chromedp.ByID),

```