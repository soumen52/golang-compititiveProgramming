func isAnagram(s string, t string) bool {
    sourceMap := make(map[string]int)
    for _, str := range s{
        if val,ok:=sourceMap[string(str)]; ok{
            sourceMap[string(str)]=val+1
            continue
        }
        sourceMap[string(str)]=1
    } 
    
    for _, str := range t{
        if val,ok := sourceMap[string(str)];ok{
            sourceMap[string(str)]=val-1
        }else{
            return false
        }
    }
    
    for _,val := range sourceMap{
        if val > 0{
            return false
        }
    }
    return true
}