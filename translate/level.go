package translate

import (
	"github.com/MeteorKL/tiger/frame"
	"github.com/MeteorKL/tiger/temp"
	"github.com/MeteorKL/tiger/util"
)

type Level *Level_
type Level_ struct {
	Parent  Level       // 指向上一级的level
	Name    temp.Label  // 这层level的标号
	Frame   frame.Frame // 这一层的level
	Formals AccessList  // frame中每一个参数的 access
}

// 新建一个Level,
func NewLevel(parent Level, name temp.Label, formals util.BoolList) Level {
	level := new(Level_)
	level.Parent = parent
	level.Name = name
	level.Frame = frame.NewFrame(name, &util.BoolList_{true, formals})
	level.Formals = makeFormalAccessList(level)
	return level
}

var outerlevel Level = nil

func Outermost() Level {
	if outerlevel == nil {
		outerlevel = NewLevel(nil, temp.Newlabel(), nil)
	}
	return outerlevel
}

func AllocLocal(level Level, escape bool) Access {
	tr_acc := new(Access_)
	tr_acc.Access = frame.AllocLocal(level.Frame, escape)
	tr_acc.Level = level
	return tr_acc
}
