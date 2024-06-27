package tree

/* AVL树 */
type AVLTree struct {
  Root *TreeNode
}

func NewAVLTree() *AVLTree {
  return &AVLTree{Root: nil}
}

/* 获取节点高度 */
func (t *AVLTree) height(node *TreeNode) int {
  if node != nil {
    return node.Height
  }
  return -1
}

/* 更新节点高度 */
func (t *AVLTree) updateHeight(node *TreeNode) {
  lh := t.height(node.Left)
  rh := t.height(node.Right)
  if lh > rh {
    node.Height = lh + 1
  } else {
    node.Height = rh + 1
  }
}

/* 获取平衡因子 */
func (t *AVLTree) balanceFactor(node *TreeNode) int {
  if node == nil {
    return 0
  }
  return t.height(node.Left) - t.height(node.Right)
}

/* 右旋操作 */
func (t *AVLTree) rightRotate(node *TreeNode) *TreeNode {
  child := node.Left
  grandChild := child.Right
  child.Right = node
  node.Left = grandChild
  t.updateHeight(node)
  t.updateHeight(child)
  return child
}

/* 左旋操作 */
func (t *AVLTree) leftRotate(node *TreeNode) *TreeNode {
  child := node.Right
  grandChild := child.Left
  child.Left = node
  node.Right = grandChild
  t.updateHeight(node)
  t.updateHeight(child)
  return child
}

/* 执行旋转操作，使该子树重新恢复平衡 */
func (t *AVLTree) rotate(node *TreeNode) *TreeNode {
  bf := t.balanceFactor(node)
  if bf > 1 {
    if t.balanceFactor(node.Left) >= 0 {
      return t.rightRotate(node)
    } else {
      t.leftRotate(node.Left)
      return t.rightRotate(node)
    }
  }
  if bf < -1 {
    if t.balanceFactor(node.Right) <= 0 {
      return t.leftRotate(node)
    } else {
      node.Right = t.rightRotate(node.Right)
      return t.leftRotate(node)
    }
  }
  return node
}

/* 插入节点 */
func (t *AVLTree) Insert(val int) {
  t.Root = t.insertHelper(t.Root, val)
}

/* 递归插入节点（辅助函数） */
func (t *AVLTree) insertHelper(node *TreeNode, val int) *TreeNode {
  if node == nil {
    return NewTreeNode(val)
  }
  if val < node.Val.(int) {
    node.Left = t.insertHelper(node.Left, val)
  } else if val > node.Val.(int) {
    node.Right = t.insertHelper(node.Right, val)
  } else {
    return node
  }
  t.updateHeight(node)
  node = t.rotate(node)
  return node
}

/* 删除节点 */
func (t *AVLTree) Remove(val int) {
  t.Root = t.removeHelper(t.Root, val)
}

func (t *AVLTree) removeHelper(node *TreeNode, val int) *TreeNode {
  if node == nil {
    return nil
  }
  if val < node.Val.(int) {
    node.Left = t.removeHelper(node.Left, val)
  } else if val > node.Val.(int) {
    node.Right = t.removeHelper(node.Right, val)
  } else {
    if node.Left == nil || node.Right == nil {
      child := node.Left
      if node.Right != nil {
        child = node.Right
      }
      if child == nil {
        return nil
      } else {
        node = child
      }
    } else {
      temp := node.Right
      for temp.Left != nil {
        temp = temp.Left
      }
      node.Right = t.removeHelper(node.Right, temp.Val.(int))
      node.Val = temp.Val
    }
  }
  t.updateHeight(node)
  node = t.rotate(node)
  return node
}

/* 查找节点 */
func (t *AVLTree) Search(val int) *TreeNode {
  cur := t.Root
  for cur != nil {
    if cur.Val.(int) < val {
      cur = cur.Right
    } else if cur.Val.(int) > val {
      cur = cur.Left
    } else {
      break
    }
  }
  return cur
}
