package backend

import (
	"sync"
	"vilac/frontend"
)

type VlcGenerator struct {
	Sources []*frontend.VlcFile
	Targets []*VlcPlayerScript
}

func (vg *VlcGenerator) Generate() {

}

var _generator *VlcGenerator
var _generatorOnce sync.Once

func NewVlcGenerator(files []*frontend.VlcFile) *VlcGenerator {
	_generatorOnce.Do(func() {
		_generator = &VlcGenerator{
			Sources: files,
		}
	})
	return _generator
}
