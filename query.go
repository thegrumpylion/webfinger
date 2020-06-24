package webfinger

import "net/url"

// Query represents a query request to the server
type Query struct {
	Resource string
	Rel      []string
}

// NewQuery returns a new query for resource res
func NewQuery(res string) *Query {
	return &Query{
		Resource: res,
	}
}

// QueryFromValues returns a new query from URL values.
// Useful when implementing the server side of the protocol
func QueryFromValues(v url.Values) *Query {
	return &Query{
		Resource: v.Get("resource"),
		Rel:      v["rel"],
	}
}

// AddRel adds rel to the query
func (q *Query) AddRel(rel string) {
	q.Rel = append(q.Rel, rel)
}

// ToValues converts the query to URL values
func (q *Query) ToValues() url.Values {
	v := url.Values{}
	v.Add("resource", q.Resource)
	if q.Rel != nil {
		v["rel"] = q.Rel
	}
	return v
}
