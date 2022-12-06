# GoCache

	c := GoCache.NewCache()
  
	key := "1"
	value := 2
  
	c.Set(key, value)
	fmt.Printf("Set element. Number of elements in the cache %d \n", c.NumberOfElement)
  
	a, err := c.Get("3")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("value %s is %v\n", key, a)
	}

	c.Delete(key)
	fmt.Printf("Delete %s number of elements in the cache %d", key, c.NumberOfElement)
