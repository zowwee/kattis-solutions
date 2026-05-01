package main

import "fmt"

func main(){
    var x int
    fmt.Scan(&x)

    sum := 0.0

    for i := 0; i < x; i ++{

        var q, y float64
        fmt.Scan(&q, &y)

        b:= q * y

        sum += b
    }
    fmt.Println(sum)
}
