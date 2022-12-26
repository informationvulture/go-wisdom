package concurrency

import (
	"fmt"
	"time"
)

// Testing concurrency, capital for export.
func TicketSender() {
	time.Sleep(5 * time.Second)
	fmt.Println("#################")
	fmt.Printf("SENDING YOUR EMAIL at %v\n", time.Now())
	fmt.Println("#################")
}
