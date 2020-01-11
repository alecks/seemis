package main

func throw(e error) {
	if e != nil {
		panic(e)
	}
}