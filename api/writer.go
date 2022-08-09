package api

import (
	"fmt"
)

func WriteToScreen(name string, author string, album string) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Now Playing:")
	fmt.Println("Title: " + name)
	fmt.Println("Artist: " + author)
	if album != "" {
		fmt.Println("Album: " + album)
	}
}
