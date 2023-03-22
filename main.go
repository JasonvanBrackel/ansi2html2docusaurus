package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReplaceSingleStyle(str string) string {
	re := regexp.MustCompile(`style="([^"]*?)(color|background-color):#([0-9a-f]{3,6})([^"]*?)"`)
	return re.ReplaceAllStringFunc(str, func(match string) string {
		parts := strings.Split(match, ":")
		output := []string{"style={{", parts[0][7:], ": '", parts[1][:len(parts[1])-1], "'}}"}
		return strings.Join(output, "")
	})
}

func ReplaceMultiStyle(str string) string {
	re := regexp.MustCompile(`style="([^"]*?)(color|background-color):#([0-9a-f]{3,6})([^"]*?);([^"]*?)(color|background-color):#([0-9a-f]{3,6})([^"]*?)"`)
	return re.ReplaceAllStringFunc(str, func(match string) string {
		styles := strings.Split(match, ";")
		var replacement []string
		for index, style := range styles {
			var output []string
			parts := strings.Split(style, ":")
			if index == 0 {
				output = append(output, "style={{", parts[0][7:], ": '", parts[1], "'}}")
			} else {
				output = append(output, "style={{", parts[0], ": '", parts[1][:len(parts[1])-1], "'}}")
			}
			replacement = append(replacement, strings.Join(output, ""))
		}

		return strings.Join(replacement, " ")
	})
}

func RemoveExtraSemicolon(str string) string {
	re := regexp.MustCompile(`style={{;`)
	return re.ReplaceAllString(str, "style={{")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = ReplaceMultiStyle(line)
		line = ReplaceSingleStyle(line)
		line = RemoveExtraSemicolon(line)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
		os.Exit(1)
	}
}
