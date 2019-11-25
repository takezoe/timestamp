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
	flag.Usage = func() {
		self := os.Args[0]
		usage(self)
	}
	var (
		// TODO: Allow to include TimeZone in timestamp?
		tz   = flag.String("tz", "UTC", "TimeZone (e.g. JST, UTC, PST)")
		sec  = flag.Bool("sec", false, "Output TimeUnit (sec)")
		msec = flag.Bool("msec", true, "Output TimeUnit (msec)")
		nano = flag.Bool("nano", false, "Output TimeUnit (nano)")
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
		time2str(i, tz, loc)
	} else {
		str2time(value, sec, msec, nano, loc)
	}
}

func time2str(value int64, tz *string, loc *time.Location) {
	// sec to msec
	if value < 10000000000 {
		value = value * 1000
	}
	// msec to nsec
	if value < 10000000000000 {
		value = value * 1000000
	}
	t := time.Unix(0, value).In(loc)
	fmt.Println(t.Format("2006-01-02 15:04:05") + " " + *tz)
}

func str2time(value string, sec *bool, msec *bool, nano *bool, loc *time.Location) {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", value, loc)
	i := t.UnixNano()
	if *sec == true {
		i = i / 1000000000
	} else if *msec == true {
		i = i / 1000000
	} else if *nano == true {
		i = i
	} else {
		i = i / 1000000 // default is msec
	}
	fmt.Println(i)
}

func usage(self string) {
	fmt.Fprintf(os.Stderr, "%s: a utility for convertig timestamp and unixtime.\n\n", self)
	fmt.Fprintf(os.Stderr, "USAGE::\n")
	fmt.Fprintf(os.Stderr, "  %s [OPTIONS] [TIMESTAMP or UNIXTIME]\n\n", self)
	fmt.Fprintf(os.Stderr, "OPTIONS:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "ARGS:\n")
	fmt.Fprintf(os.Stderr, "  <TIMESTAMP>\n")
	fmt.Fprintf(os.Stderr, "      yyyy-MM-dd HH:mm:ss\n")
	fmt.Fprintf(os.Stderr, "  <UNIXTIME>\n")
	fmt.Fprintf(os.Stderr, "      elapled time (sec, msec, nsec) from 1970-01-01 00:00:00\n")
}
