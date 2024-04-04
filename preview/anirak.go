package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/vqvw/utils/colour"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	port := flag.String("port", "3000", "Port to run the server on")
	sourceFile := flag.String("source", "source.html", "File containing the source.html file")
	themes := flag.String("themes", "night,day", "Comma-separated of themes to parse")
	template := flag.String("template", "", "Template file to run replacer on")
	dir := flag.String("dir", ".", "Directory to write output to")
	formatFlag := flag.String("format", "hex", "Options: hex, rgba, rgb")

	flag.Parse()

	if *template == "" {
		fmt.Println("No template file provided")
		os.Exit(1)
	}

	themesArray := strings.Split(*themes, ",")
	printQueue := make(map[string]string)

	templateVars := map[string]map[string]string{
		"night": {"name": "Anirak Next (night)", "type": "dark"},
		"day":   {"name": "Anirak Next (day)", "type": "light"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *sourceFile)
	})

	http.HandleFunc("/colours", func(w http.ResponseWriter, r *http.Request) {
		var anirakColours map[string]map[string]string

		err := json.NewDecoder(r.Body).Decode(&anirakColours)
		panicIf(err)

    fmt.Println(anirakColours)

		templateBytes, err := os.ReadFile(*template)
		panicIf(err)

		reComments := regexp.MustCompile("//.+")
		reWhitespace := regexp.MustCompile("\\s")
		reSubstitutions := regexp.MustCompile("{{(.+?)}}")
		reTemplateVars := regexp.MustCompile("\\$[^\\s\"]+")

		templateOutput := string(templateBytes)
		templateOutput = reComments.ReplaceAllString(templateOutput, "")
		templateOutput = reWhitespace.ReplaceAllString(templateOutput, "")

		for _, theme := range themesArray {
			printQueue[theme] = templateOutput

			printQueue[theme] = reTemplateVars.ReplaceAllStringFunc(printQueue[theme], func(s string) string {
				return templateVars[theme][s[1:]]
			})

			printQueue[theme] = reSubstitutions.ReplaceAllStringFunc(printQueue[theme], func(s string) string {
				s = s[2 : len(s)-2]

				variants := strings.Split(s, ",")
				variantsMap := make(map[string]string)

				for _, variant := range variants {
					variantSplit := strings.Split(variant, ":")
					variantName := variantSplit[0]
					variantColour := variantSplit[1]

					anirakColour, err := colour.Parse(anirakColours[theme][variantColour])
          // fmt.Println(anirakColours[theme][variantColour])
					panicIf(err)

					anirakColourFormatted, err := anirakColour.To(*formatFlag)
					panicIf(err)

					variantsMap[variantName] = anirakColourFormatted
				}

				if variantColour, exists := variantsMap[theme]; exists {
					return variantColour
				}

				var firstVariantColour string
				for _, colour := range variantsMap {
					firstVariantColour = colour
					break
				}

				return firstVariantColour
			})
		}

		for themeName, parsedFile := range printQueue {
			os.WriteFile(path.Join(*dir, fmt.Sprintf("anirak_%s%s", themeName, path.Ext(*template))), []byte(parsedFile), 0644)
		}

		fmt.Println("Theme files created.")
	})

	fmt.Println("Listening on", *port+"...")
	err := http.ListenAndServe(":"+*port, nil)
	panicIf(err)
}
