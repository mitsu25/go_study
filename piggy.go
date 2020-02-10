
package main

import (
	"fmt"
	"math/rand"
)


func main(){
	piggy := 0.0
	target_amount := 20.0

	for {
	
		switch number := rand.Intn(3); number {
			case 0:
				piggy = piggy + 0.05
			case 1:
				piggy = piggy + 0.1
			case 2:
				piggy = piggy + 0.2	
		}
		fmt.Printf("piggy = %4.2f\n", piggy)
		if(piggy >= target_amount){
			break
		}	
	}	
}
