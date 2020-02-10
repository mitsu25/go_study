
package main

import (
	"fmt"
	"math/rand"
)


func main(){
	piggy := 0
	target_amount := 20 * 100

	for {
	
		switch number := rand.Intn(3); number {
			case 0:
				piggy = piggy + 5
			case 1:
				piggy = piggy + 10
			case 2:
				piggy = piggy + 20	
		}
		fmt.Printf("piggy = $%v.%-2v\n", piggy/100, piggy%100)
		if(piggy >= target_amount){
			break
		}	
	}	
}
