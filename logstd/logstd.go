package main

import (
	"log"
	"os"
)

var myLogger *log.Logger

func main() {
	myLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

	//....
	run()

	myLogger.Println("End of Program")
}

func run() {
	myLogger.Print("Test")
}

/* 로그출력 내용
INFO: 2016/01/15 15:27:17 Test
INFO: 2016/01/15 15:27:17 End of Program
*/
