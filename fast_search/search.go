package fast_search


func Search(arr []int, search int) int{
    left := 0
    right := len(arr)-1

    for left < right {
        mid := (left + right) / 2
        value := arr[mid]

        if value == search {
            return mid
        }else if value < search {
            left = mid + 1
        }else {
            right = mid - 1
        }

    }
    return -1
}