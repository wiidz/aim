package main

import (
	goutil "github.com/wiidz/goutil/helpers"
	"log"
)

func main() {
	str := "1,2,3,4"
	var typeHelper goutil.TypeHelper //
	temp := typeHelper.ExplodeInt(str, ",")
	log.Println("temp", temp)
}