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

type CT_Formulas struct {
	F []*CT_F
}

func NewCT_Formulas() *CT_Formulas {
	ret := &CT_Formulas{}
	return ret
}

func (m *CT_Formulas) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.F != nil {
		sef := xml.StartElement{Name: xml.Name{Local: "v:f"}}
		for _, c := range m.F {
			e.EncodeElement(c, sef)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Formulas) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Formulas:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "f"}:
				tmp := NewCT_F()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.F = append(m.F, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_Formulas %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Formulas
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Formulas and its children
func (m *CT_Formulas) Validate() error {
	return m.ValidateWithPath("CT_Formulas")
}

// ValidateWithPath validates the CT_Formulas and its children, prefixing error messages with path
func (m *CT_Formulas) ValidateWithPath(path string) error {
	for i, v := range m.F {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/F[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
