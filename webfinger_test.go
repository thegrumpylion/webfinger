package webfinger_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/thegrumpylion/webfinger"
)

func getTestDB() webfinger.DB {
	return webfinger.NewMemDB(map[string][]byte{
		"http://blog.example.com/article/id/314": []byte(`
		{
			"subject" : "http://blog.example.com/article/id/314",
			"aliases" :
			[
				"http://blog.example.com/cool_new_thing",
				"http://blog.example.com/steve/article/7"
			],
			"properties" :
			{
				"http://blgx.example.net/ns/version" : "1.3",
				"http://blgx.example.net/ns/ext" : null
			},
			"links" :
			[
				{
				"rel" : "copyright",
				"href" : "http://www.example.com/copyright"
				},
				{
				"rel" : "author",
				"href" : "http://blog.example.com/author/steve",
				"titles" :
				{
					"en-us" : "The Magical World of Steve",
					"fr" : "Le Monde Magique de Steve"
				},
				"properties" :
				{
					"http://example.com/role" : "editor"
				}
				}
		
			]
		}
		`),
	})
}

func TestHandler(t *testing.T) {
	h := webfinger.NewHandler(getTestDB(), webfinger.WithAllowOrigin("*"))

	s := httptest.NewServer(h)

	q := webfinger.NewQuery("http://blog.example.com/article/id/314")

	c, err := webfinger.NewClient(s.URL)
	if err != nil {
		t.Fatal(err)
	}

	r, err := c.Query(q)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(r)
}
