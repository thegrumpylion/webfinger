package webfinger_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert := assert.New(t)

	h := webfinger.NewHandler(getTestDB(), webfinger.WithAllowOrigin("*"))

	s := httptest.NewServer(h)

	c, err := webfinger.NewClient(s.URL)
	assert.Nil(err)

	q := webfinger.NewQuery("http://blog.example.com/article/id/314")

	r, err := c.Query(q)
	assert.Nil(err)

	assert.Equal("http://blog.example.com/article/id/314", r.Subject)

	assert.Equal([]string{
		"http://blog.example.com/cool_new_thing",
		"http://blog.example.com/steve/article/7",
	}, r.Aliases)

	p, ok := r.Properties["http://blgx.example.net/ns/version"]
	assert.True(ok)
	assert.NotNil(p)
	assert.Equal("1.3", *p)

	p, ok = r.Properties["http://blgx.example.net/ns/ext"]
	assert.True(ok)
	assert.Nil(p)

	assert.Len(r.Links, 2)

	l0 := r.Links[0]
	assert.Equal("copyright", l0.Rel)
	assert.Equal("http://www.example.com/copyright", l0.Href)
	assert.Equal("copyright", l0.Rel)
	assert.Equal("copyright", l0.Rel)

	l1 := r.Links[1]
	assert.Equal("author", l1.Rel)
	assert.Equal("http://blog.example.com/author/steve", l1.Href)

	t1, ok := l1.Titles["en-us"]
	assert.True(ok)
	assert.Equal("The Magical World of Steve", t1)

	t2, ok := l1.Titles["fr"]
	assert.True(ok)
	assert.Equal("Le Monde Magique de Steve", t2)

	p, ok = l1.Properties["http://example.com/role"]
	assert.True(ok)
	assert.NotNil(p)
	assert.Equal("editor", *p)

}
