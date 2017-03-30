package sword

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type V struct {
	Name  string
	Text  string
	Words []string
	Ref   []string
}

var rPattern = regexp.MustCompile(".*:.*:[[:space:]]*(.*)")
var nPattern = regexp.MustCompile("<([HG])([0-9]+)>(.*)")
var notPattern = regexp.MustCompile("<[HG][a-z]>(.*)")

const ot = 2
const nt = 1

func longer(first, second []string) ([]string, int) {
	if len(first) > len(second) {
		return first, 1
	}
	return second, 2
}

func reverse(in string) string {
	runes := []rune(in)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (v *V) appendToWord(index int, word string) {
	v.Words[index] = string(append([]byte(v.Words[index]), []byte(word)...))
}

func (v *V) assign(s string, method int) error {
	p := strings.Split(s, " ")
	for _, r := range p {
		if notPattern.MatchString(r) {
			m := notPattern.FindStringSubmatch(r)
			if len(m[1]) > 0 {
				v.appendToWord(len(v.Words)-1, m[1])
			}
			continue
		}
		m := nPattern.FindStringSubmatch(r)
		if len(m) < 1 {
			v.Words = append(v.Words, r)
		} else {
			test := "StrongsHebrew"
			if method == nt {
				test = "StrongsGreek"
			}
			c := exec.Command("diatheke", "-b", test, "-k", m[2])
			co, err := c.CombinedOutput()
			if err != nil {
				return err
			}
			v.Ref = append(v.Ref, string(co))
			if len(m[3]) > 0 {
				v.appendToWord(len(v.Words)-1, m[3])
			}
		}
	}

	res := strings.Join(v.Words, " ")

	if method == ot {
		v.Text = reverse(res)
	} else {
		v.Text = res
	}
	return nil
}

func Verse(in string) (*V, error) {
	v := &V{Name: in}
	n := exec.Command("diatheke", "-b", "Elzevir", "-o", "n", "-k", in)
	o := exec.Command("diatheke", "-b", "OSHB", "-o", "n", "-k", in)
	nt, nerr := n.CombinedOutput()
	ot, oerr := o.CombinedOutput()
	if nerr != nil {
		return v, nerr
	}
	if oerr != nil {
		return v, oerr
	}
	nres := rPattern.FindStringSubmatch(string(nt))
	ores := rPattern.FindStringSubmatch(string(ot))
	res, which := longer(nres, ores)
	if len(res) < 2 {
		return v, fmt.Errorf("Not found!")
	}
	err := v.assign(res[1], which)
	if err != nil {
		return v, err
	}
	return v, nil
}
