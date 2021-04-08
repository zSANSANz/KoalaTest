package main 

import "fmt"

func reverse(s string) string { 
    rns := []rune(s)
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 { 
  
        rns[i], rns[j] = rns[j], rns[i] 
    } 
  
    return string(rns) 
} 
  
func main() { 
	var str string
	fmt.Scanf("%s", &str)
    
    strRev := reverse(str) 
	
	if str == strRev {
		fmt.Println("Paliondrome")
	} else {
		fmt.Println("Bukan Paliondrome")
	}
} 