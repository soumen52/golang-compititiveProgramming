//time Complexity : O(n)
//Space Complexity :  O(1)
func isPalindrome(s string) bool {
    i := 0
    j := len(s)-1
    for i <= j{
        if !isAlphaNumeric(s[i]){
            i++
            continue
        }
        if !isAlphaNumeric(s[j]){
            j--
            continue
        }
        if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])){
            return false
        }
        i++
        j--
    }
    return true
}

func isAlphaNumeric(s byte)bool{
    if (s >= 65 && s <= 90) || (s >= 97 && s <= 122) || (s >= 48 && s <= 57){
        return true
    }
    return false
}