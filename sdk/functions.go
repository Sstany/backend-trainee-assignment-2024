package sdk

import (
	"strconv"
	"strings"

	"github.com/lib/pq"
)

func IsDublicateTableErr(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "duplicate_table" {
		return true
	}

	return false
}

func CreateKey(feature, tag int) string {
	var s strings.Builder

	s.WriteString(strconv.Itoa(feature))
	s.WriteString("_")
	s.WriteString(strconv.Itoa(tag))

	return s.String()
}

func CreateKeyFromString(feature, tag string) string {
	var s strings.Builder

	s.WriteString(feature)
	s.WriteString("_")
	s.WriteString(tag)

	return s.String()
}
