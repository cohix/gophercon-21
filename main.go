package main

import (
	"fmt"
	"time"

	"github.com/suborbital/reactr/rt"
	"github.com/suborbital/reactr/rwasm"
)

func main() {
	r := rt.New()

	r.Register("rust", rwasm.NewRunner("./rust/rust.wasm"))
	r.Register("tinygo", rwasm.NewRunner("./tinygo/tinygo.wasm"))
	r.Register("assemblyscript", rwasm.NewRunner("./assemblyscript/assemblyscript.wasm"))

	rustTime := doJobWithTime(r, rt.NewJob("rust", "5000"))
	fmt.Println("RUST:", rustTime.Milliseconds())

	tinygoTime := doJobWithTime(r, rt.NewJob("tinygo", "5000"))
	fmt.Println("TINYGO:", tinygoTime.Milliseconds())

	assemblyscriptTime := doJobWithTime(r, rt.NewJob("assemblyscript", "5000"))
	fmt.Println("ASSEMBLYSCRIPT:", assemblyscriptTime.Milliseconds())
}

func doJobWithTime(r *rt.Reactr, job rt.Job) time.Duration {
	start := time.Now()

	for i := 0; i < 100000; i++ {
		r.Do(job).Then()
	}

	return time.Since(start)
}
