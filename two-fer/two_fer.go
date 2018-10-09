// Package twofer implements a function for
// conditionally printing a phrase given a name or not
package twofer

// ShareWith checks if parameter name is empty or not and 
// returns a string including that name or "you"
func ShareWith(name string) string {
        if name == "" {
            name = "you"
        }
       return "One for " + name + ", one for me."
}
