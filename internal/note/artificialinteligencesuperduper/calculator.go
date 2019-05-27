package artificialinteligencesuperduper

import (
	"strings"
	"time"
)

// Calculator is a super much wow implementation of importance calculator
type Calculator struct{}

// Calculate uses magic powers to calculate importance of title
func (c *Calculator) Calculate(title string) int {
	time.Sleep(time.Second * 5)
	if strings.Contains(title, "gay") {
		return 9001
	}
	return -1
}
