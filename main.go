package main

import "aulas/basics"

func main() {
	params := basics.URLParams{
		"limit": []string{"5"},
		"skip":  []string{"0"},
	}
	url := basics.NewURL(
		"https://dummyjson.com",
		"/products",
		params,
	)

	url.Params().Add("select", "title")
	url.Params().Add("select", "price")
	url.Params().Add("select", "description")
	url.Params().Add("select", "category")

	req := basics.NewRequest(basics.HTTPMethodGet, *url)
	req.Execute()
}
