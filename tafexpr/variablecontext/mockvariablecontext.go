package variablecontext

import "log"

type MockVariableContext struct {
	variableMap map[string]float64
	Verbose     bool
}

//var TheVariableContext *VariableContext

func (vc *MockVariableContext) SetValue(p string, v float64) {
	vc.variableMap[p] = v
}

func (vc *MockVariableContext) Init(isVerbose bool) {
	vc.Verbose = isVerbose
	vc.variableMap = map[string]float64{}
}

func (vc *MockVariableContext) GetVariableIntValue(s string) float64 {
	val, ok := vc.variableMap[s]
	if !ok {
		log.Println("Didn't find element at " + s)
		panic("Didn't find element at " + s)
	}
	return val
}

func (vc *MockVariableContext) EvaluateJSONVariableIntValue(s string, path string) float64 {
	return vc.GetVariableIntValue(s + "." + path)
}
