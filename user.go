package main

type user struct {
	name string
}

func newUser(n string) user {
	u := user{
		name: n,
	}

	return u
}
