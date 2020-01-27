
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 

func kthSmallest(root *TreeNode, k int) int {
    var elem int
    sortedNodes := indorderTraversal(root, []int{})
    if len(sortedNodes) > 0 {
        elem = sortedNodes[k-1] 
    }
    return elem
}


func indorderTraversal(node *TreeNode, list []int)[]int{
    if node==nil{
        return list
    }
    list = indorderTraversal(node.Left,list)
    list = append(list, node.Val)
    list = indorderTraversal(node.Right,list)
    return list
}