package function

import "time"

func GetTimeStr() string {
	return time.Now().Format("2006-01-02 15:03:04")
}

func GetTimeUnix() int64 {
	return time.Now().Unix()
}