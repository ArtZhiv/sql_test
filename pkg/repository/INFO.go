package repository

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/gookit/color"
	"gopkg.in/yaml.v3"
)

var (
	command     = color.New(color.FgRed, color.OpBold, color.OpUnderscore)
	description = color.New(color.FgYellow, color.OpItalic)
)

// InfoV ...
func InfoV() {
	description.Print(`
  | beCloud NOC
  +
  | +375 44 508 1535 [A1]               ОТДЕЛ ОПЕРАТИВНОГО КОНТРОЛЯ СЕТИ
  | +375 25 988 1535 [Life:)]           Управление эксплуатации сети
  | +375 29 233 1535 [MTS]
  | +375 17 399 3663 [Beltelecom]
  | +375 17 327 4444 [Fax]              noc@becloud.by   [e-mail]` + "\n",
	)
	fmt.Println()
}

// GenYaml ...
func GenYaml() []prompt.Suggest {
	var s []prompt.Suggest
	allSelect := make(map[string]string)
	openFile, err := ioutil.ReadFile("../files/info.yaml")
	if err != nil {
		Error.Printf("ERROR")
		fmt.Printf(" open YAML info: %v\n", err)
	}
	yaml.Unmarshal(openFile, &allSelect)
	for key, value := range allSelect {
		t := prompt.Suggest{
			Text:        fmt.Sprintf("%v", key),
			Description: value,
		}
		s = append(s, t)
	}
	return s
}

func GetHelp() {
	allSelect := make(map[string]string)
	openFile, err := ioutil.ReadFile("../files/help.yaml")
	if err != nil {
		Error.Printf("ERROR")
		fmt.Printf(" open YAML help: %v\n", err)
	}
	yaml.Unmarshal(openFile, &allSelect)
	for key, value := range allSelect {
		command.Printf("%v:\n", key)
		if strings.Contains(value, "|") {
			w := strings.Split(value, "|")
			for _, elem := range w {
				description.Printf("\t%v\n", elem)
			}
			fmt.Println()
		} else {
			description.Printf("\t%v\n\n", value)
		}
	}
}
