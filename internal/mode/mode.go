package mode

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/olegtemek/typing-test-cli/internal/utils"
)

func Init() (string, error) {

	availableModes, err := getAvailableModes()
	if err != nil {
		return "", err
	}

	mode := int(0)

	errorInput := false

	for {
		if !errorInput {
			utils.ClearTerminal()
			utils.ColorizePrint("Type number mode\n", "red")
			for index, mode := range availableModes {
				title := strings.Split(mode, ".txt")[0]
				utils.ColorizePrint(fmt.Sprintf("%v. %s\n", index+1, title), "yellow")
			}
		}
		fmt.Scanf("%d", &mode)

		if mode <= 0 || mode > len(availableModes) {
			if errorInput {
				errorInput = false
				continue
			}

			utils.ColorizePrint("ERROR, corrent number mode", "red")
			errorInput = true
			continue
		}

		break

	}

	return availableModes[mode-1], nil
}

func getAvailableModes() ([]string, error) {
	entries, err := os.ReadDir("./modes")
	files := []string{}
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if filepath.Ext(entry.Name()) != ".txt" {
			continue
		}

		files = append(files, entry.Name())
	}

	return files, nil
}
