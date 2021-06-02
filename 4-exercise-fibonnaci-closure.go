package main

import "fmt"

func fibonacci() func() int {
n:=0
m:=0
return func() int {
num:=n+m
if n==0 {
m=1}
n=m
m=num
return num
}
}
func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
	fmt.Println(f())
    }
}