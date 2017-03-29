package sword

import (
    "fmt"
    "os/exec"
    "regexp"
)

var rPattern = regexp.MustCompile(".*:.*:[[:space:]]*(.*)")

func longer(first, second []string) []string {
    if len(first) > len(second) {
        return first
    }
    return second
}

func Verse(in string) (string, error) {
    n := exec.Command("diatheke", "-b", "Elzevir", "-o", "mn", "-k", in)
    o := exec.Command("diatheke", "-b", "OSHB", "-k", in)
    nt, nerr := n.CombinedOutput()
    ot, oerr := o.CombinedOutput()
    if nerr != nil {
        return "", nerr
    }
    if oerr != nil {
        return "", oerr
    }
    nres := rPattern.FindStringSubmatch(string(nt))
    ores := rPattern.FindStringSubmatch(string(ot))
    res := longer(nres,ores)
    if len(res) < 2 {
        return "", fmt.Errorf("Not found!")
    }
    return res[1], nil
}
