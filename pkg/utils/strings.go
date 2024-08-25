package utils

import "fmt"

const qPattern = "%s-queue"

func GetQ(target string) string {
	return fmt.Sprintf(qPattern, target)
}
