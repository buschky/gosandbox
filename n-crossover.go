package kata
import "fmt"
import "sort"


func rvDup(el []int) []int {
    en := map[int]bool{}
    r := []int{}
    for v := range el {
        if en[el[v]] == true {
        } else {
            en[el[v]] = true
            r = append(r, el[v])
        }
    }
    return r
}

// ns : slice of indices
// xs, ys : chromosomes as slices of ints
func Crossover(ns []int, xs []int,ys []int) ([]int, []int) {
  // Your code here

ns = rvDup (ns)
sort.Ints(ns)
// clean ns 



for i := 0; i < len(ns); i++ {
	var pointer int = ns[i]
  fmt.Println(pointer)
	
  for j := 0; j < len(xs); j++ {
    
    if (j>=pointer) {
     // fmt.Println(pointer)	
      var buffer int = xs [j] 
      xs [j] = ys [j]
      ys [j] = buffer
    }
  }
}

 return xs,ys
}



func main() {
	co := []int{1, 4 }
  xs := []int{1, 2, 3, 4, 5, 6 }
  ys := []int{11, 12, 13, 14, 15, 16 }
  
  Crossover(co, xs, ys)
  return 

}
