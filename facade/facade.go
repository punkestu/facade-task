package facade

import "github.com/punkestu/facade-task/worker"

type Facade struct {
	worker1 *worker.Worker1
	worker2 *worker.Worker2
}

func NewFacade(worker1 *worker.Worker1, worker2 *worker.Worker2) *Facade {
	return &Facade{
		worker1: worker1,
		worker2: worker2,
	}
}

func (f *Facade) Work1() {
	var nama = f.worker1.Task1("ganteng")
	var umur = f.worker2.Task2()
	println("Halo saya " + nama + " " + umur)
}
