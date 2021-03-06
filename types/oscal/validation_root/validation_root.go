// Code generated by https://github.com/GoComply/metaschema; DO NOT EDIT.
package validation_root

import ()

// NOT TO BE USED IN A METASCHEMA
type VALIDATIONRoot struct {

	// A description supporting the parent item.
	Description *Description `xml:"description,omitempty" json:"description,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// References a  defined in .
	PartyUuid PartyUuid `xml:"party-uuid,omitempty" json:"partyUuid,omitempty"`
	// Provides information about the publication and availability of the containing document.
	Metadata *Metadata `xml:"metadata,omitempty" json:"metadata,omitempty"`
	// A collection of citations and resource references.
	BackMatter *BackMatter `xml:"back-matter,omitempty" json:"backMatter,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotation *Annotation `xml:"annotation,omitempty" json:"annotation,omitempty"`
	// A reference to a set of organizations or persons that have responsibility for performing a referenced role relative to the parent context.
	ResponsibleParties []ResponsibleParty `xml:"responsible-party,omitempty" json:"responsible-parties,omitempty"`
}

// Provides information about the publication and availability of the containing document.
type Metadata struct {

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// The date and time this document was published.
	Published Published `xml:"published,omitempty" json:"published,omitempty"`
	// Date and time of last modification.
	LastModified LastModified `xml:"last-modified,omitempty" json:"lastModified,omitempty"`
	// The version of the document content.
	Version Version `xml:"version,omitempty" json:"version,omitempty"`
	// OSCAL model version.
	OscalVersion OscalVersion `xml:"oscal-version,omitempty" json:"oscalVersion,omitempty"`
	// A document identifier qualified by an identifier .
	DocumentIds []DocId `xml:"doc-id,omitempty" json:"document-ids,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// An entry in a sequential list of revisions to the containing document in reverse chronological order (i.e., most recent previous revision first).
	RevisionHistory []Revision `xml:"revision,omitempty" json:"revision-history,omitempty"`
	// Defining a role to be assigned to a party
	Roles []Role `xml:"role,omitempty" json:"roles,omitempty"`
	// A location, with associated metadata that can be referenced.
	Locations []Location `xml:"location,omitempty" json:"locations,omitempty"`
	// A responsible entity, either singular (an organization or person) or collective (multiple persons)
	Parties []Party `xml:"party,omitempty" json:"parties,omitempty"`
	// A reference to a set of organizations or persons that have responsibility for performing a referenced role relative to the parent context.
	ResponsibleParties []ResponsibleParty `xml:"responsible-party,omitempty" json:"responsible-parties,omitempty"`
}

// A collection of citations and resource references.
type BackMatter struct {

	// A resource associated with the present document, which may be a pointer to other data or a citation.
	Resources []Resource `xml:"resource,omitempty" json:"resources,omitempty"`
}

// An entry in a sequential list of revisions to the containing document in reverse chronological order (i.e., most recent previous revision first).
type Revision struct {

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// The date and time this document was published.
	Published Published `xml:"published,omitempty" json:"published,omitempty"`
	// Date and time of last modification.
	LastModified LastModified `xml:"last-modified,omitempty" json:"lastModified,omitempty"`
	// The version of the document content.
	Version Version `xml:"version,omitempty" json:"version,omitempty"`
	// OSCAL model version.
	OscalVersion OscalVersion `xml:"oscal-version,omitempty" json:"oscalVersion,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
}

// A name/value pair with optional explanatory remarks.
type Annotation struct {

	// Identifying the purpose and intended use of the property, part or other object.
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`
	// A namespace qualifying the name.
	Ns string `xml:"ns,attr,omitempty" json:"ns,omitempty"`
	// Indicates the value of the characteristic.
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`

	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
}

// A location, with associated metadata that can be referenced.
type Location struct {

	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// Email address
	EmailAddresses []Email `xml:"email,omitempty" json:"email-addresses,omitempty"`
	// Contact number by telephone
	TelephoneNumbers []Phone `xml:"phone,omitempty" json:"telephone-numbers,omitempty"`
	// URL for web site or Internet presence
	URLs []Url `xml:"url,omitempty" json:"URLs,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// A postal address.
	Address *Address `xml:"address,omitempty" json:"address,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
}

// A responsible entity, either singular (an organization or person) or collective (multiple persons)
type Party struct {

	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`
	// A category describing the kind of party the object describes.
	Type string `xml:"type,attr,omitempty" json:"type,omitempty"`

	// The full (legal) name of the party.
	PartyName PartyName `xml:"party-name,omitempty" json:"partyName,omitempty"`
	// A common name, short name or acronym
	ShortName ShortName `xml:"short-name,omitempty" json:"shortName,omitempty"`
	// An identifier for a person (such as an ORCID) using a designated scheme.
	ExternalIds []ExternalId `xml:"external-id,omitempty" json:"external-ids,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Email address
	EmailAddresses []Email `xml:"email,omitempty" json:"email-addresses,omitempty"`
	// Contact number by telephone
	TelephoneNumbers []Phone `xml:"phone,omitempty" json:"telephone-numbers,omitempty"`
	// Identifies that the containing object is a member of the organization associated with the provided UUID.
	MemberOfOrganizations []MemberOfOrganization `xml:"member-of-organization,omitempty" json:"member-of-organizations,omitempty"`
	// References a  defined in .
	LocationUuids []LocationUuid `xml:"location-uuid,omitempty" json:"location-uuids,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
	// A postal address.
	Addresses []Address `xml:"address,omitempty" json:"addresses,omitempty"`
}

// A pointer to an external copy of a document with optional hash for verification
type Rlink struct {

	// A link to a document or document fragment (actual, nominal or projected)
	Href string `xml:"href,attr,omitempty" json:"href,omitempty"`
	// Describes the media type of the linked resource
	MediaType string `xml:"media-type,attr,omitempty" json:"mediaType,omitempty"`

	// A representation of a cryptographic digest generated over a resource using a hash algorithm.
	Hashes []Hash `xml:"hash,omitempty" json:"hashes,omitempty"`
}

// A postal address.
type Address struct {

	// Indicates the type of address.
	Type string `xml:"type,attr,omitempty" json:"type,omitempty"`

	// A single line of an address.
	PostalAddress []AddrLine `xml:"addr-line,omitempty" json:"postal-address,omitempty"`
	// City, town or geographical region for mailing address
	City City `xml:"city,omitempty" json:"city,omitempty"`
	// State, province or analogous geographical region for mailing address
	State State `xml:"state,omitempty" json:"state,omitempty"`
	// Postal or ZIP code for mailing address
	PostalCode PostalCode `xml:"postal-code,omitempty" json:"postalCode,omitempty"`
	// Country for mailing address
	Country Country `xml:"country,omitempty" json:"country,omitempty"`
}

// A container in which a set of bibliographic information can included. The model of this information is undefined by OSCAL.
type Biblio struct {
}

// A resource associated with the present document, which may be a pointer to other data or a citation.
type Resource struct {

	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// A short textual description
	Desc Desc `xml:"desc,omitempty" json:"desc,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A document identifier qualified by an identifier .
	DocumentIds []DocId `xml:"doc-id,omitempty" json:"document-ids,omitempty"`
	//
	Attachments []Base64 `xml:"base64,omitempty" json:"attachments,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// A citation consisting of end note text and optional structured bibliographic data.
	Citation *Citation `xml:"citation,omitempty" json:"citation,omitempty"`
	// A pointer to an external copy of a document with optional hash for verification
	Rlinks []Rlink `xml:"rlink,omitempty" json:"rlinks,omitempty"`
}

// A citation consisting of end note text and optional structured bibliographic data.
type Citation struct {

	// A line of textual content whose semantic is determined by the context of use.
	Text Text `xml:"text,omitempty" json:"text,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A container in which a set of bibliographic information can included. The model of this information is undefined by OSCAL.
	Biblio *Biblio `xml:"biblio,omitempty" json:"biblio,omitempty"`
}

// Defining a role to be assigned to a party
type Role struct {

	// Unique identifier of the containing object
	Id string `xml:"id,attr,omitempty" json:"id,omitempty"`

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// A common name, short name or acronym
	ShortName ShortName `xml:"short-name,omitempty" json:"shortName,omitempty"`
	// A short textual description
	Desc Desc `xml:"desc,omitempty" json:"desc,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
}

// A reference to a set of organizations or persons that have responsibility for performing a referenced role relative to the parent context.
type ResponsibleParty struct {

	// The role that the party is responsible for.
	RoleId string `xml:"role-id,attr,omitempty" json:"roleId,omitempty"`

	// References a  defined in .
	PartyUuids []PartyUuid `xml:"party-uuid,omitempty" json:"party-uuids,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
}

// A reference to a local or remote resource
type Link struct {
	// A link to a document or document fragment (actual, nominal or projected)
	Href string `xml:"href,attr,omitempty" json:"href,omitempty"`

	// Describes the type of relationship provided by the link. This can be an indicator of the link's purpose.
	Rel string `xml:"rel,attr,omitempty" json:"rel,omitempty"`

	// Describes the media type of the linked resource
	MediaType string `xml:"media-type,attr,omitempty" json:"mediaType,omitempty"`
	Value     string `xml:",chardata" json:"value,omitempty"`
}

// The date and time this document was published.

type Published string

// Date and time of last modification.

type LastModified string

// The version of the document content.

type Version string

// OSCAL model version.

type OscalVersion string

// A document identifier qualified by an identifier .
type DocId struct {
	// Qualifies the kind of document identifier.
	Type  string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Value string `xml:",chardata" json:"value,omitempty"`
}

// A value with a name, attributed to the containing control, part, or group.
type Prop struct {
	// Identifying the purpose and intended use of the property, part or other object.
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`

	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`

	// A namespace qualifying the name.
	Ns string `xml:"ns,attr,omitempty" json:"ns,omitempty"`

	// Indicating the type or classification of the containing object
	Class string `xml:"class,attr,omitempty" json:"class,omitempty"`
	Value string `xml:",chardata" json:"value,omitempty"`
}

// References a  defined in .

type LocationUuid string

// References a  defined in .

type PartyUuid string

// An identifier for a person (such as an ORCID) using a designated scheme.
type ExternalId struct {
	// Indicating the type of identifier, address, email or other data item.
	Type  string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Value string `xml:",chardata" json:"value,omitempty"`
}

// Identifies that the containing object is a member of the organization associated with the provided UUID.

type MemberOfOrganization string

// The full (legal) name of the party.

type PartyName string

// A common name, short name or acronym

type ShortName string

// A single line of an address.

type AddrLine string

// City, town or geographical region for mailing address

type City string

// State, province or analogous geographical region for mailing address

type State string

// Postal or ZIP code for mailing address

type PostalCode string

// Country for mailing address

type Country string

// Email address

type Email string

// Contact number by telephone
type Phone struct {
	// Indicates the type of phone number.
	Type  string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Value string `xml:",chardata" json:"value,omitempty"`
}

// URL for web site or Internet presence

type Url string

// A short textual description

type Desc string

// A line of textual content whose semantic is determined by the context of use.

type Text string

// A representation of a cryptographic digest generated over a resource using a hash algorithm.
type Hash struct {
	// Method by which a hash is derived
	Algorithm string `xml:"algorithm,attr,omitempty" json:"algorithm,omitempty"`
	Value     string `xml:",chardata" json:"value,omitempty"`
}

// A title for display and navigation

type Title string

//
type Base64 struct {
	// Name of the file before it was encoded as Base64 to be embedded in a . This is the name that will be assigned to the file when the file is decoded.
	Filename string `xml:"filename,attr,omitempty" json:"filename,omitempty"`

	// Describes the media type of the linked resource
	MediaType string `xml:"media-type,attr,omitempty" json:"mediaType,omitempty"`
	Value     string `xml:",chardata" json:"value,omitempty"`
}

// A description supporting the parent item.

type Description = Markup

// Additional commentary on the parent item.

type Remarks = Markup
