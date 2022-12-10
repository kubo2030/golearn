```markdown
# chromedp能做什么

- 反爬虫js，例如有的网页后台js自动发送心跳包，浏览器里会自动运行，不需要我们自动处理
- 针对于前端页面的自动化测试
- 解决类似VueJS和SPA之类的渲染
- 解决网页的懒加载
- 网页截图和pdf导出，而不需要额外的去学习其他的库实现
- seo训练和刷点击量
- 执行javascript 代码
- 设置dom的标签属性


# 使用前提

- 懂一点html和 css 以及js，因为操作 html 的 dom 元素需要用到 xpath 和 css 选择器之类的，如果 F12 的 element 里会右击复制 selector 也行，但是复杂的选择器还得需要 xpath 或者 css 选择器。

- chrome 打开网页 F12 后下面的调试工具出来后点击Elements,然后点击elements右边的那个框框里的鼠标箭头，点击后变蓝色，然后放到网页上选中区域点击一下，下面的内容就跳到对应地方，然后下面右击html的标签->Copy->COpy selector或者xpath，就能复制选择器了。