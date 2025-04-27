package utils

import "math/rand"

func RedString(src string) string {
	var Red = "\033[31m"
	var Reset = "\033[0m"
	return Red + src + Reset
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
