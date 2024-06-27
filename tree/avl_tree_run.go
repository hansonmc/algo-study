package tree

import (
  "fmt"
)

func RunAVLTree() {
  /* 初始化空 AVL 树 */
	avlTree := NewAVLTree()
	/* 插入节点 */
	// 请关注插入节点后，AVL 树是如何保持平衡的
	runInsert(avlTree, 1)
	runInsert(avlTree, 2)
	runInsert(avlTree, 3)
	runInsert(avlTree, 4)
	runInsert(avlTree, 5)
	runInsert(avlTree, 8)
	runInsert(avlTree, 7)
	runInsert(avlTree, 9)
	runInsert(avlTree, 10)
	runInsert(avlTree, 6)

	/* 插入重复节点 */
	runInsert(avlTree, 7)

	/* 删除节点 */
	// 请关注删除节点后，AVL 树是如何保持平衡的
	runRemove(avlTree, 8) // 删除度为 0 的节点
	runRemove(avlTree, 5) // 删除度为 1 的节点
	runRemove(avlTree, 4) // 删除度为 2 的节点

	/* 查询节点 */
	node := avlTree.Search(7)
	fmt.Printf("\n查找到的节点对象为 %#v ，节点值 = %d \n", node, node.Val)
}

func runInsert(avlTree *AVLTree, val int) {
	avlTree.Insert(val)
	fmt.Printf("\n插入节点 %d 后，AVL 树为 \n", val)
	PrintTree(avlTree.Root)
}

func runRemove(avlTree *AVLTree, val int) {
	avlTree.Remove(val)
	fmt.Printf("\n删除节点 %d 后，AVL 树为 \n", val)
	PrintTree(avlTree.Root)
}
