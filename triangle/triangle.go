// Allows to describe a triangle
package triangle

import(
    "math"
)

type Kind string

const (
    NaT = "NaT" // not a triangle
    Equ = "Equ" // equilateral
    Iso = "Iso" // isosceles
    Sca = "Sca" // scalene
)

// KindFromSides gets 3 sides and returns the type of triangle 
func KindFromSides(a, b, c float64) Kind {
        if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
           return "NaT" 
        }
        if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
           return "NaT" 
        }
        if (a <= 0 || b <= 0 || c <= 0) {
            return "NaT"
        }
        if ( a + b < c || b + c < a || c + a < b) {
            return "NaT"
        }
        if a != b && b != c && c != a {
            return "Sca"
        }
        if a == b && a == c && b == c {
            return "Equ"
        }
        if (a == b && a != c) ||  (b == c && b != a) || (a == c && a != b) {  
            return "Iso"   
        }   
        return "NaT"
}
