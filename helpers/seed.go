package helpers

import (
	"fmt"
	"time"
)

func GenerateSeed() int64 {
	seed := time.Now().UnixNano()
	fmt.Println("/*********************************************/")
	fmt.Println("/**************      SEED     ****************/")
	fmt.Printf("/**********   %d   **********/\n", seed)
	fmt.Println("/**************               ****************/")
	fmt.Println("/*********************************************/")

	return seed
}
