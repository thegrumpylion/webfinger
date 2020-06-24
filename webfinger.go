package webfinger

// https://tools.ietf.org/html/rfc7033

// Link is a webfinger resource link
type Link struct {
	// Rel is a string that is either a URI or a
	// registered relation type (see RFC 5988).  The value of the
	// "rel" member MUST contain exactly one URI or registered relation
	// type.  The URI or registered relation type identifies the type of the
	// link relation.
	//
	// The other members of the object have meaning only once the type of
	// link relation is understood.  In some instances, the link relation
	// will have associated semantics enabling the client to query for other
	// resources on the Internet.  In other instances, the link relation
	// will have associated semantics enabling the client to utilize the
	// other members of the link relation object without fetching additional
	// external resources.
	//
	// URI link relation type values are compared using the "Simple String
	// Comparison" algorithm of Section 6.2.1 of RFC 3986.
	//
	// The "rel" member MUST be present in the link relation object.
	Rel string `json:"rel"`

	// Type is a string that indicates the media
	// type of the target resource (see RFC 6838).
	// The "type" member is OPTIONAL in the link relation object.
	Type string `json:"type,omitempty"`

	// Href member is a string that contains a URI
	// pointing to the target resource.
	// The "href" member is OPTIONAL in the link relation object.
	Href string `json:"href,omitempty"`

	// Titles object comprises zero or more name/value pairs whose
	// names are a language tag or the string "und".  The string is
	// human-readable and describes the link relation.  More than one title
	// for the link relation MAY be provided for the benefit of users who
	// utilize the link relation, and, if used, a language identifier SHOULD
	// be duly used as the name.  If the language is unknown or unspecified,
	// then the name is "und".
	//
	// A JRD SHOULD NOT include more than one title identified with the same
	// language tag (or "und") within the link relation object.  Meaning is
	// undefined if a link relation object includes more than one title
	// named with the same language tag (or "und"), though this MUST NOT be
	// treated as an error.  A client MAY select whichever title or titles
	// it wishes to utilize.
	//
	// The "titles" member is OPTIONAL in the link relation object.
	Titles map[string]string `json:"titles,omitempty"`

	// Properties object within the link relation object comprises
	// zero or more name/value pairs whose names are URIs (referred to as
	// "property identifiers") and whose values are strings or null.
	// Properties are used to convey additional information about the link
	// relation.
	//
	// The "properties" member is OPTIONAL in the link relation object.
	Properties map[string]*string `json:"properties,omitempty"`

	// Template string for use with PayID
	// https://github.com/payid-org/rfcs/blob/master/dist/spec/payid-discovery.txt
	// This is not part for RFC7033
	Template string `json:"template,omitempty"`
}

// Resource is the webfinger resource
type Resource struct {
	// The value of the "subject" member is a URI that identifies the entity
	// that the JRD describes.
	//
	// The "subject" value returned by a WebFinger resource MAY differ from
	// the value of the "resource" parameter used in the client's request.
	// This might happen, for example, when the subject's identity changes
	// (e.g., a user moves his or her account to another service) or when
	// the resource prefers to express URIs in canonical form.
	//
	// The "subject" member SHOULD be present in the JRD.
	Subject string `json:"subject"`

	// The "aliases" array is an array of zero or more URI strings that
	// identify the same entity as the "subject" URI.
	//
	// The "aliases" array is OPTIONAL in the JRD.
	Aliases []string `json:"aliases,omitempty"`

	// The "properties" object comprises zero or more name/value pairs whose
	// names are URIs (referred to as "property identifiers") and whose
	// values are strings or null.  Properties are used to convey additional
	// information about the subject of the JRD.
	//
	// The "properties" member is OPTIONAL in the JRD.
	Properties map[string]*string `json:"properties,omitempty"`

	// The "links" array has any number of member objects, each of which
	// represents a link
	//
	// The "links" array is OPTIONAL in the JRD.
	Links []*Link `json:"links,omitempty"`
}
