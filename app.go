package main

type App struct {
	ambiente Ambiente
}

func (a *App) Run() {
	initLimpaTela()
	a.ambiente = Ambiente{}
	a.ambiente.Init()
	a.ambiente.Run()
}