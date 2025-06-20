package main
import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}


func main1() {
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}