// Copyright 2016 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parse

import (
	"bytes"
	"fmt"
	"strings"
)

// Markdown saves 'Proto' to markdown documentation.
// lopts are a slice of language options (C++, Java, Python, Go, Ruby, C#).
func (p *Proto) Markdown(title, fpath string, lopts ...string) error {
	p.Sort()

	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("### %s\n\n\n", title))

	for _, msg := range p.Messages {
		buf.WriteString(fmt.Sprintf("##### message `%s`\n\n", msg.Name))
		if msg.Description != "" {
			buf.WriteString(msg.Description)
			buf.WriteString("\n\n")
		}
		hd1 := "| Field | Description | Type |"
		hd2 := "| ----- | ----------- | ---- |"
		for _, lopt := range lopts {
			hd1 += fmt.Sprintf(" %s |", lopt)
			ds := strings.Repeat("-", len(lopt))
			if len(ds) < 3 {
				ds = "---"
			}
			hd2 += fmt.Sprintf(" %s |", ds)
		}
		buf.WriteString(hd1 + "\n")
		buf.WriteString(hd2 + "\n")
		for _, elem := range msg.Fields {
			ts := elem.ProtoType.String()
			if elem.UserDefinedProtoType != "" {
				ts = elem.UserDefinedProtoType
			}
			if elem.Repeated {
				ts = "[]" + ts
			}
			line := fmt.Sprintf("| %s | %s | %s |", elem.Name, elem.Description, ts)
			for _, lopt := range lopts {
				if elem.UserDefinedProtoType != "" {
					line += " |"
					continue
				}
				formatSt := " %s |"
				if elem.Repeated {
					formatSt = " []%s |"
				}
				switch lopt {
				case "C++":
					line += fmt.Sprintf(formatSt, elem.ProtoType.Cpp())
				case "Java":
					line += fmt.Sprintf(formatSt, elem.ProtoType.Java())
				case "Python":
					line += fmt.Sprintf(formatSt, elem.ProtoType.Python())
				case "Go":
					line += fmt.Sprintf(formatSt, elem.ProtoType.Go())
				case "Ruby":
					line += fmt.Sprintf(formatSt, elem.ProtoType.Ruby())
				case "C#":
					line += fmt.Sprintf(formatSt, elem.ProtoType.Csharp())
				default:
					return fmt.Errorf("%q is unknown (must be C++, Java, Python, Go, Ruby, C#)", lopt)
				}
			}
			buf.WriteString(line + "\n")
		}
		buf.WriteString("\n\n<br>\n\n")
	}

	for _, svs := range p.Services {
		buf.WriteString(fmt.Sprintf("##### service `%s`\n\n", svs.Name))
		if svs.Description != "" {
			buf.WriteString(svs.Description)
			buf.WriteString("\n\n")
		}
		hd1 := "| Method | Request Type | Response Type | Description |"
		hd2 := "| ------ | ------------ | ------------- | ----------- |"
		buf.WriteString(hd1 + "\n")
		buf.WriteString(hd2 + "\n")

		for _, elem := range svs.Methods {
			line := fmt.Sprintf("| %s | `%s` | `%s` | %s |", elem.Name, elem.RequestType, elem.ResponseType, elem.Description)
			buf.WriteString(line + "\n")
		}
		buf.WriteString("\n\n<br>\n\n")
	}

	return toFile(buf.String(), fpath)
}