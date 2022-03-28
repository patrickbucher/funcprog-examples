package main

import "fmt"

type Resource struct {
	Slots int
	Used  int
	Err   error
}

func New(slots int) *Resource {
	return &Resource{
		Slots: slots,
		Used:  0,
	}
}

func (r *Resource) Use(n int) *Resource {
	newUsed := r.Used + n
	if newUsed > r.Slots {
		return &Resource{
			Slots: r.Slots,
			Used:  newUsed,
			Err:   fmt.Errorf("use %d resources with %d slots", newUsed, r.Slots),
		}
	}
	return &Resource{
		Slots: r.Slots,
		Used:  newUsed,
		Err:   nil,
	}
}

func (r Resource) String() string {
	if r.Err != nil {
		return fmt.Sprintf("%v", r.Err)
	}
	return fmt.Sprintf("%d/%d resources used", r.Used, r.Slots)
}

func main() {
	resource := New(5).Use(1).Use(3).Use(1)
	fmt.Println(resource)

	resource = New(5).Use(4).Use(4).Use(4)
	fmt.Println(resource)
}
