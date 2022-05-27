# golang_kth_largest_ele_in_array

## 

Given an integer array `nums` and an integer `k`, return *the* $k^{th}$ *largest element in the array*.

Note that it is the $k^{th}$ largest element in the sorted order, not the $k^{th}$ distinct element.

## Examples

**Example 1:**

```
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

```

**Example 2:**

```
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4

```

**Constraints:**

- `1 <= k <= nums.length <= $10^4$`
- $`-10^4$ <= nums[i] <= $10^4$`

## 解析

給定一個整數陣列 nums 還有一個正整數 k

要求寫一個演算法去找到由大到小排序第k順位的數值

**想法一: 使用 MaxHeap**

把陣列放入MaxHeap ⇒ O(nlogn)

pop k次 每次都是 logn ⇒ klogn

第k 次的pop出的數值就是結果

這樣空間複雜度是 O(n)

時間複雜度是 O(nlogn)

**想法二： 使用 MinHeap**

要優化想法一的作法

可以發現，我們只關注較大的k個中最小的那個

也就是可以先放入 k 個陣列中的數值  ⇒ O(klogk)

依次檢查 k+1 到 n 的數值

檢查當下這個數值有沒有大於目前的 k個數值中最小的那個 (是不是前k大)

如果是 ，就把 minHeap 目前 top 數值換成當下這個數值做 heapify ⇒ log(k)

換到最後 pop 出來的數值就是第 k 大的數值

這樣的時間複雜度是 O(klogk)

空間複雜度是 O(k) 

## 程式碼
```go
package sol

import "container/heap"

type MinIntHeap []int

func (h *MinIntHeap) Len() int {
	return len(*h)
}
func (h *MinIntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *MinIntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinIntHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}
func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[1 : n-1]
	return x
}
func findKthLargest(nums []int, k int) int {
	pq := MinIntHeap(nums[:k])
	heap.Init(&pq)
	for _, num := range nums {
		if pq[0] < num {
			pq[0] = num
			// heapify
			heap.Fix(&pq, 0)
		}
	}
	return heap.Pop(&pq).(int)
}

```
## 困難點

1. 理解 Min-Heap 在第 k 個大的作用可以每次只保留最目前 k 個較大 其他 n-k+1 個都不用去管
2. 需要理解 golang  container/heap 中 Fix 代表 heapify

## Solve Point

- [x]  Understand what need to solve
- [x]  Analysis Complexity