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
	
	
	
