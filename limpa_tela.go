package main

import (
    "os"
    "os/exec"
    "runtime"
)

var mFuncOS map[string]func()

func initLimpaTela() {
    mFuncOS = make(map[string]func())
    mFuncOS["linux"] = func() { 
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    mFuncOS["windows"] = func() {
        cmd := exec.Command("cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func limpaTela() {
    f, ok := mFuncOS[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok {
        f()
    } else {
        panic("Plataforma nao suportada")
    }
}
