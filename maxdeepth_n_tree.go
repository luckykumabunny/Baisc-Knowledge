/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func maxDepth(root *Node) int {
    if root==nil{
        return 0
    }
    max_deep := 0
    for _, node := range root.Children {
        max_deep = Max(maxDepth(node), max_deep)
    }
    return 1+max_deep
}

func Max(a int, b int) int{
    if a>b{
        return a
    }   
    return b
}