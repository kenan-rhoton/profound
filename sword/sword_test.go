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
        if verse != data[1] {
            fmt.Printf("\"%s\" is not \"%s\"\n", verse, data[1])
            t.Fail()
        }
    }
}

func TestOldTestamentVerse(t *testing.T) {
    testdata := [][]string{
        []string{"Genesis 5:5", "ויהיו כל־ימי אדם אשׁר־חי תשׁע מאות שׁנה ושׁלשׁים שׁנה וימת׃ ס"},
        []string{"Ex 4:3", "ויאמר השׁליכהו ארצה וישׁליכהו ארצה ויהי לנחשׁ וינס משׁה מפניו׃"},
    }

    for _, data := range testdata {
        verse, err := Verse(data[0])
        if err != nil {
            fmt.Println(err.Error())
            t.Fail()
        }
        if verse != data[1] {
            fmt.Printf("\"%s\" is not \"%s\"\n", verse, data[1])
            t.Fail()
        }
    }
}

func TestIncorrectVerse(t *testing.T) {
    testdata := []string{
        "Geonosis 5:5",
        "Enx 4:3",
        "Xv 5:4",
        "Enx 4:3",
    }

    for _, data := range testdata {
        _, err := Verse(data)
        if err.Error() != "Not found!" {
            fmt.Println(err.Error())
            t.Fail()
        }
    }
}
