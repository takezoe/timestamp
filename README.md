timestamp
========
A utility for converting timestamp and unixtime.

```
$ timestamp -tz JST -sec "2019-11-26 23:36:00"
1574778960
$ timestamp -tz JST 1574778960
2019-11-26 23:36:00 JST
```

## Installation

```
go get github.com/takezoe/timestamp
```

## Usage

```
timestamp: a utility for converting timestamp and unixtime.

USAGE::
  timestamp [OPTIONS] [TIMESTAMP or UNIXTIME]

OPTIONS:
  -msec
    	Output TimeUnit (msec) (default true)
  -nano
    	Output TimeUnit (nano)
  -sec
    	Output TimeUnit (sec)
  -tz string
    	TimeZone (e.g. JST, UTC, PST) (default "UTC")

ARGS:
  <TIMESTAMP>
      yyyy-MM-dd HH:mm:ss
  <UNIXTIME>
      elapsed time (sec, msec, nsec) from 1970-01-01 00:00:00
```

