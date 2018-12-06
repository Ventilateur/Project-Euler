package euler

import (
    "fmt"
    "math"
)


func max(a []int) (int, int) {
    max := int(math.MinInt32)
    idx := 0;
	
    for i, v := range a {
        if v >= max {
            max = v
            idx = i
        }
    }
	
    return idx, max
}


func swap(x, y int, a []int) {
    tmp := a[x]
    a[x] = a[y]
    a[y] = tmp
}


func reorder(r []int) {
    i := false
    tmpSlice := r
    maxIdx := 0
    lowerBound, higherBound := 0, len(r)
	
    for lowerBound < higherBound {
        tmpSlice = r[lowerBound : higherBound]
        maxIdx, _ = max(tmpSlice)
		
        if i {
          swap(len(tmpSlice) - 1, maxIdx, tmpSlice)
          higherBound--
        } else {
          swap(0, maxIdx, tmpSlice)
          lowerBound++
        }
		
        i = !i
    }
}


func compute(R, n int, r []int) int {
    
    nbBalls := len(r)
    if nbBalls == 1 {
        return 2000 * r[0]
    }
    
    var sumr int
    var l float64
    sumr, l = 0, 0.0
    
    reorder(r)
    
    for i, _ := range r {
        if i < nbBalls - 1 {
            sumr = r[i] + r[i + 1]
            l += math.Sqrt(math.Pow(float64(sumr), 2) - math.Pow(float64(2 * R - sumr), 2)) * 1000
        }
    }  
	
    return int(math.Ceil(l)) + 1000 * (r[0] + r[nbBalls - 1])
}

