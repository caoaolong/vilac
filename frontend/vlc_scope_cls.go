package frontend

type VlcScope struct {
	Type int
	Name string
	List []*VlcToken
	Body []*VlcToken
}

func NewVlcScope(scopeType int, body []*VlcToken) *VlcScope {
	return &VlcScope{
		Type: scopeType,
		Body: body,
	}
}
