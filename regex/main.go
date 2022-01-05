package main

import (
	"fmt"
	"regexp"
)

var rgxNext = regexp.MustCompile(`<(.*)>; rel="next"`)
var rgxPrev = regexp.MustCompile(`<(.*)>; rel="previous"`)

// var rgxPrev = regexp.MustCompile(`\((.*?)\)`)

func main() {
	s := `<https://chillgroup.myshopify.com/admin/api/2021-07/customers/search.json?limit=5&page_info=eyJkaXJlY3Rpb24iOiJwcmV2IiwibGFzdF9pZCI6NTEzOTMxOTI1OTMyMiwibGFzdF92YWx1ZSI6MTYxOTQwNTQ1ODAwMH0>; rel="previous", <https://chillgroup.myshopify.com/admin/api/2021-07/customers/search.json?limit=5&page_info=eyJkaXJlY3Rpb24iOiJuZXh0IiwibGFzdF9pZCI6NTEzOTMxNTg1MTQ1MCwibGFzdF92YWx1ZSI6MTYxOTQwNTM4NzAwMH0>; rel="next"`
	rsNext := rgxNext.FindStringSubmatch(s)
	rsPrev := rgxPrev.FindStringSubmatch(s)
	fmt.Println(rsNext[1])
	fmt.Println(rsPrev[1])
}
