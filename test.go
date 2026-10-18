package main

import (
	"fmt"
	"time"
	"sync"
)

func fetchPage(url string){
	fmt.Println("Fetching: " + url)
	time.Sleep(1 * time.Second)
	fmt.Println("Finished: " + url)
}

func sayHello(wg *sync.WaitGroup) {
    defer wg.Done()

    time.Sleep(100 * time.Millisecond)
    fmt.Println("Hello!")
}



// func main(){
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go sayHello(&wg)
// 	go sayHello(&wg)

// 	wg.Wait()
// 	fmt.Println("testing")
// }