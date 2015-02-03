package main

import(
	"fmt"
	"time"
	"math/rand"
)
	
func fork(left, right <-chan int){
	for{
		select{
			case <- left:
				<- left
			case <- right:
				<- right
		}
	}
}

func philo(name string, leftFork, rightFork chan int, goofy bool){
	eatTime := time.Duration(100)
	thinkTime := time.Duration(200)
	randFactor := 300
	timeFactor := time.Millisecond
	eatCount := 0
	
	for{
		randTime := time.Duration(0)
		if randFactor != 0{
			randTime = time.Duration(rand.Intn(randFactor))
		}
		
		fmt.Println(name," is thinking...")
		time.Sleep((thinkTime + randTime) * timeFactor)	//Thinking
		fmt.Println(name," is hungry...")
		
		if goofy{
				leftFork <- 1	//Pick up left fork
				fmt.Println(name," picked up left fork")
				rightFork <- 1	//Pick up right fork
				
				fmt.Println(name," started eating...")
				time.Sleep((eatTime + randTime) * timeFactor)	//Eating
				fmt.Println(name," finished eating...")
				
				leftFork <- 0	//Put down left fork
				rightFork <- 0	//Put down right fork
		}else{
				rightFork <- 1	//Pick up right fork
				fmt.Println(name," picked up right fork")
				leftFork <- 1	//Pick up left fork
				
				fmt.Println(name," started eating...")
				time.Sleep((eatTime + randTime) * timeFactor)	//Eating
				fmt.Println(name," finished eating...")
				
				leftFork <- 0	//Put down left fork
				rightFork <- 0	//Put down right fork
		}
		eatCount ++
		fmt.Println(name," has eaten ", eatCount, " times now <------------")
	}
	
	
	
	
}

func main(){
	N := 5;
	lastGoofy := true	//If true makes the last philosopher left handed
	
	nameSlice := []string{"Rafael", "Aristotiles", "Sokrates", "Kant", "Hume"}
	
	chSliceLeft := make([]chan int, N)
	chSliceRight := make([]chan int, N)
	for i := range chSliceLeft{
		chSliceLeft[i] = make(chan int)
		chSliceRight[i] = make(chan int)
	}
	
	
	for i := range chSliceLeft{
		go fork(chSliceLeft[i], chSliceRight[i])
		if i+1 == N{	//Wraparound when out of range
			go philo(nameSlice[i], chSliceRight[i], chSliceLeft[0], lastGoofy)
		}else{
			go philo(nameSlice[i], chSliceRight[i], chSliceLeft[i+1], false)
		}
	}
	
	for{
		time.Sleep(time.Second)
	}
}
