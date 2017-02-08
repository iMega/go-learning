package main

import (
	"fmt"
	"net/url"
)

type googleQuery struct {
	Ie    string
	Oe    string
	GwsRd string
	Query string
}

func main() {
	gq := new(googleQuery)
	gq.Ie = "utf-8"
	gq.Oe = "utf-8"
	gq.GwsRd = "cr"
	gq.Query = "My learning curve"

	query := BuildQuery(gq)

	uri := url.URL{
		Host:       "google.com",
		Scheme:     "https",
		Path:       "search",
		ForceQuery: true,
		RawQuery:   query,
	}

	fmt.Println(uri.String())
}

func BuildQuery(gq *googleQuery) string {
	/*
	queryParams := url.Values{}
	queryParams.Add("ie", gq.Ie)
	queryParams.Add("oe", gq.Oe)
	queryParams.Add("q", gq.Query)
	*/
	queryParams := url.Values{
		"ie": []string{gq.Ie},
		"oe": []string{gq.Oe},
		"q":  []string{gq.Query},
	}

	return queryParams.Encode()
}
