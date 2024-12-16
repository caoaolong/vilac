package main

import (
	"log"
	"vilac/backend"
	"vilac/frontend"
)

func main() {
	var vc = frontend.NewVlcCompiler([]frontend.VlcFile{
		{
			Name: "scripts/1.vlc",
		},
	})
	err := vc.Compile()
	if err != nil {
		log.Fatal(err)
	}
	var vg = backend.NewVlcGenerator(vc.Sources())
	vg.Generate()
}
