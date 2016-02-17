package main

import (
"fmt"
_"math/rand"
)
import (
_"github.com/user/stringutil"
_ "io/ioutil"
_ "strings"
_ "os"
_ "time"
	"runtime"
)

//생산자
func producer(c chan<- int){
for i :=0; i < 5; i++ {
c <- i
}

c <-100

}

//소비자
func consumer (c <-chan int){
for i := range c {
fmt.Println(i)
}
}


//메인함수
func main() {

	runtime.GOMAXPROCS(8)

	c := make(chan int)

	go consumer(c)
	go producer(c)

	fmt.Scanln()


}

