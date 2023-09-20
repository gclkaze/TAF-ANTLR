package variablecontext

type IVariableContext interface {
	Init(isVerbose bool)
	GetVariableIntValue(s string) float64
	EvaluateJSONVariableIntValue(s string, path string) float64
}
