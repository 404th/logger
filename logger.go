package logger

import "fmt"

func LogWarning(str string) {
	fmt.Printf("WARNING! %s\n", str)
}

func LogInfo(str string) {
	fmt.Printf("INFO! %s\n", str)
}

func LogError(str string) {
	fmt.Printf("ERROR! %s\n", str)
}

