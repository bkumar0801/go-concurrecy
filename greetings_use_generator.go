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
   Adam := Greet("Hi, Eve!!")
   Eve  := Greet("Hi, Adam!!")
   for i := 0; i< 5; i++ {
      fmt.Println(<-Adam)
      fmt.Println(<-Eve)
   }
   fmt.Println("You are boring: I am leaving!!")
}
