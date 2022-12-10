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




```golang
chromedp.NewContext() //初始化chromedp的上下文，后续这个页面都使用这个上下文进行操作
chromedp.Run()        //运行一个chrome的一系列操作
chromedp.Navigate()   //将浏览器导航到某个页面
chromedp.WaitVisible() //等候某个元素可见，再继续执行。
chromedp.Click()       //模拟鼠标点击某个元素
chromedp.Value()       //获取某个元素的value值
chromedp.ActionFunc()   //再当前页面执行某些自定义函数
chromedp.Text()         //读取某个元素的text值
chromedp.Evaluate()     //执行某个js，相当于控制台输入js
network.SetExtraHTTPHeaders()  //截取请求，额外增加header头
chromedp.SendKeys()            //模拟键盘操作，输入字符
chromedp.Nodes()               //根据xpath获取某些元素，并存储进入数组
chromedp.NewRemoteAllocatorchromedp.OuterHTML()  //获取元素的outer html
chromedp.Screenshot()                            //根据某个元素截图
page.CaptureScreenshot()                         //截取整个页面的元素
chromedp.Submit()                                //提交某个表单
chromedp.WaitNotPresent()                        //等候某个元素不存在，比如“正在搜索。。。”
chromedp.Tasks{}                                 //一系列Action组成的任务



```