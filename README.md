# go

## The Go sdk for the tuga semantic logic

## About
The logs we printint are not there just to show a bunch of text, each log printed tells us a:
> **History about a event that occours in certain period of time.**

Computer science is a young study field compared with others like engineering for example and we can notice that most of the time what computer science try to do is to emulate the way things are structured and implement in out society. 

One example is the software development itself which trys to emulate how insfrastructure build (Civil engeneering) are constructed by first having the architecture design and use the architecture as the base do build and implement the software. 

With that in mind, as we said early logs tells a history about an ocoured event, why no tell this history using models that others sciency fields use to tell history about a event? We can emulate how historian does to tell a hostory about a event that happen in past, example: 


That is way we create **tuga**, *a semantic logging  framework* that will allow to write logs about business transation intuitivelly and sematically.

The whole goal of **tuga** is to easy the process of debugging production issue by given the log event enough context and information.


## How Tuga will achieve his goals?
**Tugas** use the well known method to discribe and related historic event, **"the 5Ws"**, with 5W we can describe:

- **What** happens
- **Why** happens
- **Where** it happens
- To **Who** happens (the routine, method, function where the error happens) and
- **When**, this later we get from granted from many log library

In Context of business transation, we can provide the logs in that formate and easy the process of debug production issues.

Ex: During payement phase in a transaction in an e-commerce system a error can happen that user try to execute the payment, we may use the following log structure to present the log:

```json 
2024/04/07 13:11:04 INFO {'what' : 'Payment failed', 'why' : 'because User has no enough balance in account', 'where' : 'in transaction payment  phase', 'who' : 'on e-commmerce.payment.pay method' }
```

The **Tuga** library will provide a framework that will allow developers to write code the semantics logs and formate in the format that you see bellow.

### Go implementation example

```go
package examples

import (
	"log/slog"
    logManager "github.com/tugasdk/go"
)

    //Initialize you prefered log library, here we will use the builtin go logging slog library
	logger := slog.Default();



    //Setup the tuga loggerManage by passing the log library you system works with in the first argument and a boolean in the second argument which setup the verbosity
	log:= loggerManager.New(logger, true)
	
	log.What("Payment failed").Why("User has no enough balance in account").Where("transaction payment  phase").Who("e-commmerce.payment.pay method").Info();
}
```

The log output will be:

```json 
2024/04/07 13:11:04 INFO {'what' : 'Payment failed', 'why' : 'because User has no enough balance in account', 'where' : 'in transaction payment  phase', 'who' : 'on e-commmerce.payment.pay method' }
```

Take a look at example folder to see who to implement using others log library

## Supported log library
The go sdk currently support the following library

- log
- logs
- uber zap
- zerolog

Further implemetation will support others log library
