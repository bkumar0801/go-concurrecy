package main
import (
         "fmt"
         "time"
      )

func Greet(msg string){
   for i:=0; ; i++ {
      fmt.Println(msg, i)
      time.Sleep(time.Second)
   }
}

func main() {
   go Greet("Hello,Go Dev!!")
   time.Sleep(2*time.Second)
}
