package logManager

import (
	"fmt"
	//"https://pkg.go.dev/go.uber.org/zap"
	"log"
	"log/slog"

	"github.com/rs/zerolog"

	"go.uber.org/zap"
)

/*
* Used to setup the configuration about logging, like the log framework to use
* The default value of the log framework is the goland log, library
 */
type LogManager struct{
	whyMsg string
	whatMsg string
	whereMsg string // this should try to give context about at which module/package, class, method the erro was originated
	whoMsg string
	trace any
	//whenMsg string - this is  not need to pass because all log library prints the timestamp
	Verbose bool // default to false, when set to true, it prints the stackTrace fo the error
	Logger any
}

func Default() *LogManager{

	return &LogManager{
		Verbose: false,
		Logger: log.Default(),
	}
}

func New (logger interface{}, verbose bool) *LogManager {
	return &LogManager{
		Logger: logger,
		Verbose: verbose,
	}
}


func (logman *LogManager)Why(msg string) *LogManager{

	logman.whyMsg = msg;

	return logman;

}

func (logman *LogManager)Where(msg string) *LogManager{

	logman.whereMsg = msg;

	return logman;

}

func (logman *LogManager)Who(msg string) *LogManager{

	logman.whoMsg = msg;

	return logman;

}

func (logman *LogManager)What(msg string) *LogManager{

	logman.whatMsg = msg;

	return logman;

}

func (logman *LogManager)PrintStackTrace(msg any) *LogManager{

	logman.trace = msg;

	return logman;

}


//Generates the formated message
func (logman *LogManager) generateFormatedMsg() string {

	generateMsg := fmt.Sprintf("{'what' : '%s', 'why' : 'because %s', 'where' : 'in %s', 'who' : 'on %s' ", logman.whatMsg, logman.whyMsg, logman.whereMsg, logman.whoMsg)
	
	if logman.trace == nil{
		generateMsg = fmt.Sprintf("%s}", generateMsg)
	}else{
		generateMsg = fmt.Sprintf("%s, 'stackTrace' : '%s'}", generateMsg, logman.trace)
	}

	return generateMsg;

}



//Send msg to factory which sends to the specified log library
func (logman *LogManager) Error (){

	logman.errorLogFactory(logman.generateFormatedMsg())
}

//Send msg to factory which sends to the specified log library
func (logman *LogManager) Info (){

	logman.infoLogFactory(logman.generateFormatedMsg())
}

//Send msg to factory which sends to the specified log library
func (logman *LogManager) Warn (){

	logman.warningLogFactory(logman.generateFormatedMsg())
}


//Send msg to factory which sends to the specified log library
func (logman *LogManager) Debug (){
	logman.debugLogFactory(logman.generateFormatedMsg())
}



//This factory function generate error message acording to the specified log library by asserting the log library passed through initialization
func (logman *LogManager) errorLogFactory(errormsg string){
	logger := logman.Logger;

	log, ok := logger.(*log.Logger)

	if ok  {
		log.Println(errormsg)
		return
	}

	zapLog, ok:= logger.(*zap.Logger)
	if ok  {
		zapLog.Error(errormsg)
		return
	}

	slog, ok := logger.(*slog.Logger)

	if ok {
		slog.Error(errormsg)
		return
	}

	zlog, ok := logger.(*zerolog.Logger)

	if ok{
		zlog.Error().Msg(errormsg)
		return
	}

	println("No looger or no supported log library was found")

	// Todo, verify for others loggers libary
}



//This factory function generate information message acording to the specified log library by asserting the log library passed through initialization
func (logman *LogManager) infoLogFactory(errormsg string){
	logger := logman.Logger;

	log, ok := logger.(*log.Logger)

	if ok  {
		log.Println(errormsg)
		return
	}

	zapLog, ok:= logger.(*zap.Logger)
	if ok  {
		zapLog.Info(errormsg)
		return
	}

	slog, ok := logger.(*slog.Logger)

	if ok {
		slog.Info(errormsg)
		return
	}

	zlog, ok := logger.(*zerolog.Logger)

	if ok{
		zlog.Info().Msg(errormsg)
		return
	}

	println("No looger or no supported log library was found")

	// Todo, verify for others loggers libary
}


//This factory function generate warning message acording to the specified log library by asserting the log library passed through initialization
func (logman *LogManager) warningLogFactory(errormsg string){
	logger := logman.Logger;

	log, ok := logger.(*log.Logger)

	if ok  {
		log.Println(errormsg)
		return
	}

	zapLog, ok:= logger.(*zap.Logger)
	if ok  {
		zapLog.Warn(errormsg)
		return
	}

	slog, ok := logger.(*slog.Logger)

	if ok {
		slog.Warn(errormsg)
		return
	}

	zlog, ok := logger.(*zerolog.Logger)

	if ok{
		zlog.Warn().Msg(errormsg)
		return
	}

	println("No looger or no supported log library was found")

	// Todo, verify for others loggers libary
}


//This factory function generate debug message acording to the specified log library by asserting the log library passed through initialization
func (logman *LogManager) debugLogFactory(errormsg string){
	logger := logman.Logger;

	log, ok := logger.(*log.Logger)

	if ok  {
		log.Println(errormsg)
		return
	}

	zapLog, ok:= logger.(*zap.Logger)
	if ok  {
		zapLog.Debug(errormsg)
		return
	}

	slog, ok := logger.(*slog.Logger)

	if ok {
		slog.Debug(errormsg)
		return
	}

	zlog, ok := logger.(*zerolog.Logger)

	if ok{
		zlog.Debug().Msg(errormsg)
		return
	}

	println("No looger or no supported log library was found")

	// Todo, verify for others loggers libary
}



