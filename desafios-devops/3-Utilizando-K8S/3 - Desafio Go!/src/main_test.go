package main

import "testing"

func TestPrint(t *testing.T) {

    text := greeting("Code.education Rocks!")
    
    if text != "<b>Code.education Rocks!</b>" {
       t.Errorf("incorrect text, got: %s, want: %s.", text, "<b>Code.education Rocks!</b>")
    }
    
}