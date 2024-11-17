package helpers

import (
	"strconv"
	"strings"
	"time"
)

func GenerateUsername(email string) string {
	splittedEmail := strings.Split(email, "@")
	unixTimestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	return splittedEmail[0] + unixTimestamp
}
