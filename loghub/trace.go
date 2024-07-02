package loghub

import (
	"strconv"
)

const DEBUG = 1
const INFO = 1 << 1
const WARNING = 1 << 2
const ERROR = 1 << 3

func FormatStringToNumber(trace string) int {
	traceNum, _ := strconv.Atoi(trace)
	return traceNum
}
