package tree

// 二叉树
type TreeNode struct {
  Val any
  Height int
  Left *TreeNode
  Right *TreeNode
}

func NewTreeNode(val any) *TreeNode {
  return &TreeNode{
    Val: val,
    Left: nil,
    Right: nil,
  }
}
