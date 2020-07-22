package errorGo

import "log"

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func LogFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
