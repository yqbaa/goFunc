package Libs

import (
	"math/rand"
	"time"
)

type Number struct {

}

/*
 *随机整数
 */
func (this *Number) rand(start int64, end int64) int64 {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Int63n(end-start) + start
}

func (this *Number) Max(a int, args ...int) (interface{}) {
	if !this.isInt(a) || !this.isNumber(a) {
		return false
	}
	var argsLen int = len(args)
	if argsLen > 0 {
		if this.isInt(a) && !this.isInt(args[0]) {
			return false
		}
		for _, v := range args {
			if !this.isInt(v) {
				return false
			}
		}
	} else {
		return a
	}
	items := args[:]
	items = append(items, a)
	this.quickSort(items, 0, argsLen)
	return items[argsLen]
}

/*
 * 取最小值
 */
func (this *Number) Min(a int, args ...int) (interface{}) {
	if !this.isInt(a) || !this.isNumber(a) {
		return false
	}
	var argsLen int = len(args)
	if argsLen > 0 {
		if this.isInt(a) && !this.isInt(args[0]) {
			return false
		}
		for _, v := range args {
			if !this.isInt(v) {
				return false
			}
		}
	} else {
		return a
	}
	items := args[:]
	items = append(items, a)
	this.quickSort(items, 0, argsLen)
	return items[0]
}

/*
 *判断是否为Number
 */
func (this *Number) isNumber(a interface{}) bool {
	if this.isInt(a) {
		return true
	}
	switch a.(type) {
	case float32:
		return true
	case float64:
		return true

	default:
		return false
	}
}

/*
 *判断是否为整型
 */
func (this *Number) isInt(a interface{}) bool {
	switch a.(type) {
	case int8:
		return true
	case int16:
		return true
	case int32:
		return true
	case int:
		return true
	case int64:
		return true

	default:
		return false
	}
}

/*
 * 快排
 */
func (this *Number) quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			this.quickSort(arr, start, j)
		}
		if end > i {
			this.quickSort(arr, i, end)
		}
	}

}
