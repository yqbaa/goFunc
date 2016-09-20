package Libs

import (
	"math/rand"
	"time"

)

type Str struct {
}

func (this *Str) Rand(len int8) interface{} {
	if len <= 0 {
		return ""
	}
	var str string = "0123456789qwertyuioplkjhgfdsazxcvbnm"
	s := []byte(str)
	var i int8 = 0
	r := rand.New(rand.NewSource( int64(time.Now().Nanosecond()) ))
	var newstr string
	for i = 0; i < len; i++ {
		newstr += string(s[r.Intn(35)])

	}
	return newstr

}
