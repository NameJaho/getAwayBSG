# 逃离北上广

[这里是python-scrapy-mysql实现](https://github.com/jiangwei1995910/getAwayBSG/tree/python)

> **注意！**\
> 1.本项目仅供学习研究，禁止用于任何商业项目\
> 2.运行的时候为被爬方考虑下！尽量不要爬全站。请在配置文件中设置你需要的城市爬取即可！



## What!

如果你是一个正准备逃离北上广等一线城市却又找不到去处的IT人士，或许这个项目能给你点建议。

## Desc

或许你跟我一样困惑，为此我通过爬虫抓取了智联招聘跟链家这2个平台的全部数据。最终拿到了18W+全国各个城市的招聘数据与81W+全国各地的房屋成交记录数据。

其中，招聘数据我抓取了工作年限，公司名称，公司规模，公司类型，工作类型，创建时间，工作名称，结束时间，教育情况，薪资字段。

职位我使用了['php', 'java', 'python', 'c/c++', 'c#', 'mysql', 'oracle', 'javascript', 'linux', 'SQL', '软件', '程序员']作为关键词搜索。基本上涵盖了程序猿们绝大部分工作

对应房屋成交记录，我抓取了成交时间，成交价格，每平米均价，地址字段。

## 分析

### 综合分析

首先，房价和薪资都没法代表一个地方的生活成本情况。因此我使用了（月薪/每平米房价）的倒数值来表示一个城市的生活成本。这个值越大，表面这个地方的生活成本越高。结果如下图
![](./docs/img/shcb.png)
结果很明显了，如果你想去一个安稳生活的地方，这个表中前几的城市都不错，买房压力较低，并且我相信经常逛Github的程序猿肯定都是平均薪资几倍的收入。反之，如果你想挑战人生的地狱模式，emmmmm


另外，也附上各个城市的月薪，每平米房价对比情况
![](./docs/img/fjxz.png)

### 工作机会

统计了各个规模公司的招聘数量
![](./docs/img/gzjh.png)



### 薪资

首先，我计算了各个城市的平均薪资情况，为什么不包含博士学历呢？因为智联上面写明要博士的职位很少(硕士其实也不多，只有几百的量)，抓下来每个城市都是几十的量，这种数据不具有统计意义如下：

![](./docs/img/avg.png)

可见，硕士的薪资跟本科比并没多大差距。另外，别小看了拉萨，拉萨的薪资并不低，但是方差特别大，说明如果你愿意去拉萨，其实你的薪资会很高，平均薪资低是因为有很多1000 2000的工作拉下去的

其他都是意料之中的了


其实，平均值并没有多大参考意义，因为被平均的东西太多了，比如工龄就是很重要一个，于是我又取了各工龄的情况，如下

![](./docs/img/workTime.png)

### 房价

对于一个地方是否合适自己发展，房价非常重要，于是我也分析了各个城市的房价数据。

首先，也先来个平均
![](./docs/img/avgRoom.png)


其次，也来个近10年房价的走势图吧
![](./docs/img/room.png)
似乎有个大致规律，上半年涨，下半年跌。


## How to run

**这应该是你在Github上能找到的运行最简单的爬虫项目！**  如果你想要运行这个项目，在[releases](https://github.com/jiangwei1995910/getAwayBSG/releases)里面下载你需要的操作系统平台，修改配置文件，双击运行，搞定！


## 未完成 TODO

1.Go语言重构后目前只写了链家和智联招聘的爬虫，自如的还没写

2.有空再加一些其他网站的数据进来

3.加入租房数据，计算各个城市的租售比




## 2019-07-03 Update Log

为了能够在树莓派(没错，我就是拿树莓派跑爬虫服务的)上面有更好的运行效率，使用了Golang语言重构了整个项目，数据库换为MongoDB

爬取数据源修改：链家改为爬取二手房数据，而不是以前的二手房交易记录





