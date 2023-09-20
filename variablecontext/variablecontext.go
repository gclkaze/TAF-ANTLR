package variablecontext

type VariableContext struct {
}

//var TheVariableContext *VariableContext

func (vc *VariableContext) Init(isVerbose bool) {
}

func (vc *VariableContext) GetVariableIntValue(s string) float64 {
	//log.Println(s, "=>", 1)
	return 1
}

func (vc *VariableContext) EvaluateJSONVariableIntValue(s string, path string) float64 {
	//log.Println(s, path, "=>", 2)
	return 2
}
