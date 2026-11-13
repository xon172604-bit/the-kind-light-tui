// The Kind Light - a terminal horror story.
// Standard library only, so it builds to one self-contained binary per OS.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	dim   = "\x1b[2m"
	red   = "\x1b[31m"
	amber = "\x1b[33m"
	bold  = "\x1b[1m"
	reset = "\x1b[0m"
	clear = "\x1b[2J\x1b[H"
)

// Choice leads from one scene to another.
type Choice struct {
	Label string
	Next  string
}

// Scene is one beat of the story. An empty Choices slice ends the game.
type Scene struct {
	Text    []string
	Choices []Choice
}

// story is the whole game as data, so adding scenes never touches the engine.
var story = map[string]Scene{
	"start": {
		Text: []string{
			"The power went out an hour ago.",
			"One light is still on, at the end of the hall.",
			"It is warm. It almost looks kind.",
		},
		Choices: []Choice{
			{"Walk toward the light", "hall"},
			{"Stay where you are", "wait"},
		},
	},
	"hall": {
		Text: []string{
			"The floor is cold under your feet.",
			"The light hums. It sounds like someone trying to remember a lullaby.",
		},
		Choices: []Choice{
			{"Reach out and touch it", "touch"},
			{"Turn back", "start"},
		},
	},
	"wait": {
		Text: []string{
			"You count your breaths.",
			"The light does not move. But it is closer than it was.",
		},
		Choices: []Choice{
			{"Look away", "start"},
			{"Look back at it", "hall"},
		},
	},
	"touch": {
		Text: []string{
			"It is warm.",
			"It is so warm.",
			"Somewhere behind you, the dark exhales.",
			"THE END (more scenes coming soon)",
		},
	},
}

var in = bufio.NewReader(os.Stdin)

// slow prints text one rune at a time for the typewriter effect.
func slow(s string, delay time.Duration) {
	for _, r := range s {
		fmt.Print(string(r))
		time.Sleep(delay)
	}
	fmt.Println()
}

func prompt(max int) int {
	for {
		fmt.Printf("%s> %s", amber, reset)
		line, err := in.ReadString('\n')
		if err != nil {
			return 0 // stdin closed: quit cleanly
		}
		line = strings.TrimSpace(line)
		if line == "q" || line == "quit" {
			return 0
		}
		var n int
		if _, err := fmt.Sscanf(line, "%d", &n); err == nil && n >= 1 && n <= max {
			return n
		}
		fmt.Printf("%sChoose 1-%d, or q to quit.%s\n", dim, max, reset)
	}
}

func play() {
	current := "start"
	for {
		sc, ok := story[current]
		if !ok {
			fmt.Println(red + "The story lost its way." + reset)
			return
		}
		fmt.Print(clear)
		for _, line := range sc.Text {
			slow(line, 22*time.Millisecond)
			time.Sleep(350 * time.Millisecond)
		}
		fmt.Println()
		if len(sc.Choices) == 0 {
			return
		}
		for i, c := range sc.Choices {
			fmt.Printf("  %s%d%s) %s\n", bold, i+1, reset, c.Label)
		}
		n := prompt(len(sc.Choices))
		if n == 0 {
			return
		}
		current = sc.Choices[n-1].Next
	}
}

func main() {
	enableANSI()
	fmt.Print(clear)
	fmt.Println(bold + "THE KIND LIGHT" + reset)
	fmt.Println(dim + "a terminal horror story  |  type a number, or q to quit" + reset)
	fmt.Println()
	fmt.Print("Press Enter to begin...")
	in.ReadString('\n')
	play()
	fmt.Println(dim + "\nThanks for playing." + reset)
}
