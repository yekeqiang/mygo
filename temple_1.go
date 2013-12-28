package main

import (
	"log"
	"os"
	"text/template"
	// "fmt"
)

func tmpl(text string, data interface{}) {
	t, _ := template.New("test").Parse(text)
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	t := `
	      Data:
	          {{.}}
	          {{$}}
	      Var、Key:
	          {{/* 注意变量使用 := 赋值 */}}
	          {{$name := .name}} Data.Name = {{$name}}
	      Range:
	          Map: {{range $key,$value :=$}} {{$key}} = {{$value}};{{end}}
              
              {{/* 此处的.代表当前上下文，也就算item*/}}
              Slice:{{range .ints}} {{.}},{{end}}

              {{/* 索引序号 */}}
              Slice: {{range $index,$value := .ints}} [{{$index}}] = {{$value}},{{end}}


              {{/*用内置函数index 访问指定序号的元素*/}}
              Index:ints[2] = {{index .ints 2}}

          With:
              {{/*with 定义块上下文，可以减少前缀输入*/}}
              {{with .meta}} {{.X}} {{end}}

          IF:
              {{/*  0、false、nil、len=0的string、array、slice、map，或者不存在的成员都市False*/}}    
              {{if .ints}} Y {{else}} N {{end}}
              {{with .SomeOne}} Y {{else}} N {{end}}

    `
	d := map[string]interface{}{
		"name": "Q.yuhen",
		"ints": []int{1, 2, 3, 4},
		"meta": struct{ X int }{100},
	}

	tmpl(t, d)

}
