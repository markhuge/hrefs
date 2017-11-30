package hrefs_test

import (
	"strings"
	"testing"

	"github.com/markhuge/hrefs"

	"golang.org/x/net/html"
)

func TestGetA(t *testing.T) {
	t.Log("hrefs.Get <a> Tag")
	input := fixtureHelper("<a href=\"http://test.com\">test</a>")
	expected := "http://test.com"
	result := hrefs.Get(input)

	if result != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, result)
	}
}

func TestGetA_Hash(t *testing.T) {
	t.Log("Expect empty result hrefs.Get <a href=\"#\"> Tag")
	input := fixtureHelper("<a href=\"#\">test</a>")
	expected := ""
	result := hrefs.Get(input)

	if result != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, result)
	}
}

func TestGetImg(t *testing.T) {
	t.Log("hrefs.Get <img> Tag")
	input := fixtureHelper("<img src=\"/images/test.png\">")
	expected := "/images/test.png"
	result := hrefs.Get(input)

	if result != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, result)
	}
}

func TestGetLink(t *testing.T) {
	t.Log("hrefs.Get <link> Tag")
	input := fixtureHelper("<link rel=\"stylesheet\" href=\"/css/test.css\">")
	expected := "/css/test.css"
	result := hrefs.Get(input)

	if result != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, result)
	}
}

func TestGetScript(t *testing.T) {
	t.Log("hrefs.Get <script> Tag")
	input := fixtureHelper("<script src=\"https://some.cdn.test.js\"></script>")
	expected := "https://some.cdn.test.js"
	result := hrefs.Get(input)

	if result != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, result)
	}
}

func fixtureHelper(s string) html.Token {
	r := strings.NewReader(s)
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch {
		case tt == html.StartTagToken:
			return z.Token()
		}
	}

}
