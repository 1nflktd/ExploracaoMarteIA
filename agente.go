package main

import (
	"time"
	"math/rand"
)

type Agente struct {
	posicao Posicao
	base Posicao
}

func (a *Agente) Init(base Posicao) {
	a.base = base
}

func (a *Agente) setPosicao(p Posicao) {
	a.posicao = p
}

func (a *Agente) setPosicaoXY(x, y int) {
	a.posicao = Posicao{x, y}
}

func (a *Agente) getPosicao() Posicao {
	return a.posicao
}

func (a *Agente) moveAleatorio() Posicao {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if a.posicao.X >= (TamanhoMapa - 1) {
		a.posicao.X--
	} else if a.posicao.X == 0 {
		a.posicao.X++
	} else {
		a.posicao.X += (r.Intn(2) - 1) // random de 0 a 2, se 0 volta uma (-1), 1 fica parado, 2 avanca
	}

	if a.posicao.Y >= (TamanhoMapa - 1) {
		a.posicao.Y--
	} else if a.posicao.Y == 0 {
		a.posicao.Y++
	} else {
		a.posicao.Y += (r.Intn(2) - 1) // random de 0 a 2, se 0 volta uma (-1), 1 fica parado, 2 avanca
	}

	return a.posicao
}

func (a *Agente) voltaBase() {

}