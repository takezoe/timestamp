package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tkuchiki/go-timezone"
)

func main() {
	flag.Usage = func() {
		self := os.Args[0]
		usage(self)
	}
	var (
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
	fmt.Print(t.Format("2006-01-02 15:04:05") + " " + *tz)
}

func str2time(value string, sec *bool, msec *bool, nano *bool, loc *time.Location) {
	// timezone
	s := strings.Split(value, " ")
	if len(s) == 3 {
		value = s[0] + " " + s[1]
		offset, _ := timezone.GetOffset(s[2])
		loc = time.FixedZone(s[2], offset)
	}

	patterns := [...]string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"02 Jan 2006 15:04:05",
		"02 Jan 2006 15:04",
		"02 Jan 2006",
	}
	result := int64(-1)

	for i := 0; i < len(patterns); i++ {
		t, _ := time.ParseInLocation(patterns[i], value, loc)
		result = t.UnixNano()
		if result >= 0 {
			break
		}
	}

	if *sec == true {
		result = result / 1000000000
	} else if *msec == true {
		result = result / 1000000
	} else if *nano == true {
		result = result
	} else {
		result = result / 1000000 // default is msec
	}
	fmt.Print(result)
}

func usage(self string) {
	fmt.Fprintf(os.Stderr, "%s: a utility for converting timestamp and unixtime.\n\n", self)
	fmt.Fprintf(os.Stderr, "USAGE::\n")
	fmt.Fprintf(os.Stderr, "  %s [OPTIONS] [TIMESTAMP or UNIXTIME]\n\n", self)
	fmt.Fprintf(os.Stderr, "OPTIONS:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "ARGS:\n")
	fmt.Fprintf(os.Stderr, "  <TIMESTAMP>\n")
	fmt.Fprintf(os.Stderr, "      yyyy-MM-dd [HH:mm[:ss] [timezone]]\n")
	fmt.Fprintf(os.Stderr, "  <UNIXTIME>\n")
	fmt.Fprintf(os.Stderr, "      elapsed time (sec, msec, nsec) from 1970-01-01 00:00:00 UTC\n")
}
