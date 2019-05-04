// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package relationships

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type Relationships struct {
	CT_Relationships
}

func NewRelationships() *Relationships {
	ret := &Relationships{}
	ret.CT_Relationships = *NewCT_Relationships()
	return ret
}

func (m *Relationships) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/package/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "Relationships"
	return m.CT_Relationships.MarshalXML(e, start)
}

func (m *Relationships) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Relationships = *NewCT_Relationships()
lRelationships:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/relationships", Local: "Relationship"}:
				tmp := NewRelationship()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Relationship = append(m.Relationship, tmp)
			default:
				gooxml.Log("skipping unsupported element on Relationships %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lRelationships
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Relationships and its children
func (m *Relationships) Validate() error {
	return m.ValidateWithPath("Relationships")
}

// ValidateWithPath validates the Relationships and its children, prefixing error messages with path
func (m *Relationships) ValidateWithPath(path string) error {
	if err := m.CT_Relationships.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
