package main

import (
	"github.com/heronhurpia/celeritas"
)

type application struct {
	App *celeritas.Celeritas
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
