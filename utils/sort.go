package utils

// 冒泡排序，从大到小排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[i]
			}
		}
	}
}

// 选择排序，从大到小排序
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		maxVal := arr[i]
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			if maxVal < arr[j] {
				maxVal = arr[j]
				maxIndex = j
			}
		}
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
}

// 插入排序，从大到小排序
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && insertVal > arr[insertIndex] {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}
}

// 归并排序，从大到小排序
func MergeSort(arr []int, left, right int) {
	if right-left <= 1 {
		return
	}
	middle := (left + right) / 2
	MergeSort(arr, left, middle)
	MergeSort(arr, middle, right)

	arrLeft := make([]int, middle-left)
	arrRight := make([]int, right-middle)
	copy(arrLeft, arr[left:middle])
	copy(arrRight, arr[middle:right])

	i := 0
	j := 0
	for k := left; k < right; k++ {
		if i >= middle-left {
			arr[k] = arrRight[j]
			j++
		} else if j >= right-middle {
			arr[k] = arrLeft[i]
			i++
		} else if arrLeft[i] > arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}
}

// 快速排序，从大到小排序
func QuickSort(arr []int, left, right int) {
	l := left
	r := right
	pivot := arr[(left+right)/2]

	for l < r {
		for arr[l] > pivot {
			l++
		}
		for arr[r] < pivot {
			r--
		}

		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]

		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}

	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(arr, left, r)
	}
	if right > l {
		QuickSort(arr, l, right)
	}
}
