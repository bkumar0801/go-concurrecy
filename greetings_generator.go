package main
import (
         "fmt"
         "time"
      )

func Greet(msg string) <- chan string {
   c := make(chan string)
   go func() {
      for i:=0; ; i++ {
      	c <- fmt.Sprintf("%s %d", msg, i)
      	time.Sleep(2*time.Second)
      }
   }()
  return c
}

func main() {
   c := Greet("Hi, Go Dev!!")
   for i := 0; i< 5; i++ {
      fmt.Printf("Sub-routine Say: %q\n", <-c)
   }
   fmt.Println("You are boring: I am leaving!!")
}
