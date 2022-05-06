package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`  // 添加两个标签，标签需要是键值对
	Sex string `info:"sex"` // 添加一个标签
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", tagInfo, "doc: ", tagDoc)
	}
}

func main() {
	var r resume
	findTag(&r)
}