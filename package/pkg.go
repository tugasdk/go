package logManager

import (
	"fmt"
	"log"
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
	//whenMsg string - this is  not need to pass because all log library prints the timestamp
	Verbose bool // default to false, when set to true, it prints the stackTrace fo the error
	Logger interface{}
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

func (logman *LogManager) Error () error{

	/*Error structure example:
	*(what) Enable to add the item on the cart, (why) because there is no stock for this item, (where) the error accoured during validation stock by the validation method on (who) github.com/module/package/struct/method
	*/
	generateMsg := fmt.Sprintf("(what)%s, (why) because %s, (where) in %s, (who) on %s ", logman.whatMsg, logman.whyMsg, logman.whereMsg, logman.whoMsg)
	logman.errorLogFactory(generateMsg)
	return nil
}

func (logman *LogManager) errorLogFactory(errormsg string){
	logger := logman.Logger;

	myDefault, ok := logger.(*log.Logger)

	if ok  {
		myDefault.Fatal(errormsg)
	}

	println("No looger found")

	// Todo, verify for others loggers libary
}

func Package(){
	println("Hello world from package")
}