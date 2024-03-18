package game

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/olegtemek/typing-test-cli/internal/utils"
)

type Game struct {
	words       []string
	c           chan rune
	d           chan bool
	lastWord    []byte
	indexInWord int
	indexInText int
	writedWords []string
	start       time.Time
	end         time.Time
}

func New(words []string) *Game {
	return &Game{
		words:       words,
		writedWords: make([]string, len(words)),
		d:           make(chan bool),
		c:           make(chan rune),
	}
}

func (g *Game) Start() {

	go g.startKeyboardReader()

	for {
		g.prints()
		if g.indexInText == len(g.words)-1 && g.indexInWord == len(g.words[g.indexInText]) {
			g.writedWords[g.indexInText] = string(g.lastWord)
			close(g.d)
			break
		}

		val, ok := <-g.c
		if !ok {
			break
		}

		if val == 0 {
			continue
		}

		if len(g.lastWord) == 0 {
			g.lastWord = make([]byte, len(g.words[g.indexInText]))
			g.start = time.Now()
		}

		if len(g.lastWord) > g.indexInWord {
			g.lastWord[g.indexInWord] = byte(val)

			g.writedWords[g.indexInText] = string(g.lastWord)
			g.indexInWord++
		}

	}

	g.end = time.Now()
}

func (g *Game) Stats() {
	elapsed := g.end.Sub(g.start)
	correctedWords := 0

	for i, word := range g.words {
		if word == g.writedWords[i] {
			correctedWords++
		}
	}

	wpm := math.Round((float64(correctedWords) / elapsed.Seconds()) * 60)

	fmt.Printf("\n\nWRONG WORDS: ")
	utils.ColorizePrint(fmt.Sprintf("%d", len(g.words)-correctedWords), "red")
	fmt.Printf("\nTIME: ")
	utils.ColorizePrint(elapsed.String(), "yellow")

	fmt.Printf("\nAVERAGE SPEED: %v wpm\n", wpm)
}

func (g *Game) restart() {
	g.writedWords = make([]string, len(g.words))
	g.lastWord = []byte{}
	g.indexInText = 0
	g.indexInWord = 0
}

func (g *Game) addSpace() {
	if g.indexInText < len(g.words)-1 {
		for i, b := range g.lastWord {
			if b == 0 {
				g.lastWord[i] = byte(' ')
			}
		}

		if len(g.lastWord) > 0 {
			g.writedWords[g.indexInText] = string(g.lastWord)
			g.indexInText++
			g.lastWord = make([]byte, len(g.words[g.indexInText]))
			g.indexInWord = 0
		}

	}
}

func (g *Game) deleteChar() {
	if g.indexInWord <= len(g.words[g.indexInText]) {

		if g.indexInWord > 0 {
			g.indexInWord--
			g.lastWord[g.indexInWord] = 0
			g.writedWords[g.indexInText] = string(g.lastWord)
			return
		}

		if g.indexInText > 0 {
			g.writedWords[g.indexInText] = ""
			g.indexInText--
		}
		g.lastWord = []byte(g.writedWords[g.indexInText])
		g.indexInWord = len(g.lastWord)

	}
}

func (g *Game) startKeyboardReader() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
		os.Exit(1)
	}()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in keyboard reader", r)
		}
	}()

	for {
		select {
		case <-g.d:
			return
		default:
			char, key, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}

			g.c <- char

			if key == keyboard.KeySpace {
				g.addSpace()
			}

			if key == keyboard.KeyBackspace || key == keyboard.KeyDelete || key == keyboard.KeyBackspace2 {
				g.deleteChar()
			}
			if key == keyboard.KeyTab {
				g.restart()
			}

			if key == keyboard.KeyCtrlC {
				close(g.d)
			}

		}
	}
}

func (g *Game) prints() {
	utils.ClearTerminal()
	utils.ColorizePrint("for restart game, press TAB\n", "red")
	text := strings.Join(g.words, " ")
	fmt.Println(text)

	writedText := strings.Join(g.writedWords, " ")
	fmt.Println(writedText)

}
