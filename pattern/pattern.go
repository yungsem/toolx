package pattern

import "os"

const (
	Date             = "2006-01-02"
	DateTime         = "2006-01-02 15:04:05"
	DateTimeT        = "2006-01-02T15:04:05"
	DateTimeMil      = "2006-01-02 15:04:05.99999"
	DateTimeMilShort = "2006-01-02 15:04:05.999"
	DateTimeMilTight = "2006010215040599999"
	SlashDate        = "2006/01/02"
	SlashDateTime    = "2006/01/02 15:04:05"
	DateTight        = "20060102"
	HourMinute       = "15:04"
)

const (
	NewLine       = "\n"
	Space         = " "
	Tab           = "\t"
	PathSeparator = string(os.PathSeparator)
	Underscore    = "_"
	Equal         = "="
)
