package frontend

type VlcToken struct {
	Token    string
	Type     int
	IsArray  bool
	Modifier uint8
}

var keywords = []string{
	"return",
}

var dataTypes = []string{
	"void",
	"string",
	"db32",
	"db64",
	"db64",
	"db16",
	"db8",
	"unsigned",
	"signed",
}

const (
	MiConst    = 0b00000001
	MiStatic   = 0b00000010
	MiUnsigned = 0b00000100
	MiSigned   = 0b00001000
	MiVolatile = 0b00010000
)

var modifiers = map[string]uint8{
	"const":    MiConst,
	"static":   MiStatic,
	"unsigned": MiUnsigned,
	"signed":   MiSigned,
	"volatile": MiVolatile,
}
