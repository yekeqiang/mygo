package main

var vname1,vname2,vname3 type

var variableName type = value

var vname1,vname2,vname3 type = v1,v2,v3


vname1,vname2,vname3 := v1,v2,v3
//下划线是个特殊的变量名，任何赋予它的值都会被丢弃，在这个例子中，我们将值35赋予b,并同时丢弃34
_,b = 34,35

const constantName = value
const Pi float32 = 3.1415926
const Pi = 3.1415926
const i = 10000
const MaxThread = 10
const prefix = "yekeqiang_"
var s string = "hello"
s[0] = 'c'
s := "hello"
c := []byte(s) //将字符串s转换成[]byte类型
c[0] = 'c'
s2 := string(c)
fmt.Printf("%s\n",s2)


func myfunc() {
	i := 0
	Here:
	println(i)
	i++
	goto Here //跳转到Here去



if integer == 3 {
		fmt.Println("Ther integer is equal to 3")
	}	else if integer < 3 {
		fmt.Println("Ther integer is less than 3")
	}   else {
		fmt.Println("The integer is greater than 3")
	}
}