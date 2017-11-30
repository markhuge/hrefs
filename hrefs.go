/*
Package hrefs is a library for extracting URIs from HTML elements.

This includes:
 - URLs from <a> tags
 - image URIs from <img> tags
 - stylesheet URIs from <link> tags
 - script URIs from <script> tags
*/
package hrefs

import "golang.org/x/net/html"

/*
Get is a wrapper for all the other functions. It attempts to determine the
correct function based on the tag type
*/
func Get(t html.Token) (uri string) {
	switch {
	case t.Data == "a":
		uri = GetHref(t)
	case t.Data == "img":
		uri = GetSrc(t)
	case t.Data == "script":
		uri = GetSrc(t)
	case t.Data == "link":
		uri = GetCSS(t)
	}
	return
}

// GetHref returns href values when present
func GetHref(t html.Token) (href string) {
	for _, a := range t.Attr {
		if a.Key == "href" && a.Val != "#" {
			href = a.Val
		}
	}
	return
}

// GetSrc returns src values when present
func GetSrc(t html.Token) (src string) {
	for _, a := range t.Attr {
		if a.Key == "src" {
			src = a.Val
		}
	}
	return
}

// GetCSS is a helper to get a stylesheet href from a <link> tag
func GetCSS(t html.Token) (href string) {
	for _, a := range t.Attr {
		if a.Key == "rel" && a.Val == "stylesheet" {
			href = GetHref(t)
		}
	}
	return
}
