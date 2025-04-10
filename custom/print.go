package custom

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/fatih/color"
)

type GetParameter struct {
	Url   string
	Param []string
}

var (
	LastLenght         = 0
	MiniBannerSize     = 50
	MiniBannerBlock    = "-"
	MiniBannerPosition = "left"
	MiniBannerColor    = "yellow"
)

func RandomColorPicker() string {
	var colors []string
	colors = append(colors, "blue",
		"red",
		"green",
		"yellow",
		"magenta")
	n := rand.Int() % len(colors)
	return colors[n]
}

func Println(textorecv, cor string) {
	texto := textorecv + "\n"
	Print(texto, cor)
}

func Printm(ntexto []string, cor string) {
	texto := strings.Join(ntexto, " ")
	Print(texto, cor)
}

func Print(texto, cor string) {
	switch cor {
	case "blue":
		color.Blue(texto)
	case "red":
		color.Red(texto)
	case "green":
		color.Green(texto)
	case "yellow":
		color.Yellow(texto)
	case "magenta":
		color.Magenta(texto)
	}
	color.Unset()
}

func Aviso(texto string) {
	Println(texto, "blue")
}

func ExitOnError(aviso string, err error, mostrar bool) {
	if err == nil {
		return
	}
	if mostrar == true {
		erro := aviso + err.Error()
		Println(erro, "red")
	}
	os.Exit(1)
}

func PrintRecursive(txt, cor string) {
	if LastLenght == 0 {
		LastLenght = 50
	}
	texto := "\033[F\r" + strings.Repeat(" ", LastLenght) + "\r" + txt
	switch cor {
	case "blue":
		color.Blue(texto)
	case "red":
		color.Red(texto)
	case "green":
		color.Green(texto)
	case "yellow":
		color.Yellow(texto)
	case "magenta":
		color.Magenta(texto)
	case "":
		fmt.Println("\033[F\r" + strings.Repeat(" ", LastLenght) + "\r")
		color.Unset()
		LastLenght = 0
	}
	LastLenght = len(txt)
}

func MiniBanner(text string) {
	Println("\n\n\n"+strings.Repeat(MiniBannerBlock, MiniBannerSize), MiniBannerColor)
	switch MiniBannerPosition {
	case "left":

		Println(text+strings.Repeat(MiniBannerBlock, MiniBannerSize-len(text)), MiniBannerColor)
	case "right":
		Println(strings.Repeat(MiniBannerBlock, MiniBannerSize-len(text))+text, MiniBannerColor)
	}
	Println(strings.Repeat(MiniBannerBlock, MiniBannerSize), MiniBannerColor)
}
