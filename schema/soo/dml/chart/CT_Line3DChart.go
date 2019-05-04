// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_Line3DChart struct {
	Grouping   *CT_Grouping
	VaryColors *CT_Boolean
	Ser        []*CT_LineSer
	DLbls      *CT_DLbls
	DropLines  *CT_ChartLines
	GapDepth   *CT_GapAmount
	AxId       []*CT_UnsignedInt
	ExtLst     *CT_ExtensionList
}

func NewCT_Line3DChart() *CT_Line3DChart {
	ret := &CT_Line3DChart{}
	ret.Grouping = NewCT_Grouping()
	return ret
}

func (m *CT_Line3DChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	segrouping := xml.StartElement{Name: xml.Name{Local: "c:grouping"}}
	e.EncodeElement(m.Grouping, segrouping)
	if m.VaryColors != nil {
		sevaryColors := xml.StartElement{Name: xml.Name{Local: "c:varyColors"}}
		e.EncodeElement(m.VaryColors, sevaryColors)
	}
	if m.Ser != nil {
		seser := xml.StartElement{Name: xml.Name{Local: "c:ser"}}
		for _, c := range m.Ser {
			e.EncodeElement(c, seser)
		}
	}
	if m.DLbls != nil {
		sedLbls := xml.StartElement{Name: xml.Name{Local: "c:dLbls"}}
		e.EncodeElement(m.DLbls, sedLbls)
	}
	if m.DropLines != nil {
		sedropLines := xml.StartElement{Name: xml.Name{Local: "c:dropLines"}}
		e.EncodeElement(m.DropLines, sedropLines)
	}
	if m.GapDepth != nil {
		segapDepth := xml.StartElement{Name: xml.Name{Local: "c:gapDepth"}}
		e.EncodeElement(m.GapDepth, segapDepth)
	}
	seaxId := xml.StartElement{Name: xml.Name{Local: "c:axId"}}
	for _, c := range m.AxId {
		e.EncodeElement(c, seaxId)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Line3DChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Grouping = NewCT_Grouping()
lCT_Line3DChart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "grouping"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "grouping"}:
				if err := d.DecodeElement(m.Grouping, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "varyColors"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "varyColors"}:
				m.VaryColors = NewCT_Boolean()
				if err := d.DecodeElement(m.VaryColors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "ser"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "ser"}:
				tmp := NewCT_LineSer()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ser = append(m.Ser, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dLbls"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dLbls"}:
				m.DLbls = NewCT_DLbls()
				if err := d.DecodeElement(m.DLbls, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dropLines"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dropLines"}:
				m.DropLines = NewCT_ChartLines()
				if err := d.DecodeElement(m.DropLines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "gapDepth"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "gapDepth"}:
				m.GapDepth = NewCT_GapAmount()
				if err := d.DecodeElement(m.GapDepth, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "axId"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "axId"}:
				tmp := NewCT_UnsignedInt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AxId = append(m.AxId, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Line3DChart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Line3DChart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Line3DChart and its children
func (m *CT_Line3DChart) Validate() error {
	return m.ValidateWithPath("CT_Line3DChart")
}

// ValidateWithPath validates the CT_Line3DChart and its children, prefixing error messages with path
func (m *CT_Line3DChart) ValidateWithPath(path string) error {
	if err := m.Grouping.ValidateWithPath(path + "/Grouping"); err != nil {
		return err
	}
	if m.VaryColors != nil {
		if err := m.VaryColors.ValidateWithPath(path + "/VaryColors"); err != nil {
			return err
		}
	}
	for i, v := range m.Ser {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ser[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.DLbls != nil {
		if err := m.DLbls.ValidateWithPath(path + "/DLbls"); err != nil {
			return err
		}
	}
	if m.DropLines != nil {
		if err := m.DropLines.ValidateWithPath(path + "/DropLines"); err != nil {
			return err
		}
	}
	if m.GapDepth != nil {
		if err := m.GapDepth.ValidateWithPath(path + "/GapDepth"); err != nil {
			return err
		}
	}
	for i, v := range m.AxId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AxId[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
