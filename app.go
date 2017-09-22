package main

type App struct {
	ambiente Ambiente
}

func (a *App) Run() {
	a.ambiente = Ambiente{}
	a.ambiente.Init()
	a.ambiente.PrintMapa()
}