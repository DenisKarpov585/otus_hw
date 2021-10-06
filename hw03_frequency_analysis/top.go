package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

type Countword struct {
	Name  string
	Count int
}

type Everyword []Countword

func (c *Countword) Plusone() {
	c.Count++
}

func SortSameCount(slc Everyword) {
	for i := range slc {
		for j := range slc {
			if slc[j].Count == slc[i].Count {
				if slc[i].Name < slc[j].Name {
					slc[i], slc[j] = slc[j], slc[i]
				}
			}
		}
	}
}

func OnlyLetters(s string) string {
	var b strings.Builder
	var buf string
	for n, i := range s {
		if unicode.IsLetter(i) {
			if buf != "" {
				b.WriteString(buf + string(i))
				buf = ""
			} else {
				b.WriteString(string(i))
			}
		} else if n != 0 {
			buf += string(i)
		}
	}
	return b.String()
}

func Top10(input string) []string {
	var answer []string
	if input == "" {
		return answer
	}
	newstr := strings.Fields(input)
	var slc Everyword
	var flag bool
	for _, v := range newstr {
		flag = true
		for i := range slc {
			if slc[i].Name == v {
				slc[i].Plusone()
				flag = false
				break
			}
		}
		if flag {
			slc = append(slc, Countword{v, 1})
		}
	}
	sort.Slice(slc, func(i, j int) bool {
		return slc[i].Count > slc[j].Count
	})
	SortSameCount(slc)
	if len(slc) < 10 {
		for i := 0; i < len(slc); i++ {
			answer = append(answer, slc[i].Name)
		}
	} else {
		for i := 0; i < 10; i++ {
			answer = append(answer, slc[i].Name)
		}
	}
	return answer
}

func Top10custom(input string) []string {
	var answer []string
	if input == "" {
		return answer
	}
	newstr := strings.Fields(input)
	var slc Everyword
	var flag bool
	for _, v := range newstr {
		v = strings.ToLower(v)
		v = OnlyLetters(v)
		flag = true
		for i := range slc {
			if slc[i].Name == v && slc[i].Name != "" {
				slc[i].Plusone()
				flag = false
				break
			}
		}
		if flag && v != "" {
			slc = append(slc, Countword{v, 1})
		}
	}
	sort.Slice(slc, func(i, j int) bool {
		return slc[i].Count > slc[j].Count
	})
	SortSameCount(slc)
	if len(slc) < 10 {
		for i := 0; i < len(slc); i++ {
			answer = append(answer, slc[i].Name)
		}
	} else {
		for i := 0; i < 10; i++ {
			answer = append(answer, slc[i].Name)
		}
	}
	return answer
}
