package main

type App struct {
	ambiente Ambiente
}

func (a *App) Run(nDiamantes, nPedras, nAgentes, milissegundos int) {
	initLimpaTela()
	a.ambiente = Ambiente{}
	a.ambiente.Init(nDiamantes, nPedras, nAgentes)
	a.ambiente.Run(milissegundos)
}