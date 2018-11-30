package main;
import "fmt"
import "math/rand"

func main() {
  var a int;
  a = 5;
  var b = 5;
  c := 5;
  fmt.Printf("Value of variable a is %d \n", a);
  fmt.Printf("Value of variable b is %d \n", b);
  fmt.Printf("Value of variable c is %d \n", c);
  rand.Seed(2);
  fmt.Println("Random Value is: ", rand.Intn(10));
}
