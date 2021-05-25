package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func FoundKey(key string) {
	c := color.New(color.FgGreen)
	c.Print("[+] ")
	fmt.Print("Key Found : ")
	c = color.New(color.FgYellow)
	c.Println(key)
}
