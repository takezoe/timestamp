package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tkuchiki/go-timezone"
)

func main() {
	var (
		tz   = flag.String("tz", "UTC", "TimeZone (e.g. JST, UTC, PST)")
		unit = flag.String("u", "msec", "TimeUnit (e.g. sec, msec, nano)")
	)
	flag.Parse()
	args := flag.Args()

	offset, _ := timezone.GetOffset(*tz)
	loc := time.FixedZone(*tz, offset)

	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	value := args[0]

	i, e := strconv.ParseInt(value, 10, 64)
	if e == nil {
		// sec to msec
		if i < 10000000000 {
			i = i * 1000
		}
		// msec to nsec
		if i < 10000000000000 {
			i = i * 1000000
		}
		t := time.Unix(0, i).In(loc)
		fmt.Println(t.Format("2006-01-02 15:04:05") + " " + *tz)

	} else {
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", value, loc)
		i := t.UnixNano()
		if *unit == "sec" {
			i = i / 1000000000
		} else if *unit == "msec" {
			i = i / 1000000
		} else if *unit == "nano" {
			i = i
		} else {
			i = i / 1000000 // default is msec
		}
		fmt.Println(i)
	}
}
