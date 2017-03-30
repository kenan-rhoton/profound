package sword

import (
    "fmt"
    "testing"
)

func TestNewTestamentVerse(t *testing.T) {
    testdata := [][]string{
        []string{"John 3:16", "ουτως γαρ ηγαπησεν ο θεος τον κοσμον ωστε τον υιον αυτου τον μονογενη εδωκεν ινα πας ο πιστευων εις αυτον μη αποληται αλλ εχη ζωην αιωνιον"},
        []string{"Rev 5:15", "και ειδον οτε ηνοιξεν το αρνιον μιαν εκ των σφραγιδων και ηκουσα ενος εκ των τεσσαρων ζωων λεγοντος ως φωνης βροντης ερχου και βλεπε"},
    }

    for _, data := range testdata {
        verse, err := Verse(data[0])
        if err != nil {
            fmt.Println(err.Error())
            t.Fail()
        }
        if verse.Text != data[1] {
            fmt.Printf("\"%s\"\nis not\n\"%s\"\n", verse.Text, data[1])
            t.Fail()
        }
    }
}

func TestOldTestamentVerse(t *testing.T) {
    testdata := [][]string{
        []string{"Genesis 5:5", "ס ׃תמיו הנׁש םיׁשלׁשו הנׁש תואמ עׁשת יח־רׁשא םדא ימי־לכ ויהיו"},
        []string{"Ex 4:3", "׃וינפמ הׁשמ סניו ׁשחנל יהיו הצרא והכילׁשיו הצרא והכילׁשה רמאיו"},
    }

    for _, data := range testdata {
        verse, err := Verse(data[0])
        if err != nil {
            fmt.Println(err.Error())
            t.Fail()
        }
        if verse.Text != data[1] {
            fmt.Printf("\"%s\"\nis not\n\"%s\"\n", verse.Text, data[1])
            t.Fail()
        }
    }
}

func TestIncorrectVerse(t *testing.T) {
    testdata := []string{
        "Geonosis 5:5",
        "Enx 4:3",
        "Xv 5:4",
        "Pselm 443:3",
    }

    for _, data := range testdata {
        x, err := Verse(data)
        if err != nil {
            if err.Error() != "Not found!" {
                fmt.Println(err.Error())
                t.Fail()
            }
        } else {
            fmt.Println("Missing error:", x)
            t.Fail()
        }
    }
}

func TestHebrewAnalysis(t *testing.T) {
    testdata := []struct{
        verse string
        ref int
        res string
    }{
        struct{
            verse string
            ref int
            res string
        }{"Ex 4:3",
        0,
        `0055900559:  559  'amar  aw-mar'


 a primitive root; to say (used with great
 latitude):--answer, appoint, avouch, bid, boast self, call,
 certify, challenge, charge, + (at the, give) command(-ment),
 commune, consider, declare, demand, X desire, determine, X
 expressly, X indeed, X intend, name, X plainly, promise,
 publish, report, require, say, speak (against, of), X still, X
 suppose, talk, tell, term, X that is, X think, use (speech),
 utter, X verily, X yet.(StrongsHebrew)
`},
    }
    for _, data := range testdata {
        verse, err := Verse(data.verse)
        if err != nil {
            fmt.Println(err.Error())
            t.Fail()
        }
        if len(verse.Ref) <= data.ref{
            fmt.Printf("Not enough references")
            t.Fail()
        } else if verse.Ref[0] != data.res {
            fmt.Printf("\"%s\"\nis not\n\"%s\"\n", verse.Ref[0], data.res)
            t.Fail()
        }
    }
}
