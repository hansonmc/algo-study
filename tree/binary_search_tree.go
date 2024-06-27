package tree

import (
  "fmt"
)

type binarySearchTree struct {
  root *TreeNode
}

func newBinarySearchTree() *binarySearchTree {
  bst := &binarySearchTree{}
  bst.root = nil
  return bst
}

/* 获取根节点 */
func (bst *binarySearchTree) getRoot() *TreeNode {
  return bst.root
}

/* 查找节点 */
func (bst *binarySearchTree) search(num int) *TreeNode {
  node := bst.root
  for node != nil {
    if node.Val.(int) < num {
      node = node.Right
    } else if node.Val.(int) > num {
      node = node.Left
    } else {
      break
    }
  }
  return node
}

/* 插入节点 */
func (bst *binarySearchTree) insert(num int) {
  cur := bst.root
	if cur == nil {
		bst.root = NewTreeNode(num)
		return
	}
	var pre *TreeNode = nil
	for cur != nil {
		if cur.Val.(int) == num {
			return
		}
		pre = cur
		if cur.Val.(int) < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	node := NewTreeNode(num)
	if pre.Val.(int) < num {
		pre.Right = node
	} else {
		pre.Left = node
	}
}

/*删除节点*/
func (bst *binarySearchTree) remove(num int) {
 cur := bst.root
	// 若树为空，直接提前返回
	if cur == nil {
		return
	}
	// 待删除节点之前的节点位置
	var pre *TreeNode = nil
	// 循环查找，越过叶节点后跳出
	for cur != nil {
		if cur.Val == num {
			break
		}
		pre = cur
		if cur.Val.(int) < num {
			// 待删除节点在右子树中
			cur = cur.Right
		} else {
			// 待删除节点在左子树中
			cur = cur.Left
		}
	}
	// 若无待删除节点，则直接返回
	if cur == nil {
		return
	}
	// 子节点数为 0 或 1
	if cur.Left == nil || cur.Right == nil {
		var child *TreeNode = nil
		// 取出待删除节点的子节点
		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}
		// 删除节点 cur
		if cur != bst.root {
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		} else {
			// 若删除节点为根节点，则重新指定根节点
			bst.root = child
		}
		// 子节点数为 2
	} else {
		// 获取中序遍历中待删除节点 cur 的下一个节点
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		// 递归删除节点 tmp
		bst.remove(tmp.Val.(int))
		// 用 tmp 覆盖 cur
		cur.Val = tmp.Val
	}  
}

/* 打印二叉搜索树 */
func (bst *binarySearchTree) print() {
	PrintTree(bst.root)
}


func RunBinarySearchTree() {
  bst := newBinarySearchTree()
	nums := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	// 请注意，不同的插入顺序会生成不同的二叉树，该序列可以生成一个完美二叉树
	for _, num := range nums {
		bst.insert(num)
	}
	fmt.Println("\n初始化的二叉树为:")
	bst.print()
	// 获取根节点
	node := bst.getRoot()
	fmt.Println("\n二叉树的根节点为:", node.Val)
	// 查找节点
	node = bst.search(7)
	fmt.Println("查找到的节点对象为", node, "，节点值 =", node.Val)
	// 插入节点
	bst.insert(16)
	fmt.Println("\n插入节点后 16 的二叉树为:")
	bst.print()
	// 删除节点
	bst.remove(1)
	fmt.Println("\n删除节点 1 后的二叉树为:")
	bst.print()
	bst.remove(2)
	fmt.Println("\n删除节点 2 后的二叉树为:")
	bst.print()
	bst.remove(4)
	fmt.Println("\n删除节点 4 后的二叉树为:")
	bst.print()
}

