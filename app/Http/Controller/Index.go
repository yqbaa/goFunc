package Controller

import "Social/Libs"

type IndexController struct {
	BaseController
}

func (this *IndexController)Get(){
	//s :=Libs.Str{}
	//this.success(s.Rand(5))
	n:=Libs.Number{}
	this.success(n.Max(10,2,3,54,6))

}