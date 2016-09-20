package Libs

import (
	"time"
	"math/rand"
	"math"
)

type Number struct {

}

/*
 *随机整数
 */
func (this *Number)rand(start int64, end int64)int64{
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Int63n(end - start) + start
}

func (this *Number)Max(a interface{} , args ...interface{})(interface{}, Error){

	if !this.isInt(a){
		return 0,error("param is not an number")
	}
	if len(args) > 0{
		for _,v:=range args {
			if(!this.isInt(v)){
				return 0,error("param is not an number")
			}
		}
	}else {
		return a , nil
	}
	items :=args[:]
	items = append(items,a)


	return 0,nil
}

/*
 *判断是否为Number
 */
func (this *Number)isNumber(a interface{}) bool{
	if(this.isInt(a)){
		return true
	}
	switch a.(type) {
		case float32 :
			return true
		case float64 :
			return true


		default:
			return false
	}
}
/*
 *判断是否为整型
 */
func (this *Number)isInt(a interface{}) bool{
	switch a.(type) {
		case int8:
			return true
		case int16:
			return true
		case int32:
			return true
		case int:
			return true
		case int64 :
			return true

		default:
			return false
	}
}