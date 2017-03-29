package sword

import (
    "fmt"
    "os/exec"
    "regexp"
)

type V struct {
    Text string
    Refs []string
}

var rPattern = regexp.MustCompile(".*:.*:[[:space:]]*(.*)")

func longer(first, second []string) ([]string, string) {
    if len(first) > len(second) {
        return first, "first"
    }
    return second, "second"
}

func reverse(in string) string {
    runes := []rune(in)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func Verse(in string) (*V, error) {
    v := &V{}
    n := exec.Command("diatheke", "-b", "Elzevir", "-k", in)
    o := exec.Command("diatheke", "-b", "OSHB", "-k", in)
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
    res, which := longer(nres,ores)
    if len(res) < 2 {
        return v, fmt.Errorf("Not found!")
    }
    if which == "second" {
        v.Text = reverse(res[1])
    } else {
        v.Text = res[1]
    }
    return v, nil
}
