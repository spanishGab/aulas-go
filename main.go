package main

import "aulas/basics"

func main() {
	params := basics.URLParams{
		"limit": []string{"5"},
		"skip":  []string{"0"},
	}
	params.Add("select", "title")
	params.Add("select", "price")
	params.Add("select", "description")
	params.Add("select", "category")

	url := basics.NewURL(
		"https://dummyjson.com",
		"/products",
		params,
	)
	req := basics.NewRequest(basics.HTTPMethodGet, *url)
	req.Execute()
}
