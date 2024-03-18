package word

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func Init(fileName string) ([]string, error) {
	words := []string{}

	file, err := os.Open(fmt.Sprintf("./modes/%s", fileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if len(words) < 15 {
			word := sc.Text()
			words = append(words, word)
			continue
		}
		break
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })
	return words, nil
}
