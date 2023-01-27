package project_types

// Options - Тип данных, реализующий структуру для удобного хранения
// входных опций, заданных пользователем
type Options struct {
	Input           string
	Output          string
	UniqParam       uint8
	IParam          bool
	FParamNumFields int
	SParamNumChars  int
}
