// Ikmal Nawawi Thoha
// 19102255
// MM4
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// penghitung waktu
	start := time.Now()
	
	fmt.Printf("\n\n")
	fmt.Println("***QUICK SORT***")
	slice := generateSlice(20)

	// print data sebelum diurutkan
	fmt.Println("Data sebelum diurutkan:\n", slice)
	quicksort(slice)

	// print data setelah diurutkan
	fmt.Println("Data setelah diurutkan:\n", slice, "\n")
	
	// durasi ekseskusi program
	fmt.Printf("Durasi eksekusi program: %v Detik \n", time.Since(start).Seconds())

	var space string
	fmt.Printf("Tekan enter untuk keluar")
	fmt.Scanln(&space)
}

// Menghasilkan sepotong ukuran, ukuran diisi dengan angka acak
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
 
func quicksort(a []int) []int {
    if len(a) < 2 {
        return a
    }
     
    left, right := 0, len(a)-1
     
    pivot := rand.Int() % len(a)
     
    a[pivot], a[right] = a[right], a[pivot]
     
    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }
     
    a[left], a[right] = a[right], a[left]
     
    quicksort(a[:left])
    quicksort(a[left+1:])
     
    return a
}