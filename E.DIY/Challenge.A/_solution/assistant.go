package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)

	help := flag.Bool("help", false, "Show help message")
	timeOption := flag.Bool("time", false, "Print current time")
	hello := flag.Bool("hello", false, "Print a greeting message")
	inspireMe := flag.Bool("inspire-me", false, "Get inspired by a quote")

	flag.Parse()

	if *help {
		printHelp()
		return
	}

	if *timeOption {
		fmt.Printf("Current time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	}

	if *hello {
		fmt.Println("Hello. Congrats on completing the exercise. You are now a Rockstar!")
	}

	if *inspireMe {
		inspirationalQuotes := []string{
			"The Force will be with you, always. - Obi-Wan Kenobi, Star Wars",
			"Life is like a box of chocolates; you never know what you're gonna get. - Forrest Gump, Forrest Gump",
			"Oh yes, the past can hurt. But the way I see it, you can either run from it, or learn from it. - Rafiki, The Lion King",
			"To infinity and beyond! - Buzz Lightyear, Toy Story",
			"May the odds be ever in your favor. - Effie Trinket, The Hunger Games",
			"Just keep swimming. - Dory, Finding Nemo",
			"Carpe Diem. Seize the day, boys. Make your lives extraordinary. - John Keating, Dead Poets Society",
			"Even the darkest night will end and the sun will rise. - Victor Hugo, Les Misérables",
			"Life moves pretty fast. If you don't stop and look around once in a while, you could miss it. - Ferris Bueller, Ferris Bueller's Day Off",
			"We are Groot. - Groot, Guardians of the Galaxy",
		}

		rand.New(rand.NewSource(time.Now().UnixNano()))
		randomQuote := inspirationalQuotes[rand.Intn(len(inspirationalQuotes))]
		fmt.Println(randomQuote)
	}
}

func printHelp() {
	fmt.Println("Usage: assistant [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}
