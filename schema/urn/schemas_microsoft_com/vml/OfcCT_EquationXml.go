// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type OfcCT_EquationXml struct {
	ContentTypeAttr *string
	Any             gooxml.Any
}

func NewOfcCT_EquationXml() *OfcCT_EquationXml {
	ret := &OfcCT_EquationXml{}
	return ret
}

func (m *OfcCT_EquationXml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ContentTypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "contentType"},
			Value: fmt.Sprintf("%v", *m.ContentTypeAttr)})
	}
	e.EncodeToken(start)
	if m.Any != nil {
		m.Any.MarshalXML(e, xml.StartElement{})
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_EquationXml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "contentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentTypeAttr = &parsed
			continue
		}
	}
lOfcCT_EquationXml:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			default:
				if anyEl, err := gooxml.CreateElement(el); err != nil {
					return err
				} else {
					if err := d.DecodeElement(anyEl, &el); err != nil {
						return err
					}
					m.Any = anyEl
				}
			}
		case xml.EndElement:
			break lOfcCT_EquationXml
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcCT_EquationXml and its children
func (m *OfcCT_EquationXml) Validate() error {
	return m.ValidateWithPath("OfcCT_EquationXml")
}

// ValidateWithPath validates the OfcCT_EquationXml and its children, prefixing error messages with path
func (m *OfcCT_EquationXml) ValidateWithPath(path string) error {
	return nil
}
