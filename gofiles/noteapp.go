// mini golang note taking script.

package main

import (
    "fmt"
    "strings"
)

func main () {
    notebook := &notesApplication{author:"Samuel James",}
    notebook.create("This is my first note")
    notebook.create("This is my second note")
    notebook.create("This is my third note")
    notebook.list()
    notebook.get(2)
    notebook.search("third")
    notebook.edit(2, "The fundamental of javascript.")
    notebook.list()
    notebook.delete(1)
}

// instantiate the noteapplication struct object

type notesApplication struct {
    author string
    notes []string
}

// defines struct methods here.

func (napp *notesApplication) create(noteContent string) {
    napp.notes = append(napp.notes, noteContent)
}

func (napp *notesApplication) list() {
    for idx, note := range napp.notes {
        label1 := fmt.Sprintf("%s%d", "Note ID: ", idx + 1)
        label2 := fmt.Sprintf("%s\n", note)
        fmt.Println(label1)
        fmt.Println(label2)
    }
}

func (napp *notesApplication) get(noteId int) {
    note := napp.notes[noteId - 1]
    label := fmt.Sprintf("%s\n", note)
    fmt.Println(label)
}

func (napp * notesApplication) search(searctText string) {
    label := fmt.Sprintf("%s%s\n", "Showing results for search ", searctText)
    fmt.Println(label)

    for idx, note := range napp.notes {
        if strings.Contains(note, searctText) {
            label1 := fmt.Sprintf("%s%d", "Note ID: ", idx + 1)
            label2 := fmt.Sprintf("%s\n", note)
            fmt.Println(label1)
            fmt.Println(label2)
        }
    }
}

func (napp *notesApplication) edit(noteId int, newContent string) {
    napp.notes[noteId - 1] = newContent
}

func (napp *notesApplication) delete(noteId int) {
    napp.notes = append(napp.notes[:noteId - 1], napp.notes[noteId:]...)
    fmt.Println(napp.notes)
}
