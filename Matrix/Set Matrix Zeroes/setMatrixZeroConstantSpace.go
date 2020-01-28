var (
   //Declare row length and column length as global so that it can be used from other function 
   rowLength int 
   colLength int
)

func setZeroes(matrix [][]int)  {
    //find the numbers of row
    rowLength = len(matrix)
    if rowLength > 0{
       //find the number of column 
       colLength = len(matrix[0]) 
    }
    
    var ifRowHasZero bool
    var ifColHasZero bool
    
    //check if 1st row has zero
    for col :=0; col < colLength; col++{
        if matrix[0][col]==0{
           ifRowHasZero = true 
            break
        }
    }
    
    
    //check if 1st column has zero
    for row := 0; row < rowLength; row++{
        if matrix[row][0]==0{
            ifColHasZero = true
            break
        }
    }
    
    //set the correding row and column as zero
    for row := 1; row < rowLength; row++{
        for col := 1; col < colLength ;col++{
            if matrix[row][col]==0{
                matrix[row][0]=0
                matrix[0][col]=0
            }
        }
    }
    
    for row := 1; row < rowLength; row++{
        if matrix[row][0] == 0{
            setRowzero(matrix, row)
        } 
    }    
    
    for col := 1; col < colLength; col++{
        if matrix[0][col] == 0{
           setColzero(matrix, col)
        }
    }
    
    if ifRowHasZero {
        setRowzero(matrix, 0)         
    }
    
    if ifColHasZero {
        setColzero(matrix, 0)
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