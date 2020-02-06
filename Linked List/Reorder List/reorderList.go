/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil{
        return 
    }
    var newTail *ListNode
    var secList *ListNode
    slowP := head
    fastP := head.Next.Next
    for fastP != nil{
        slowP = slowP.Next
        if fastP.Next == nil{
           break
        }
        fastP = fastP.Next.Next        
    }
    newTail = slowP
    secList = slowP.Next
    newTail.Next = nil
    secList = reverseList(secList)
    head = mergeList(head, secList)
}

func mergeList(l1, l2 *ListNode) *ListNode{
    if l1 == nil{
        return l1
    }
    if l2 == nil{
        return l1
    }
    head := l1
    for l1 != nil && l2 != nil{
        tmp := l1.Next        
        l1.Next = l2
        l2 = l2.Next
        l1.Next.Next=tmp
        l1=tmp        
    }
    return head
}

func reverseList(head *ListNode) *ListNode{
    first := head
    second := head.Next
    for second != nil{
        tmp := second.Next
        second.Next = first
        first = second
        second = tmp
    }
    head.Next = nil
    return first
}