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
	diamantes int
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
	a.diamantes++
	a.mapa[5][7] = C_Diamante
	a.diamantes++
	a.mapa[9][15] = C_Diamante

	// coloca pedras (aleatorio)
	a.mapa[4][17] = C_Pedra
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

func (a *Ambiente) PrintInfo() {
	for i, ag := range a.agentes {
		fmt.Printf("Agente %d:...", i)
		if ag.getTemDiamante() {
			fmt.Printf("Tem diamante. Voltando para base\n")
		} else {
			fmt.Printf("Procurando diamantes\n")
		}
	}
}

func (a *Ambiente) Run() {
	// laco ate encontrar todos os diamantes
	for {
		limpaTela()
		a.PrintMapa()
		a.PrintInfo()
		a.moveAgentes()
		if a.diamantes == 0 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Todos diamantes foram encontrados e levados para a base!\n")
}

func (a *Ambiente) moveAgentes() {
	for _, ag := range a.agentes {
		posAtual := ag.getPosicao()
		var p_ag Posicao
		if ag.getTemDiamante() {
			p_ag = ag.voltaBase()
		} else {
			p_ag = ag.movePosAleatorio()
		}

		if ok, caracter := a.verificaColisao(p_ag); ok {
			atualizarPos := true
			if caracter == C_Diamante {
				if ag.getTemDiamante() { // ja tem diamante, pula
					atualizarPos = false
				} else {
					ag.setTemDiamante(true)
				}
			}

			if atualizarPos {
				a.mapa[posAtual.X][posAtual.Y] = C_Vazio
				ag.setPosicao(p_ag) // move o elemento
				a.mapa[p_ag.X][p_ag.Y] = C_Agente
			}
		} else if caracter == C_Base {
			if ag.getTemDiamante() {
				ag.setTemDiamante(false)
				a.diamantes--
			}
		}
	}
}

func (a *Ambiente) verificaColisao(posAgente Posicao) (bool, Caracter) {
	c := a.mapa[posAgente.X][posAgente.Y]
	if c == C_Diamante || c == C_Vazio {
		return true, c
	} else {
		return false, c
	}
}
