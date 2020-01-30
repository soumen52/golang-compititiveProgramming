
//Time Complexity : O(n*klogk) n is length of input string Array . K is length of each string
//Space Complexity: O(n*k) 
func groupAnagrams(strs []string) [][]string {
    groupMap := make(map[string][]string)
    var result [][]string
    for _, str := range strs{
        strArray := strings.Split(str,"")
        sort.Strings(strArray)
        strS := strings.Join(strArray,"")
        groupMap[strS] = append(groupMap[strS],str)
    }
    for _,val := range groupMap{
        result = append(result,val)
    }
    return result
}

//Time Complexity : O(n*k) n is length of input string Array . K is length of each string
//Space Complexity: O(n*k) 
func groupAnagrams(strs []string) [][]string {
    var result [][]string
    if len(strs)==0{
        return result
    }
    resultMap := make(map[string][]string)
    count := make([]string,26)
    for _, str := range strs{
        fillArrayWithZero(count)
        for _, c := range str{
            val := count[c-'a']
            iVal,_ := strconv.Atoi(val)
            count[c-'a']=strconv.Itoa(iVal+1)            
        }
        key :=strings.Join(count,"")
        resultMap[key]=append(resultMap[key], str)        
    }
    for _,val := range resultMap{
        result = append(result, val)
    }
    return result
}

func fillArrayWithZero(count []string){
    for i, _ := range count{
        count[i]="0"
    }
}