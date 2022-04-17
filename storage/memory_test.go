package storage

import (
	"fmt"
	"sync"
	"testing"
)
//проверяем защиту при конкурентом доступе к map
func TestPostStorage(t *testing.T) {
	fmt.Println(len(urls))
	wg := sync.WaitGroup{}
	n := 200
	wg.Add(n)
	for i := 0; i < n; i++ {
		srt := "a"
		orig := "b"
		go func(i int) {
			PostStorage(srt, orig)
			wg.Done()
		}(i)
		srt += string(rune(i))
		orig += string(rune(i))
	}
	wg.Wait()
	if len(urls) != 200 {
		t.Fatalf("Want %v, but got %v", 200, len(urls))
	}
}
//в словаре уже хранится 200 элементов от TestPostStorage
func TestGetStorage(t *testing.T) {
	wg := sync.WaitGroup{}
	n := 200
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			GetStorage("a")
			wg.Done()
		}(i)
	}
	wg.Wait()
}

