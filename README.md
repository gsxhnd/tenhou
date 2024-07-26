# Tenhou

天凤牌谱地址 <https://tenhou.net/mjlog.html>
最近七天的牌谱数据， 每小时有一个文件<https://tenhou.net/sc/raw/list.cgi>
归档的数据列表， 每天有一个文件 <https://tenhou.net/sc/raw/list.cgi?old>
<!-- <https://tenhou.net/sc/raw/dat/2023/scc20230101.html.gz> -->
如果需要下载每天的牌桌数据 <https://tenhou.net/sc/raw/dat/[YEAR]/scc[DATE].html.gz>
往年的牌桌数据 <https://tenhou.net/sc/raw/scraw2022.zip>
Log转换成对应json <https://tenhou.net/5/mjlog2json.cgi?[LOG_ID]>

在下载完压缩包后 只有scc开头的html文件是包含凤凰桌的数据

## log

牌桌数据中包含了全部的牌谱的ID为`log`

> 2009022100gm-00b9-0000-f1c46168

四凤东：标准四人麻将，不允许食断。
四凤东食：允许食断。
四凤东食赤：允许食断，引入赤宝牌。
四凤东食赤速：允许食断，引入赤宝牌，游戏节奏更快。

00a9 四鳳南喰赤－
00e9 四鳳南喰赤－

00a1 四鳳東喰赤速
00e1 四鳳東喰赤速

00b9 三鳳南喰赤－
00f9 三鳳南喰赤速

00b1 三鳳東喰赤－
00f1 三鳳東喰赤速

## 生成训练数据

1. 下载牌桌数据，只下到对应天数的， 每小时的数据等天凤归档成对应天的再下载
2. 手动解压缩筛选出 scc 的牌桌数据
3. 通过脚本下载转换成牌谱数据（json）
