# Simple Time formatting in Go

package "time"


## Predefined Layouts
**1. time.ANSIC**
- Layout: `"Mon Jan _2 15:04:05 2006"`
- Example: `Mon Jan 2 15:04:05 2006`

time.UnixDate
Layout: "Mon Jan _2 15:04:05 MST 2006"
Example: Mon Jan 2 15:04:05 MST 2006

time.RubyDate
Layout: "Mon Jan 02 15:04:05 -0700 2006"
Example: Mon Jan 02 15:04:05 -0700 2006

time.RFC822
Layout: "02 Jan 06 15:04 MST"
Example: 02 Jan 06 15:04 MST
time.RFC822Z

Layout: "02 Jan 06 15:04 -0700"
Example: 02 Jan 06 15:04 -0700
time.RFC850

Layout: "Monday, 02-Jan-06 15:04:05 MST"
Example: Monday, 02-Jan-06 15:04:05 MST
time.RFC1123

Layout: "Mon, 02 Jan 2006 15:04:05 MST"
Example: Mon, 02 Jan 2006 15:04:05 MST
time.RFC1123Z

Layout: "Mon, 02 Jan 2006 15:04:05 -0700"
Example: Mon, 02 Jan 2006 15:04:05 -0700
time.RFC3339

Layout: "2006-01-02T15:04:05Z07:00"
Example: 2006-01-02T15:04:05Z07:00
time.RFC3339Nano

Layout: "2006-01-02T15:04:05.999999999Z07:00"
Example: 2006-01-02T15:04:05.999999999Z07:00
time.Kitchen

Layout: "3:04PM"
Example: 3:04PM
time.Stamp

Layout: "Jan _2 15:04:05"
Example: Jan 2 15:04:05
time.StampMilli

Layout: "Jan _2 15:04:05.000"
Example: Jan 2 15:04:05.000
time.StampMicro

Layout: "Jan _2 15:04:05.000000"
Example: Jan 2 15:04:05.000000
time.StampNano

Layout: "Jan _2 15:04:05.000000000"
Example: Jan 2 15:04:05.000000000