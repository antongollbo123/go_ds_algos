package bubble_sort

func Sort(list *[]int) {
	if list == nil || len(*list) == 0 {
		return // early return if the list is nil or empty
	}
	l := *list

	for idx := 0; idx < len(l); idx++ {
		swap := false
		for jdx := 0; jdx < len(l)-idx-1; jdx++ {
			if l[jdx] > l[jdx+1] {
				l[jdx], l[jdx+1] = l[jdx+1], l[jdx]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}
