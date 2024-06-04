package main

import (
	"github.com/punkestu/facade-task/facade"
	"github.com/punkestu/facade-task/worker"
)

func main() {
	var worker1 = &worker.Worker1{Nama: "Bima"}
	var worker2 = &worker.Worker2{Umur: 25}
	var facade = facade.NewFacade(worker1, worker2)

	// tanpa facade
	var nama = worker1.Task1("ganteng")
	var umur = worker2.Task2()
	println("Halo saya " + nama + " " + umur)

	// dengan facade
	facade.Work1()
}
