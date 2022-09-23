// Code generated by astool. DO NOT EDIT.

package propertysharedinbox

import (
	"fmt"
	anyuri "github.com/superseriousbusiness/activity/streams/values/anyURI"
	vocab "github.com/superseriousbusiness/activity/streams/vocab"
	"net/url"
)

// ActivityStreamsSharedInboxProperty is the functional property "sharedInbox". It
// is permitted to be a single nilable value type.
type ActivityStreamsSharedInboxProperty struct {
	xmlschemaAnyURIMember *url.URL
	unknown               interface{}
	alias                 string
}

// DeserializeSharedInboxProperty creates a "sharedInbox" property from an
// interface representation that has been unmarshalled from a text or binary
// format.
func DeserializeSharedInboxProperty(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsSharedInboxProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "sharedInbox"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "sharedInbox")
	}
	i, ok := m[propName]

	if ok {
		if v, err := anyuri.DeserializeAnyURI(i); err == nil {
			this := &ActivityStreamsSharedInboxProperty{
				alias:                 alias,
				xmlschemaAnyURIMember: v,
			}
			return this, nil
		}
		this := &ActivityStreamsSharedInboxProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsSharedInboxProperty creates a new sharedInbox property.
func NewActivityStreamsSharedInboxProperty() *ActivityStreamsSharedInboxProperty {
	return &ActivityStreamsSharedInboxProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsXMLSchemaAnyURI
// afterwards will return false.
func (this *ActivityStreamsSharedInboxProperty) Clear() {
	this.unknown = nil
	this.xmlschemaAnyURIMember = nil
}

// Get returns the value of this property. When IsXMLSchemaAnyURI returns false,
// Get will return any arbitrary value.
func (this ActivityStreamsSharedInboxProperty) Get() *url.URL {
	return this.xmlschemaAnyURIMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this ActivityStreamsSharedInboxProperty) GetIRI() *url.URL {
	return this.xmlschemaAnyURIMember
}

// HasAny returns true if the value or IRI is set.
func (this ActivityStreamsSharedInboxProperty) HasAny() bool {
	return this.IsXMLSchemaAnyURI()
}

// IsIRI returns true if this property is an IRI.
func (this ActivityStreamsSharedInboxProperty) IsIRI() bool {
	return this.xmlschemaAnyURIMember != nil
}

// IsXMLSchemaAnyURI returns true if this property is set and not an IRI.
func (this ActivityStreamsSharedInboxProperty) IsXMLSchemaAnyURI() bool {
	return this.xmlschemaAnyURIMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsSharedInboxProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	var child map[string]string

	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this ActivityStreamsSharedInboxProperty) KindIndex() int {
	if this.IsXMLSchemaAnyURI() {
		return 0
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsSharedInboxProperty) LessThan(o vocab.ActivityStreamsSharedInboxProperty) bool {
	if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsXMLSchemaAnyURI() && !o.IsXMLSchemaAnyURI() {
		// Both are unknowns.
		return false
	} else if this.IsXMLSchemaAnyURI() && !o.IsXMLSchemaAnyURI() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsXMLSchemaAnyURI() && o.IsXMLSchemaAnyURI() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return anyuri.LessAnyURI(this.Get(), o.Get())
	}
}

// Name returns the name of this property: "sharedInbox".
func (this ActivityStreamsSharedInboxProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "sharedInbox"
	} else {
		return "sharedInbox"
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsSharedInboxProperty) Serialize() (interface{}, error) {
	if this.IsXMLSchemaAnyURI() {
		return anyuri.SerializeAnyURI(this.Get())
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsXMLSchemaAnyURI afterwards will
// return true.
func (this *ActivityStreamsSharedInboxProperty) Set(v *url.URL) {
	this.Clear()
	this.xmlschemaAnyURIMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *ActivityStreamsSharedInboxProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.Set(v)
}