package fast_sort


func Qsort (arr []int) []int {
   if len(arr) <= 1 {
      return arr
   }

   pivot := arr[0]
   var left, right []int
   for _, v := range arr[1:] {
      if v < pivot {
         left = append(left, v)
      }
      if v > pivot {
         right = append(right, v)
      }
      
   }

   left = Qsort(left)
   right = Qsort(right)

   return append(append(left, pivot), right...)
}
