package main
import (
         "fmt"
         "time"
      )

func Greet(msg string, c chan string){
   for i:=0; ; i++ {
      c <- fmt.Sprintf("%s %d", msg, i)
      time.Sleep(2*time.Second)
   }
}

func main() {
   c := make(chan string)
   go Greet("Hi, Go Dev!!", c)
   for i := 0; i< 5; i++ {
      fmt.Printf("Sub-routine Say: %q\n", <-c)
   }
   fmt.Println("You are boring: I am leaving!!")
}
