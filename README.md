# PACKAGE DOCUMENTATION

package hrefs
    import "github.com/markhuge/hrefs"

    Package hrefs is a library for extracting URIs from HTML elements.

    This includes:

	- URLs from <a> tags
	- image URIs from <img> tags
	- stylesheet URIs from <link> tags
	- script URIs from <script> tags

## FUNCTIONS

func Get(t html.Token) (uri string)
    Get is a wrapper for all the other functions. It attempts to determine
    the correct function based on the tag type

func GetCSS(t html.Token) (href string)
    GetCSS is a helper to get a stylesheet href from a <link> tag

func GetHref(t html.Token) (href string)
    GetHref returns href values when present

func GetSrc(t html.Token) (src string)
    GetSrc returns src values when present


