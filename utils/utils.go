package utils

func RedString(src string) string {
	var Red = "\033[31m"
	var Reset = "\033[0m"
	return Red + src + Reset
}
