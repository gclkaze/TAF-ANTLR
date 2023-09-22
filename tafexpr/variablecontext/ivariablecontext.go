package variablecontext

type IVariableContext interface {
	Init(isVerbose bool)
	GetVariableIntValue(s string) (float64, error)
	EvaluateJSONVariableIntValue(s string, path string) (float64, error)
}
