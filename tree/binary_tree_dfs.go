package tree

var Nums []any

/*
前序、中序和后序遍历都属于深度优先遍历，也称尝试优先搜索(depth-first search, DFS)，
它体现了一种“先走到尽头，再回溯继续”的遍历方式。
*/
// 前序遍历 根节点 -> 左子树 -> 右子树
func PreOrder(root *TreeNode) {
  if root == nil {
    return
  }
  
  Nums = append(Nums, root.Val)
  PreOrder(root.Left)
  PreOrder(root.Right)
}

// 中序遍历 左子树 -> 根节点 -> 右子树
func InOrder(root *TreeNode) {
  if root == nil {
    return
  }
  
  InOrder(root.Left)
  Nums = append(Nums, root.Val)
  InOrder(root.Right)
}

// 后序遍历 左子树 -> 右子树 -> 根节点
func PostOrder(root *TreeNode) {
  if root == nil {
    return
  }
  PostOrder(root.Left)
  PostOrder(root.Right)
  Nums = append(Nums, root.Val)
}
