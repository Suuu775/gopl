package ex91

type withdrawChan struct{
	amount int
	result chan bool
}

var withdraws = make(chan withdrawChan)
var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int){deposits<-amount}

func Balance() int {return <-balances}

func Withdraw(amount int) bool{
	withdr := withdrawChan{
		amount: amount,
		result: make(chan bool),
	}
	return <-withdr.result
}

func teller(){
	var balance int
	for{
		select{
		case amount := <-deposits:
			balance+=amount
        case balances<-balance:
		case withdr := <-withdraws:
			amount := withdr.amount
			if amount>balance{
				withdr.result<-false
			} else {
				balance -= amount
				withdr.result<-true
			}
		}
	}
}

func Init(){
	go teller()
}
