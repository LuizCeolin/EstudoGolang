package main

import (
	"testing"
)

func TestOla(t *testing.T) {

    verificaMensagemCorreta := func(t *testing.T, resp, esp string) {
        t.Helper()
        if resp != esp {
            t.Errorf("resultado '%s', esperado '%s'", resp, esp)
        }
    }

    t.Run("diz olá para as pessoas", func(t *testing.T) {
        resp := Ola("Luiz!", "")
        esp := "Olá, Luiz!"
        verificaMensagemCorreta(t, resp, esp)
    })

    t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
        resp := Ola("", "")
        esp := "Olá, Mundo"
        verificaMensagemCorreta(t, resp, esp)
    })

    t.Run("Em espanhol", func(t *testing.T) {
        resp := Ola("Karine", "espanhol")
        esp := "Holla, Karine"
        verificaMensagemCorreta(t, resp, esp)
    })

    t.Run("Em Frances", func(t *testing.T) {
        resp := Ola("Karine", "frances")
        esp := "Bonjour, Karine"
        verificaMensagemCorreta(t, resp, esp)
    })
    
}