package examples
import(
	loggerManager "github.com/tugasdk/go/package"
)


var log loggerManager.LogManager = *loggerManager.Default();

func Examples(){
	
	log.What("Payment failed").Why("System shows not balance was available").Where("When processing the transaction").Who("path/to/error/method").Error();
}