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

