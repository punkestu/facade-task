package worker

import "strconv"

type Worker1 struct {
	Nama string
}

func (w *Worker1) Task1(tambahan string) string {
	return w.Nama + " " + tambahan
}

type Worker2 struct {
	Umur int
}

func (w *Worker2) Task2() string {
	return strconv.Itoa(w.Umur) + " tahun"
}
