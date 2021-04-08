package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Scanf("%d", &bilangan)

    for y:=1; y<=bilangan; y++ {
    	a := 0;
        for g:=1; g<=y; g++ {
            if y % g == 0 {
            a++;
        	}
    	}

		if a == 2 {
			if bilangan == y {
				fmt.Println("Bilangan Prima")
			}
		}
    }

}