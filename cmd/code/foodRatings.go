package code

import "container/heap"

type FoodRatings struct {
	foodToCuisine map[string] string
	foodToRatings map[string] int
	cuisineHeap map[string] *MaxHeap
}

type FoodItem struct {
	food string
	rating int
}

// custom max-heap
type MaxHeap []FoodItem

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].food < h[j].food
	}

	return h[i].rating < h[j].rating
}

func (h MaxHeap) Swap(i,j int) {
	h[i], h[j] = h[j], h[i]
}

func(h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(FoodItem))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	
	return item
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRating:  make(map[string]int),
		cuisineHeap:   make(map[string]*MaxHeap),
	}

	for i := 0; i < len(foods); i++ {
		food, cuisine, rating := foods[i], cuisines[i], ratings[i]
		fr.foodToCuisine[food] = cuisine
		fr.foodToRating[food] = rating
		if _, ok := fr.cuisineHeap[cuisine]; !ok {
			fr.cuisineHeap[cuisine] = &MaxHeap{}
		}
		heap.Push(fr.cuisineHeap[cuisine], FoodItem{food, rating})
	}
	return fr
}

// ChangeRating
func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	cuisine := fr.foodToCuisine[food]
	fr.foodToRating[food] = newRating
	heap.Push(fr.cuisineHeap[cuisine], FoodItem{food, newRating})
}

// HighestRated
func (fr *FoodRatings) HighestRated(cuisine string) string {
	h := fr.cuisineHeap[cuisine]
	for {
		top := (*h)[0]
		// Check if top is still valid
		if fr.foodToRating[top.food] == top.rating {
			return top.food
		}
		heap.Pop(h) // discard outdated entry
	}
}

/**
 * Example usage:
 * obj := Constructor([]string{"kimchi","miso","sushi","moussaka","ramen","bulgogi"},
 *                    []string{"korean","japanese","japanese","greek","japanese","korean"},
 *                    []int{9,12,8,15,14,7});
 *
 * fmt.Println(obj.HighestRated("korean"))   // kimchi
 * fmt.Println(obj.HighestRated("japanese")) // ramen
 * obj.ChangeRating("sushi",16);
 * fmt.Println(obj.HighestRated("japanese")) // sushi
 */