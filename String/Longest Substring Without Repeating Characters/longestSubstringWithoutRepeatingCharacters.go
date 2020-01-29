func lengthOfLongestSubstring(s string) int {
    var ans int
    var i, j int
    strMap := make(map[string]bool)
    for i < len(s) && j < len(s){
        strChar := string(s[j])
        if _,ok := strMap[strChar];ok{
            delete(strMap, string(s[i]))
            i++
        }else{
            j++
            ans = int(math.Max(float64(ans),float64(j-i)))
            strMap[strChar]=true
        }
    }
    return ans
}