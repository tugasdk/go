package internal


type LoggerNotSupported struct{

}

func (lo *LoggerNotSupported) Error(msg string) string{
	return "Log library no support,further implementation may contain the specified log library"
}