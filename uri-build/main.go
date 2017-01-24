package main

import (
	"fmt"
	"net/url"
)

func main() {
	queryParams := url.Values{}
	queryParams.Add("ie", "utf-8")
	queryParams.Add("oe", "utf-8")
	queryParams.Add("gws_rd", "cr")
	queryParams.Add("q", "My learning curve")
	query := queryParams.Encode()

	uri := url.URL{
		Host: "google.com",
		Scheme: "https",
		Path: "search",
		ForceQuery: true,
		RawQuery: query,
	}

	fmt.Println(uri.String())
}