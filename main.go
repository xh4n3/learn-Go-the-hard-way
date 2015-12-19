package main

import "reflect"

//Reverse reverses a slice.
func Reverse(slice interface{}) {
	sli := reflect.ValueOf(slice).Elem()
	var temp int64
	var j int
	for i := 0 ; i < sli.Len() / 2; i += 1 {
		j = sli.Len() - i - 1
		temp = sli.Index(i).Int()
		sli.Index(i).SetInt(sli.Index(j).Int())
		sli.Index(j).SetInt(temp)
	}
}

func main() {
	println("Please edit main.go,and complete the 'Reverse' function to pass the test.\nYou should use reflect package to reflect the slice type and make it applly to any type.\nTo run test,please run 'go test'\nIf you pass the test,please run 'git checkout l2' ")
}
