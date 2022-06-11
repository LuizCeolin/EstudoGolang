package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"

const prefixoOlaBr = "Olá, "
const prefixoOlaEs = "Holla, "
const prefixoOlaFr = "Bonjour, "

func Ola(nome string, idioma string) string {

    if nome == "" {
        nome = "Mundo"
    }

    return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {
    switch idioma {
    case espanhol:
        prefixo = prefixoOlaEs
    case frances:
        prefixo = prefixoOlaFr
    default:
        prefixo = prefixoOlaBr
    }

    return
}

func main() {
    fmt.Println(Ola("Luiz!", ""))
}