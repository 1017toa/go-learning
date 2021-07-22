package main

import (
	"fmt"

	"github.com/1017toa/go-learning/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
	del := dictionary.Delete(baseWord)
	if err != nil {
		fmt.Println(del)
	}
	word2, _ := dictionary.Search(baseWord)
	fmt.Println(word2)

}
