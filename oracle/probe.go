package main

import (
	__json "encoding/json"
	__fmt "fmt"
)



import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	regex := regexp.MustCompile("[A-Z]+[a-z']*|[a-z]+")
	words := regex.FindAllString(s, -1)

	var abbr []string
	for _, word := range words {
		abbr = append(abbr, string(word[0]))
	}

	return strings.ToUpper(strings.Join(abbr, ""))
}

type R struct {
	Ok bool        `json:"ok"`
	V  interface{} `json:"v"`
}

func main() {
	inputs := []string{"Portable Network Graphics", "Ruby on Rails", "First In, First Out", "GNU Image Manipulation Program", "Complementary metal-oxide semiconductor", "Rolling On The Floor Laughing So Hard That My Dogs Came Over And Licked Me", "Something - I made up from thin air", "Halley's Comet", "The Road _Not_ Taken"}
	out := []R{}
	for _, x := range inputs {
		func() {
			defer func() { if r := recover(); r != nil { out = append(out, R{false, __fmt.Sprint(r)}) } }()
			out = append(out, R{true, Abbreviate(x)})
		}()
	}
	b, _ := __json.Marshal(map[string]interface{}{"out": out})
	__fmt.Println(string(b))
}
