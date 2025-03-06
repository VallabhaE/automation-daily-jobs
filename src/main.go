package main

import (
	"fmt"
	logreader "main/src/utils/logReader"
)


const (
	message = "1.Make Utils for reading log fine (fileReader) STATUS: COMPLETED"+
	"\n2.Telegram Chat Bot (chatModule)"+
	"\n3.load all timings to info in a file with xml format by using switch statements and selecting required command"+
	"\n4.unload data to local objects and sort with time and process them accordingly" +
	"\nFlow of Code"  + 
	"\n1.unload and sort data from xml or json"+
	"\n2.go one by one and when time reached get log data and send message to user telegram"+
	"\nwait for user and once message reached verify yes or no"+
	"\nif yes,update xl file and move to next timer and process else no means it is also expected for wait time with no message"+
	"\nso code will sleep that many min"
)
func main() {
	fmt.Println(message)

	fmt.Println("==============Test=====================")

	logInfo,err := logreader.GetLogData("config")
	if err!=nil{
		panic(err)
	}

	fmt.Println(logInfo)
}
