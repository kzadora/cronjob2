package naming

import (
	"fmt"
	"time"
)

const (
	// A child Job will have a name with a suffix like "-1234567890", which is the Unix epoch.
	// For foreseeable future the number of digits in the epoch will be equal to 10,
	// so the total length of the suffix is 11.
	ChildJobSuffixLength = 11
)

func CreateJobName(baseName string, startTime time.Time) string {
	name := fmt.Sprintf("%s-%d", baseName, startTime.Unix())
	return name
}
