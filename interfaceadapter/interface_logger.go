package interfaceadapter

type Logger interface {
	Printf(format string, args ...interface{})
	LogWithFields(map[string]interface{}, string)
}
