package main
import (
"fmt"
"runtime"
)
func say(s string) {


fmt.Println(s)
runtime.Gosched()
}

func main() {
say("1") 
for i:=0;i<20;i++{
go say("world")
} //开一个新的Goroutines执行
say("2")
}
