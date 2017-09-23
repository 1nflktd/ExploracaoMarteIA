package main

type App struct {
	ambiente Ambiente
}

func (a *App) Run(nDiamantes, nPedras, nAgentes int) {
	initLimpaTela()
	a.ambiente = Ambiente{}
	a.ambiente.Init(nDiamantes, nPedras, nAgentes)
	a.ambiente.Run()
}