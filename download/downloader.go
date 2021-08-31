package main

import (
	"fmt"
	"learngo/download/infra"
	"learngo/download/retrieve"
)

func getRetriever() retrieve.Retriever {
	return infra.Retriever{}
}

func main() {
	r := getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
