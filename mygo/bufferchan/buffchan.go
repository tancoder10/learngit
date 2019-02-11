package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numGoroutine = 4
	taskLoad     = 100
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	task := make(chan string, taskLoad)
	wg.Add(numGoroutine)
	for gr := 1; gr <= numGoroutine; gr++ {
		go worker(task, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		task <- fmt.Sprintf("Task:%d", post)
	}
	close(task)

	wg.Wait()
}

func worker(task chan string, worker int) {
	//通知main函数，此函数已经结束并返回结果
	defer wg.Done()
	for {
		//从通道理取数据，如果通道里没数据会一直阻塞
		item, ok := <-task
		if !ok {
			fmt.Printf("worker:%d done\n", worker)
			return
		}
		fmt.Printf("worker:%d deal task %s\n", worker, item)
	}
}
