# GoCache
Ensure the library is installed and up to date by running  
go get -u github.com/bigmihan/GoCache

	//   Sorry for my english. I`m a programmer 1C :)

	// key and value in Cache. For example
	key := "1"
	value := 2
	timeToLive := time.Second * 2

	CleanupInterval := time.Second * 1
	startCleanup := true // start goruten in NewCache

	c := GoCache.NewCache(CleanupInterval, startCleanup)
	c.Set(key, value, timeToLive)

	//c.Set("long", 100500, time.Second*10)

	_, err := c.Get(key)
	if err == nil {
		fmt.Println("found element")

	}

	// wait cleanup data
	time.Sleep(time.Second * 5)
	//c.Cleanup()

	v, err := c.Get("1")
	if err == nil {
		fmt.Printf("found element %v \n", v)
	} else {
		fmt.Println(err)
	}

	ok := c.Delete(key)
	if ok {
		fmt.Printf("Delete %s, number of elements in the cache %d", key, c.CountOfElement())
	} else {
		fmt.Printf("not found %s, number of elements in the cache %d", key, c.CountOfElement())
	}
	
	
	
	
	
For test
Worker pool
	
	countCache := 15

	CleanupInterval := time.Second * 100
	startCleanup := false // start goruten in NewCache
	c := GoCache.NewCache(CleanupInterval, startCleanup)

	for i := 0; i < countCache; i++ {

		if i < 5 {
			c.Set(fmt.Sprintf("key %d", i), i, time.Second*100)
		} else {
			c.Set(fmt.Sprintf("key %d", i), i, time.Second*1)
		}

	}

	time.Sleep(time.Second * 2)

	go c.Cleanup()

	time.Sleep(time.Second * 1)
	fmt.Printf("len(cache)=%d", c.CountOfElement())
	
	
