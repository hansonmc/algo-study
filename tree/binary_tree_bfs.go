package tree

import (
  "container/list"
  "fmt"
)

/*
层序遍历 本质上属于广度优先遍历,也称广度优先搜索(breadth-first search, BFS),
它体现了一种“一圈一圈向外扩展”的逐层遍历方式。
*/
func LevelOrder(root *TreeNode) []any {
  queue := list.New()
  queue.PushBack(root)
  nums := make([]any, 0)
  for queue.Len() > 0 {
    node := queue.Remove(queue.Front()).(*TreeNode)
    nums = append(nums, node.Val)
    if node.Left != nil {
      queue.PushBack(node.Left)
    }
    if node.Right != nil {
      queue.PushBack(node.Right)
    }
    fmt.Println(queue.Len())
  }
  return nums
}
