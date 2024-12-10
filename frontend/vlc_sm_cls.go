package frontend

import (
	"math"
	"sync"
	"vilac/consts"
)

func initSM(end int) [][]uint8 {
	var sm = make([][]uint8, end+1)
	for i := range sm {
		sm[i] = make([]uint8, math.MaxUint8)
		for j := range sm[i] {
			sm[i][j] = math.MaxUint8
		}
	}
	return sm
}

var _funcSM [][]uint8
var _funcSMOnce sync.Once

func FuncSM() ([][]uint8, int) {
	const end = 7
	_funcSMOnce.Do(func() {
		_funcSM = initSM(end)
		var s uint8 = 0
		_funcSM[0][consts.TokenTypeDataType] = s + 1
		s++
		_funcSM[1][consts.TokenTypeIdentifier] = s + 1
		s++
		_funcSM[2]['('] = s + 1
		s++
		_funcSM[3][consts.TokenTypeDataType] = s + 1
		s++
		_funcSM[4][consts.TokenTypeIdentifier] = s + 1
		s++
		_funcSM[5][','] = s - 1
		_funcSM[5][')'] = s + 1
		s++
		_funcSM[6][';'] = s + 1
		_funcSM[6]['{'] = s + 1
	})
	return _funcSM, end
}

var _callSM [][]uint8
var _callSMOnce sync.Once

func CallSM() ([][]uint8, int) {
	const end = 6
	_callSMOnce.Do(func() {
		_callSM = initSM(end)
		var s uint8 = 0
		_callSM[0][consts.TokenTypeIdentifier] = s + 1
		s++
		_callSM[1]['('] = s + 1
		s++
		_callSM[2][consts.TokenTypeString] = s + 1
		_callSM[2][consts.TokenTypeNumber] = s + 1
		_callSM[2][consts.TokenTypeIdentifier] = s + 1
		_callSM[2][')'] = s + 3
		s++
		_callSM[3][','] = s + 1
		_callSM[3][')'] = s + 2
		s++
		_callSM[4][consts.TokenTypeString] = s - 1
		_callSM[4][consts.TokenTypeNumber] = s - 1
		_callSM[4][consts.TokenTypeIdentifier] = s - 1
		s++
		_callSM[5][';'] = s + 1
	})
	return _callSM, end
}
