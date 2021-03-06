// Code generated by https://github.com/GoComply/metaschema; DO NOT EDIT.
package catalog

import (
	"encoding/xml"

	"github.com/gocomply/oscalkit/types/oscal/validation_root"

	"github.com/gocomply/oscalkit/types/oscal/nominal_catalog"
)

// A collection of controls.
type Catalog struct {
	XMLName xml.Name `xml:"http://csrc.nist.gov/ns/oscal/1.0 catalog" json:"-"`
	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`

	// Provides information about the publication and availability of the containing document.
	Metadata *Metadata `xml:"metadata,omitempty" json:"metadata,omitempty"`
	// Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
	Parameters []Param `xml:"param,omitempty" json:"parameters,omitempty"`
	// A structured information object representing a security or privacy control. Each security or privacy control within the Catalog is defined by a distinct control instance.
	Controls []Control `xml:"control,omitempty" json:"controls,omitempty"`
	// A group of controls, or of groups of controls.
	Groups []Group `xml:"group,omitempty" json:"groups,omitempty"`
	// Back matter including references and resources.
	BackMatter *BackMatter `xml:"back-matter,omitempty" json:"backMatter,omitempty"`
}

// A group of controls, or of groups of controls.
type Group struct {

	// Unique identifier of the containing object
	Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Indicating the type or classification of the containing object
	Class string `xml:"class,attr,omitempty" json:"class,omitempty"`

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
	Parameters []Param `xml:"param,omitempty" json:"parameters,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
	// A partition or component of a control or part
	Parts []Part `xml:"part,omitempty" json:"parts,omitempty"`
	// A group of controls, or of groups of controls.
	Groups []Group `xml:"group,omitempty" json:"groups,omitempty"`
	// A structured information object representing a security or privacy control. Each security or privacy control within the Catalog is defined by a distinct control instance.
	Controls []Control `xml:"control,omitempty" json:"controls,omitempty"`
}

// A structured information object representing a security or privacy control. Each security or privacy control within the Catalog is defined by a distinct control instance.
type Control struct {

	// Unique identifier of the containing object
	Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Indicating the type or classification of the containing object
	Class string `xml:"class,attr,omitempty" json:"class,omitempty"`

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
	Parameters []Param `xml:"param,omitempty" json:"parameters,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
	// A partition or component of a control or part
	Parts []Part `xml:"part,omitempty" json:"parts,omitempty"`
	// A structured information object representing a security or privacy control. Each security or privacy control within the Catalog is defined by a distinct control instance.
	Controls []Control `xml:"control,omitempty" json:"controls,omitempty"`
}

type Annotation = validation_root.Annotation

type BackMatter = validation_root.BackMatter

type Link = validation_root.Link

type Metadata = validation_root.Metadata

type Param = nominal_catalog.Param

type Part = nominal_catalog.Part

type Prop = validation_root.Prop

type Title = validation_root.Title
