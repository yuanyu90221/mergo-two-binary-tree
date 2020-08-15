# merge-two-binary-tree

## 題目解讀：

### 題目來源:
[merge-two-binary-trees](https://leetcode.com/problems/merge-two-binary-trees/)

### 原文:
Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

### 解讀：

給定兩個二元樹 Tree1, Tree2

把Tree1跟Tree2 做merge

merge 規則如下:

1 假設對應的節點 兩個原本二元樹都有值 則把兩個值相加作為新的節點的值

2 假設原本只有一個二元數有值 則取那個有值得節點的值 作為新節點的值


## 初步解法:
### 初步觀察:
同步兩個二元樹做遍歷

遇到兩個樹都有值則設定新的值為兩個樹原本的值相加

遇到其中一個有值 把有值得那個設定為新的值

### 初步設計:

```code
func transerval(t1 *TreeNode, t2 *TreeNode, result *TreeNode) {
	if t1 == nil && t2 == nil {
		return
	}
	switch {
	case t1 == nil:
		result.Val = t2.Val
		if t2.Left != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Left = &valNode
			transerval(nil, t2.Left, result.Left)
		}
		if t2.Right != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Right = &valNode
			transerval(nil, t2.Right, result.Right)
		}
	case t2 == nil:
		result.Val = t1.Val
		if t1.Left != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Left = &valNode
			transerval(t1.Left, nil, result.Left)
		}
		if t1.Right != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Right = &valNode
			transerval(t1.Right, nil, result.Right)
		}
	default:
		result.Val = t1.Val + t2.Val
		if t1.Left != nil || t2.Left != nil {
			setUpNextNode(t1, t2, result, true)
		}
		if t1.Right != nil || t2.Right != nil {
			setUpNextNode(t1, t2, result, false)
		}
	}
}

func setUpNextNode(t1 *TreeNode, t2 *TreeNode, result *TreeNode, isLeft bool) {
	valNode := TreeNode{Val: 0, Left: nil, Right: nil}
	if isLeft {
		result.Left = &valNode
		transerval(t1.Left, t2.Left, result.Left)
	} else {
		result.Right = &valNode
		transerval(t1.Right, t2.Right, result.Right)
	}
}
```

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package merge_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var result *TreeNode
	if t1 == nil && t2 == nil {
		return nil
	}
	valNode := TreeNode{Val: 0, Left: nil, Right: nil}
	result = &valNode
	transerval(t1, t2, result)
	return result
}

func transerval(t1 *TreeNode, t2 *TreeNode, result *TreeNode) {
	if t1 == nil && t2 == nil {
		return
	}
	switch {
	case t1 == nil:
		result.Val = t2.Val
		if t2.Left != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Left = &valNode
			transerval(nil, t2.Left, result.Left)
		}
		if t2.Right != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Right = &valNode
			transerval(nil, t2.Right, result.Right)
		}
	case t2 == nil:
		result.Val = t1.Val
		if t1.Left != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Left = &valNode
			transerval(t1.Left, nil, result.Left)
		}
		if t1.Right != nil {
			valNode := TreeNode{Val: 0, Left: nil, Right: nil}
			result.Right = &valNode
			transerval(t1.Right, nil, result.Right)
		}
	default:
		result.Val = t1.Val + t2.Val
		if t1.Left != nil || t2.Left != nil {
			setUpNextNode(t1, t2, result, true)
		}
		if t1.Right != nil || t2.Right != nil {
			setUpNextNode(t1, t2, result, false)
		}
	}
}

func setUpNextNode(t1 *TreeNode, t2 *TreeNode, result *TreeNode, isLeft bool) {
	valNode := TreeNode{Val: 0, Left: nil, Right: nil}
	if isLeft {
		result.Left = &valNode
		transerval(t1.Left, t2.Left, result.Left)
	} else {
		result.Right = &valNode
		transerval(t1.Right, t2.Right, result.Right)
	}
}

```
## 測資的撰寫
```golang
package merge_trees

import (
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example1",
			args: args{
				t1: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				t2: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  1,
						Left: nil,
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:  5,
					Left: nil,
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTrees(tt.args.t1, tt.args.t2)
			gotVal := DumpFromTree(got)
			wantVal := DumpFromTree(tt.want)
			if !reflect.DeepEqual(gotVal, wantVal) {
				t.Errorf("mergeTrees() = %v, want %v", gotVal, wantVal)
			}
		})
	}
}

func DumpFromTree(t *TreeNode) []int {
	if t == nil {
		return []int{}
	} else {
		current := []int{}
		current = append(DumpFromTree(t.Left), t.Val)
		return append(current, DumpFromTree(t.Right)...)
	}
}

```

## my self record
[golang leetcode 30 day 17thday](https://hackmd.io/bbR2i-F_QPORfQQ0knQ-jg?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)