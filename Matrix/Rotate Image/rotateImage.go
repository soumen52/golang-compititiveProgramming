func rotate(matrix [][]int)  {
    if len(matrix)==0 || len(matrix) != len(matrix[0]){
        return 
    }
    for layer := 0; layer < len(matrix)/2; layer++{
        first := layer 
        last := len(matrix)-1-layer
        for i := first; i < last ; i++{
            offset := i-first
            
            //save the top
            tmp := matrix[first][i]
            
            //left - > top
            matrix[first][i] = matrix[last-offset][first]
            
            //right -> left 
            matrix[last-offset][first] = matrix[last][last-offset]
            
            //right -> bottom
            matrix[last][last-offset] = matrix[i][last]
            
            //top -> right
            matrix[i][last] = tmp
        }
    }
}