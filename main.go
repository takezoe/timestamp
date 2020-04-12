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
		tz   = flag.String("tz", "", "TimeZone (e.g. JST, UTC, PST)")
		out  = flag.String("out", "", "Output TimeZone (e.g. JST, UTC, PST)")
		sec  = flag.Bool("sec", false, "Output TimeUnit (sec)")
		msec = flag.Bool("msec", true, "Output TimeUnit (msec)")
		nano = flag.Bool("nano", false, "Output TimeUnit (nano)")
	)
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		showNow(*tz, getTimeunit(*sec, *msec, *nano))
		os.Exit(0)
	}

	if *tz == "" {
		*tz = "UTC"
	}
	loc := getLocation(*tz)

	value := args[0]
	i, e := strconv.ParseInt(value, 10, 64)

	if e == nil {
		time2str(i, *tz, loc)
	} else {
		str2time(value, getTimeunit(*sec, *msec, *nano), loc, *out)
	}
}

func showNow(tz string, unit int64) {
	now := time.Now().UnixNano()
	if tz == "" {
		fmt.Println(now / unit)
	} else {
		time2str(now, tz, getLocation(tz))
	}
}

func getLocation(tz string) time.Location {
	offset, e := timezone.GetOffset(tz)
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
	loc := time.FixedZone(tz, offset)
	return *loc
}

func getTimeunit(sec bool, msec bool, nano bool) int64 {
	if sec == true {
		return 1000000000
	} else if msec == true {
		return 1000000
	} else if nano == true {
		return 1
	} else {
		return 1000000 // default is msec
	}
}

func time2str(value int64, tz string, loc time.Location) {
	// sec to msec
	if value < 10000000000 {
		value = value * 1000
	}
	// msec to nsec
	if value < 10000000000000 {
		value = value * 1000000
	}
	t := time.Unix(0, value).In(&loc)
	fmt.Println(t.Format("2006-01-02 15:04:05") + " " + tz)
}

func str2time(value string, unit int64, loc time.Location, out string) {
	// timezone
	s := strings.Split(value, " ")
	offset, e := timezone.GetOffset(s[len(s)-1])
	if e != nil {
		loc = *time.FixedZone(s[len(s)-1], offset)
	}

	result := int64(-1)

	patterns := [...]string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"2006/01/02 15:04:05",
		"2006/01/02 15:04",
		"2006/01/02",
		"02 Jan 2006 15:04:05",
		"02 Jan 2006 15:04",
		"02 Jan 2006",
	}

	for i := 0; i < len(patterns); i++ {
		t, _ := time.ParseInLocation(patterns[i], value, &loc)
		result = t.UnixNano()
		if result >= 0 {
			break
		}
	}

	if result < 0 {
		fmt.Fprintln(os.Stderr, "Invalid input: "+value)
		os.Exit(1)
	}

	if out != "" {
		time2str(result, out, getLocation(out))
	} else {
		fmt.Println(result / unit)
	}
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
	fmt.Fprintf(os.Stderr, "      yyyy/MM/dd [HH:mm[:ss] [timezone]]\n")
	fmt.Fprintf(os.Stderr, "      yyyy-MM-dd [HH:mm[:ss] [timezone]]\n")
	fmt.Fprintf(os.Stderr, "      dd MMM yyyy [HH:mm[:ss] [timezone]]\n")
	fmt.Fprintf(os.Stderr, "  <UNIXTIME>\n")
	fmt.Fprintf(os.Stderr, "      elapsed time (sec, msec, nsec) from 1970-01-01 00:00:00 UTC\n")
}
