package variablecontext

import "errors"

type NilVariableContext struct {
	variableMap map[string]float64
	Verbose     bool
}

//var TheVariableContext *VariableContext

func (vc *NilVariableContext) SetValue(p string, v float64) {
	vc.variableMap[p] = v
}

func (vc *NilVariableContext) Init(isVerbose bool) {
	vc.Verbose = isVerbose
	vc.variableMap = map[string]float64{}
}

func (vc *NilVariableContext) GetVariableIntValue(s string) (res float64, err error) {
	return 0, errors.New("Undeclared variable \"" + s + "\"")
}

func (vc *NilVariableContext) EvaluateJSONVariableIntValue(s string, path string) (res float64, err error) {
	return vc.GetVariableIntValue(s + "." + path)
}
