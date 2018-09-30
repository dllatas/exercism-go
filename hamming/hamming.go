package hamming

import "errors"

func Distance(a, b string) (int, error) {
    var diff int = 0
    if len(a) != len(b) {
        return -1, errors.New("Mismatch between strings")     
    }
    for i := 0; i < len(a); i++ {
       if a[i] != b[i] {
           diff = diff + 1
       } 
    }
    return diff, nil
}
