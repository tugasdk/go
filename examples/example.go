package examples

import (
	//"errors"

	"log"
	"log/slog"

	//"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	loggerManager "github.com/tugasdk/go/package"
	"go.uber.org/zap"
)

//var log loggerManager.LogManager = *loggerManager.Default();

func WithLog(){

	logger := log.Default();

	log:= loggerManager.New(logger, true)
	
	log.What("Payment failed").Why("User has no enough balance in account").Where("transaction payment  phase").Who("e-commmerce.payment.pay method").Info();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Warn();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Error();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Debug();
}


func WithSlog(){

	logger := slog.Default();

	log:= loggerManager.New(logger, true)
	
	log.What("Payment failed").Why("User has no enough balance in account").Where("transaction payment  phase").Who("e-commmerce.payment.pay method").Info();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Warn();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Error();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Debug();
}


func WithZap(){

	logger, err := zap.NewDevelopment();

	if err != nil{
		println(" Error setup the zap error")
	}

	log:= loggerManager.New(logger, true)
	
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Info();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Warn();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Error();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Debug();
}

func WithZerolog(){

	logger := zlog.Logger

	log:= loggerManager.New(&logger, true)
	
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Info();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Warn();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Error();
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Debug();
}