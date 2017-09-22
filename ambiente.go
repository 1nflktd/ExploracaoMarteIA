package main

import (
	"fmt"
	"time"
)

const TamanhoMapa = 20

type Caracter string
var C_Diamante Caracter = "*"
var C_Pedra Caracter = "#"
var C_Agente Caracter = "o"
var C_Base Caracter = "B"
var C_Vazio Caracter = " "

type Ambiente struct {
	mapa [TamanhoMapa][TamanhoMapa]Caracter
	pedras []Posicao
	diamantes []Posicao
	agentes []*Agente
	base Posicao
}

func (a *Ambiente) Init() {
	// inicia todos em branco
	for i := 0; i < TamanhoMapa; i++ {
		for w := 0; w < TamanhoMapa; w++ {
			a.mapa[i][w] = C_Vazio
		}
	}

	// coloca a base na posicao certa
	a.base = Posicao{0, 0}
	a.mapa[0][0] = C_Base

	// coloca diamantes (aleatorio)
	a.diamantes = append(a.diamantes, Posicao{5, 7})
	a.mapa[5][7] = C_Diamante
	a.diamantes = append(a.diamantes, Posicao{9, 15})
	a.mapa[9][15] = C_Diamante

	// coloca pedras (aleatorio)
	a.pedras = append(a.pedras, Posicao{4, 17})
	a.mapa[4][17] = C_Pedra
	a.pedras = append(a.pedras, Posicao{8, 10})
	a.mapa[8][10] = C_Pedra

	// coloca agente (aleatorio)
	a.mapa[4][4] = C_Agente
	agente1 := &Agente{}
	agente1.Init(a.base)
	agente1.setPosicaoXY(4, 4)
	a.agentes = append(a.agentes, agente1)
}

func (a *Ambiente) PrintMapa() {
	for i := 0; i < TamanhoMapa; i++ {
		for w := 0; w < TamanhoMapa; w++ {
			fmt.Printf(string(a.mapa[i][w]))
		}
		fmt.Printf("\n")
	}
}

func (a *Ambiente) Run() {
	// laco ate encontrar todos os diamantes
	for {
		limpaTela()
		a.PrintMapa()
		a.moveAgentes()

		time.Sleep(1 * time.Second)
	}
}

func (a *Ambiente) moveAgentes() {
	for _, ag := range a.agentes {
		posAtual := ag.getPosicao()
		a.mapa[posAtual.X][posAtual.Y] = C_Vazio
		p_ag := ag.moveAleatorio()
		// verificaColisao(p_ag)
		a.mapa[p_ag.X][p_ag.Y] = C_Agente
	}
}

func (a *Ambiente) verificaColisao(posAgente Posicao) {

}