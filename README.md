你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
timestamp
========
A utility for converting timestamp and unixtime.

```bash
# Convert datetime string to unixtime
$ timestamp -tz JST -sec "2019-11-26 23:36:00"
1574778960

# Convert unixtime to datetime string
$ timestamp -tz JST 1574778960
2019-11-26 23:36:00 JST

# Show current timestamp
$ timestamp
1586706426398
$ timestamp -tz JST
2020-04-13 00:47:12 JST
```

## Installation

```
go get github.com/takezoe/timestamp
```

## Usage

```
timestamp: a utility for converting timestamp and unixtime.

USAGE::
  timestamp [OPTIONS]            ... Show current timestamp
  timestamp [OPTIONS] [DATETIME] ... Convert date time string to unixtime
  timestamp [OPTIONS] [UNIXTIME] ... Convert unixtime to date time string

OPTIONS:
  -msec
        Output TimeUnit (msec) (default true)
  -nano
        Output TimeUnit (nano)
  -out string
        Output TimeZone (e.g. JST, UTC, PST)
  -sec
        Output TimeUnit (sec)
  -tz string
        TimeZone (e.g. JST, UTC, PST) (default "UTC")

ARGS:
  <DATETIME>
      yyyy/MM/dd [HH:mm[:ss] [timezone]]
      yyyy-MM-dd [HH:mm[:ss] [timezone]]
      dd MMM yyyy [HH:mm[:ss] [timezone]]
  <UNIXTIME>
      elapsed time (sec, msec, nsec) from 1970-01-01 00:00:00 UTC
```

