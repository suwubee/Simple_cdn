golang实现简易单点cdn
==========

项目起因是upyun流量太大,无法承受,[魅族迷](http://www.meizumi.com)一个页面近200个图片

线上环境用于[魅族迷](http://www.meizumi.com)

只做到了简单的 pull资源(本地无资源则回源)

因为代码原因,修改图片不需要(无实现)

目前还没有删除资源功能,只是抛砖引玉,需要的可以自己写(线上环境已经实现单点通信删除)

压测没有问题!

linux 64位编译下载 [Simple_cdn](https://raw.github.com/sun8911879/Simple_cdn/master/Simple_cdn_x86_64.tar.gz)

注意
==========
[魅族迷](http://www.meizumi.com)之前压测出现大面积404

原因是句柄被跑满(感谢 [ASTA谢](http://weibo.com/533452688)帮找出问题！)

使用者也需要注意的问题！

open files 默认是1024(线上已经修改为20W)

不适用于windows平台,由于系统内核原因 跑30分钟左右自动挂掉(windows就表挣扎了)

关于
==========
本程序由golang开发

如果你对golang感兴趣 

请看go开源图书 [Go Web 编程](https://github.com/astaxie/build-web-application-with-golang)

你也可以通过以下途径购买书籍:

- [chinapub](http://product.china-pub.com/3767290)
- [当当网](http://product.dangdang.com/product.aspx?product_id=23231404)
- [京东](http://book.jd.com/11224644.html)
- [Amazon](http://www.amazon.cn/Go-Web%E7%BC%96%E7%A8%8B-%E8%B0%A2%E5%AD%9F%E5%86%9B/dp/B00CHWVAHQ/ref=sr_1_1?s=books&ie=UTF8&qid=1369323453&sr=1-1)

本人不接受各种定制 请勿扰
- by [雪虎](http://weibo.com/sun8911879)
