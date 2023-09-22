package variablecontext

type VariableContext struct {
}

//var TheVariableContext *VariableContext

func (vc *VariableContext) Init(isVerbose bool) {
}

func (vc *VariableContext) GetVariableIntValue(s string) (res float64, ok error) {
	//log.Println(s, "=>", 1)
	return 1, nil
}

func (vc *VariableContext) EvaluateJSONVariableIntValue(s string, path string) (res float64, ok error) {
	//log.Println(s, path, "=>", 2)
	return 2, nil
}
