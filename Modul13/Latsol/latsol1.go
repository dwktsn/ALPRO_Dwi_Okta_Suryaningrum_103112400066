package main

import "fmt"

func main(){
    var bil, digit int

    digit = 0
    fmt.Scan(&bil)

    for bil > 0{
        bil = bil / 10
        digit++
    }
    fmt.Println(digit)
}