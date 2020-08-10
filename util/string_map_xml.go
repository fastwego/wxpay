// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"bytes"
	"encoding/xml"
	"io"
)

type stringMap map[string]string

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

// MarshalXML marshals the map to XML, with each key in the map being a
// tag and it's corresponding value being it's contents.
func (m stringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

// UnmarshalXML unmarshals the XML into a map of string to strings,
// creating a key in the map for each tag and setting it's value to the
// tags contents.
//
// The fact this function is on the pointer of stringMap is important, so that
// if m is nil it can be initialized, which is often the case if m is
// nested in another xml structurel. This is also why the first thing done
// on the first line is initialize it.
func (m *stringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = stringMap{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

func Map2XML(kvs map[string]string) (text []byte, err error) {
	text, err = xml.Marshal(stringMap(kvs))
	if err != nil {
		return
	}
	text = bytes.ReplaceAll(text, []byte(`<stringMap>`), []byte(`<xml>`))
	text = bytes.ReplaceAll(text, []byte(`</stringMap>`), []byte(`</xml>`))

	return
}

func XML2Map(text []byte) (result map[string]string, err error) {
	err = xml.Unmarshal(text, (*stringMap)(&result))
	if err != nil {
		return
	}
	return
}
