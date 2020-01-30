func minWindow(s string, t string) string {
    requiredChars := buildCharacterMap(t)
    charTobeFound := len(t)
    charMatched:=0
    windowCharMap := make(map[string]int)
    var right, left int
    var minWindow string
    
    for right < len(s){
        rightChar := string(s[right])
        currentCount := putInMap(rightChar, windowCharMap)
                      
        if val,ok := requiredChars[rightChar];ok{
            if val == currentCount{
               charMatched+=val
            }
        }
        for charMatched == charTobeFound && left <= right{
            windowLength := right-left+1
            if len(minWindow) >  windowLength{
                minWindow = s[left:right+1]
            }
            if minWindow == ""{
               minWindow = s[left:right+1]
            }
            leftChar := string(s[left])
            currentCount = deleteFromMap(leftChar, windowCharMap)
            
            
            if val,ok := requiredChars[leftChar];ok{
                if currentCount < val{
                    charMatched-=val                    
                }
            }
            left++
        }
        right++
    }
    return minWindow
}

func buildCharacterMap(t string)map[string]int{
    tMap := make(map[string]int)
    for _, char := range t{
        strChar := string(char)
        if val, ok := tMap[strChar]; ok{
            tMap[strChar]=val+1
            continue
        }
        tMap[strChar]=1
    }
    return tMap
}

func putInMap(strChar string, sMap map[string]int)int{
    if val, ok := sMap[strChar]; ok{
        sMap[strChar]=val+1        
    }else{
        sMap[strChar]=1
    }
    return sMap[strChar]
}

func deleteFromMap(strChar string, windowCharMap map[string]int)int{
    if val, ok := windowCharMap[strChar]; ok{
        if val > 0{
            val = val-1
            windowCharMap[strChar]=val
            return val
        }
    }
    return 0
}