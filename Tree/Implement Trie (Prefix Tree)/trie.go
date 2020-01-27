type Node struct{
    childrenMap map[string]*Node
    endOfString bool
}

type Trie struct {
    root *Node
}

func newNode()*Node{
    node := new(Node)
    node.childrenMap = make(map[string]*Node)
    node.endOfString = false
    return node
}

/** Initialize your data structure here. */
func Constructor() Trie {
    node := newNode()
    return Trie{root : node} 
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    currNode := this.root
    for _, val := range word{
        strVal := string(val)
        if nxtNode,ok := currNode.childrenMap[strVal];ok{
            currNode = nxtNode            
        }else {
            node := newNode()
            currNode.childrenMap[strVal] = node
            currNode = node
        }
    }
    currNode.endOfString = true    
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    currNode := this.root
    for _, val := range word{
        strVal := string(val)
        if nxtNode,ok := currNode.childrenMap[strVal];ok{
            currNode = nxtNode            
        }else {
            return false
        }
    }
    if currNode.endOfString{
        return true
    }
    return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    currNode := this.root
    for _, val := range prefix{
        strVal := string(val)
        if nxtNode,ok := currNode.childrenMap[strVal];ok{
            currNode = nxtNode            
        }else {
            return false
        }
    }
    return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */