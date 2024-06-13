package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var userMap = map[int]int{
	1:  1,
	2:  2,
	3:  3,
	4:  4,
	5:  5,
	6:  6,
	7:  7,
	8:  8,
	9:  9,
	10: 10,
	40: 40,
	50: 50,
	70: 70,
}

func FindUserByID(userID int) int {
	val := userMap[userID]
	return val
}

// This function simulates checking a user ID with a 1-second delay
func checkUser(workerID int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer func(wID int) {
		fmt.Println("EXIT GOROUTINE: ", wID)
		wg.Done()
	}(workerID)

	for id := range jobs {
		fmt.Printf("workerID: %d, check userID: %v\n", workerID, id)
		userID := FindUserByID(id)

		if userID == 0 {
			fmt.Printf("userID: %v not found\n", id)
			continue
		}

		time.Sleep(1000 * time.Millisecond)

		fmt.Printf("userID: %v found\n", userID)
		results <- userID
	}
}

func main() {
	r := gin.Default()
	fmt.Println("Init: ", runtime.NumGoroutine())

	// Route for asynchronous user ID check with worker pool
	r.POST("/checkUserIDs", func(c *gin.Context) {
		fmt.Println("==================== NEW REQUEST [Async] ====================")
		fmt.Println("Start Request: ", runtime.NumGoroutine())

		// Define struct for request body
		var request struct {
			UserIDs []int `json:"user_ids"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create wait group and channels
		var wg sync.WaitGroup
		jobs := make(chan int, len(request.UserIDs))
		results := make(chan int, len(request.UserIDs))

		// Spawn 10 workers
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go checkUser(i, jobs, results, &wg)
		}

		// Send user IDs to the jobs channel
		// The idle worker will then retrieve the data and process it.
		for _, userID := range request.UserIDs {
			jobs <- userID
		}
		close(jobs) // Close jobs channel after all data is sent

		// Wait for all goroutines to finish
		wg.Wait()
		close(results) // Close results channel to ensure proper closing

		fmt.Println("After Close Result Channel: ", runtime.NumGoroutine())

		// Collect found user IDs
		var foundUserIDs []int
		for res := range results {
			foundUserIDs = append(foundUserIDs, res)
		}

		c.JSON(http.StatusOK, gin.H{"found_user_ids": foundUserIDs})
	})

	r.POST("/checkUserIDsWithoutWorkerPool", func(ctx *gin.Context) {
		fmt.Println("==================== NEW REQUEST [Sync] ====================")
		var request struct {
			UserIDs []int `json:"user_ids"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var foundUserIDs []int

		for _, userID := range request.UserIDs {
			fmt.Printf("check userID: %v\n", userID)
			gotID := FindUserByID(userID)

			if gotID == 0 {
				fmt.Printf("userID: %v not found\n", userID)
				continue
			}
			time.Sleep(1000 * time.Millisecond)

			fmt.Printf("userID: %v found\n", gotID)
			foundUserIDs = append(foundUserIDs, gotID)
		}
		ctx.JSON(http.StatusOK, gin.H{"found_user_ids": foundUserIDs})
	})
	r.Run()
}
