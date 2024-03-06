package mystrings

func Reverse(s string) string {
   
    runes := []rune(s)

    result := ""

    // Iterate over the string in reverse order
    for i := len(runes) - 1; i >= 0; i-- {
        // Append each rune to the result string
        result += string(runes[i])
    }

    // Return the reversed string
    return result
}
