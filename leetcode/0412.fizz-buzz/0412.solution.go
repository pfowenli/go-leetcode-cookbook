package leetcode

import "fmt"

func fizzBuzz(n int) []string {
    results := make([]string, n)
    for index := 0; index < n; index++ {
        num := index + 1

        element := ""

        if num % 3 == 0 {
            element = fmt.Sprintf("%sFizz", element)
        }
        if num % 5 == 0 {
            element = fmt.Sprintf("%sBuzz", element)
        }
        
        if len(element) == 0 {
            element = fmt.Sprintf("%d", num)
        }

        results[index] = element
    }
    
    return results    
}

