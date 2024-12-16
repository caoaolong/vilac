package backend

type VlcPlayerScript struct {
	Code   []string
	Script string
}

func NewVlcPlayerScript(code string) *VlcPlayerScript {

	return &VlcPlayerScript{
		Code: make([]string, 0),
	}
}
