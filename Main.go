package main

import "fmt"

func main() {
   var x float64
   x = 20.0
   fmt.Println(x)
   fmt.Printf("x is of type %T\n", x)

   /* local variable definition */
   var a int = 100
   var b int = 200
   var ret int

   /* calling a function to get max value */
   ret = max(a, b)
   fmt.Printf( "Max value is : %d\n", ret )

   a1, b1 := swap("Mahesh", "Kumar")
   fmt.Println(a1, b1)

   /* nextNumber is now a function with i as 0 */
   nextNumber := getSequence()

   /* invoke nextNumber to increase i by 1 and return the same */
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())

   /* create a new sequence and see the result, i is 0 again*/
   nextNumber1 := getSequence()
   fmt.Println(nextNumber1())
   fmt.Println(nextNumber1())
}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
   /* local variable declaration */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}

func swap(x, y string) (string, string) {
   return y, x
}

func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
      return i
   }
}
