package raindrops

import(
        "fmt"
)

func Convert(number int) string {
    var convert string = "" 
    if number % 3  == 0 {
        convert += "Pling"
    }
    if number % 5 == 0 {
        convert += "Plang"
    }
    if number % 7 == 0 {
        convert += "Plong"
    }
    if convert == "" { 
        return fmt.Sprintf("%d", number);
    }
    return convert
}
