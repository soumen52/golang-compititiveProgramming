var (
    rowLength int
    colLength int
)

func setZeroes(matrix [][]int)  {
    rowLength = len(matrix)
    if rowLength > 0{
       colLength = len(matrix[0]) 
    }
    rowArray := make([]bool, rowLength)
    colArray := make([]bool, colLength)
    
    for row := 0; row < rowLength; row++{
        for col := 0; col < colLength; col++{
            if matrix[row][col]==0{
               rowArray[row]=true
               colArray[col]=true
            }
        }        
    }
    for i, val := range rowArray{
        if val{
           setRowzero(matrix, i) 
        }
    }
    
    for i, val := range colArray{
        if val{
           setColzero(matrix, i) 
        }
    }
    
}

func setRowzero(matrix [][]int, row int){
    for col := 0; col < colLength; col++{
        matrix[row][col]=0        
    }
}

func setColzero(matrix [][]int, col int){
    for row := 0; row < rowLength; row++{
        matrix[row][col]=0
    }
}