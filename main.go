package main

import (
	"flag"
)

func main() {
	nDiamantes := flag.Int("diamantes", 5, "Numero de diamantes")
	nPedras := flag.Int("pedras", 5, "Numero de pedras")
	nAgentes := flag.Int("agentes", 5, "Numero de agentes")
	milissegundos := flag.Int("milissegundos", 100, "Numero de milissegundos para refresh tela")
	flag.Parse()

	app := &App{}
	app.Run(*nDiamantes, *nPedras, *nAgentes, *milissegundos)
}