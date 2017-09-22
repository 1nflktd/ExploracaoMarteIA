package main

type Agente struct {
	posicao Posicao
}

func (a *Agente) setPosicao(p Posicao) {
	a.posicao = p
}

func (a *Agente) setPosicaoXY(x, y int) {
	a.posicao = Posicao{x, y}
}