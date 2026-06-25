package banner

import (
	"bufio"
	"os"
	"strings"
)

// load opens a font banner file and parses it line by line into a clean slice of strings
func Load(name string) ([]string, error) {
	// look inside the local system assets folder for the specific text banner
	file, err := os.Open("banners/" + name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// strip out hidden Windows carriage returns (\r) to safeguard the thinkertoy banner
		line = strings.ReplaceAll(line, "\r", "")
		lines = append(lines, line)
	}

	// duble check for any structural scanner reading errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
