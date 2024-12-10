package main

import (
	"log"
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
	for _, token := range vc.CurrentFile.Tokens {
		log.Println(token)
	}
}
