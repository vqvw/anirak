package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Colour theme definition
var colours = map[string]string{
	"black":   "#171921",
	"red":     "#ff4766",
	"green":   "#00f58b",
	"yellow":  "#ffdb66",
	"blue":    "#5c85ff",
	"magenta": "#85c2ff",
	"cyan":    "#b8f0ff",
	"white":   "#ebf2ff",

	"hi_black":   "#3a434d",
	"hi_red":     "#ff7a91",
	"hi_green":   "#75ffc3",
	"hi_yellow":  "#ffe799",
	"hi_blue":    "#85a3ff",
	"hi_magenta": "#a8d2ff",
	"hi_cyan":    "#dbf4ff",
	"hi_white":   "#ffffff",

	"black_light":   "#dee1e8",
	"red_light":     "#ff3557",
	"green_light":   "#00d368",
	"yellow_light":  "#bbac00",
	"blue_light":    "#4c6aff",
	"magenta_light": "#3d91ff",
	"cyan_light":    "#2eafff",
	"white_light":   "#2b2e4b",

	"hi_black_light":   "#c0c6d3",
	"hi_red_light":     "#ff6680",
	"hi_green_light":   "#00eb74",
	"hi_yellow_light":  "#f0dc00",
	"hi_blue_light":    "#8a9fff",
	"hi_magenta_light": "#8abdff",
	"hi_cyan_light":    "#85d6ff",
	"hi_white_light":   "#181a2a",

	"real_black":       "#000000",
	"real_black_light": "#ffffff",

	"background0":              "#111127",
	"background1":              "#161731",
	"background1_border":       "#262a4d",
	"background2":              "#1f203d",
	"background2_border":       "#363d6b",
	"background0_light":        "#ffffff",
	"background1_light":        "#f8f8fc",
	"background1_border_light": "#ededf7",
	"background2_light":        "#f0f1fa",
	"background2_border_light": "#c5cddb",

	"transparent":       "#00000000",
	"transparent_light": "#00000000",
	"red40":             "#ff476640",
	"red40_light":       "#ff355740",
	"green40":           "#00f58b40",
	"green40_light":     "#00d36840",
	"blue80":            "#5c85ff80",
	"blue80_light":      "#4c6aff80",
	"blue40":            "#5c85ff40",
	"blue40_light":      "#4c6aff40",
	"blue10":            "#5c85ff10",
	"blue10_light":      "#4c6aff10",
	"magenta40":         "#85c2ff40",
	"magenta40_light":   "#3d91ff40",
	"white80":           "#ebf2ff80",
	"white80_light":     "#2b2e4b80",
	"white40":           "#ebf2ff40",
	"white40_light":     "#2b2e4b40",
	"white20":           "#ebf2ff20",
	"white20_light":     "#2b2e4b20",
	"bluewhite":         "#b3c7ff",
	"bluewhite_light":   "#16216b",
	"bluewhiteA0":       "#b3c7ffA0",
	"bluewhiteA0_light": "#16216bA0",
	"bluewhite80":       "#b3c7ff80",
	"bluewhite80_light": "#16216b80",
	"bluewhite60":       "#b3c7ff60",
	"bluewhite60_light": "#16216b60",
	"bluewhite50":       "#b3c7ff50",
	"bluewhite50_light": "#16216b50",
	"bluewhite40":       "#b3c7ff40",
	"bluewhite40_light": "#16216b40",
	"bluewhite30":       "#b3c7ff30",
	"bluewhite30_light": "#16216b30",
	"bluewhite20":       "#b3c7ff20",
	"bluewhite20_light": "#16216b20",
	"bluewhite10":       "#b3c7ff10",
	"bluewhite10_light": "#16216b10",
	"test":              "#d9ffd9",
}

func main() {
	tplFile := flag.String("t", "", "Template theme file")
	fileout := flag.String("o", "theme.json", "Output theme file")
	colourFormat := flag.String("c", "hex", "Colour format, options: hex [#ff0000], hex0 [0xff0000], rgb [rgb(255, 0, 0)]")

	flag.Parse()

	colourFmt := func(colourName string) (value string) {
		switch *colourFormat {
		case "hex":
			value = colours[colourName]
		case "hex0":
			value = "0x" + colours[colourName][1:]
		case "rgb":
			rgb, err := strconv.ParseUint(colours[colourName][1:], 16, 32)
			check(err)
			value = fmt.Sprintf("rgb(%d, %d, %d)", uint8(rgb>>16), uint8((rgb>>8)&0xFF), uint8(rgb&0xFF))
		default:
			fmt.Println("Invalid colour format")
			os.Exit(1)
		}
		return
	}

	tplFileBytes, err := ioutil.ReadFile(*tplFile)
	check(err)

	re := regexp.MustCompile("//.+")
	tplFileString := re.ReplaceAllString(string(tplFileBytes), "")

	re = regexp.MustCompile("--[^(--)]+--")

	themeFile := re.ReplaceAllStringFunc(tplFileString, func(s string) string {
		colourOptions := strings.Split(s[2:len(s)-2], ":")
		colourName := colourOptions[0]
		return colourFmt(colourName)
	})

	themeFileLight := re.ReplaceAllStringFunc(tplFileString, func(s string) string {
		colourOptions := strings.Split(s[2:len(s)-2], ":")
		colourName := colourOptions[0]
		if len(colourOptions) > 1 {
			colourName = colourOptions[1]
		}
		return colourFmt(colourName)
	})

	err = ioutil.WriteFile(*fileout, []byte(themeFile), 0644)
	check(err)

	err = ioutil.WriteFile(filepath.Dir(*fileout)+"/light-"+filepath.Base(*fileout), []byte(themeFileLight), 0644)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
