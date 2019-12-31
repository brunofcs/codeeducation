package main

import "testing"

func TestGreetings(t *testing.T) {
    res := greetings("Code.education Rocks !")
    if (res != "<b>Code.education Rocks !</b>") {
       t.Errorf("A mensagem está incorreta, resultado: %s, esperado: %s.", res, "<b>Code.education Rocks !</b>")
    }
}
