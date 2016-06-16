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

import "testing"

func TestMarkdown(t *testing.T) {
	proto, err := ReadDir("testdata", "")
	if err != nil {
		t.Fatal(err)
	}
	if txt, err := proto.Markdown("this is test...", "etcdserverpb", []ParseOption{ParseService, ParseMessage}, "Go", "Java", "Python", "C++"); err != nil {
		t.Fatal(err)
	} else {
		err = toFile(txt, "testdata/README.md")
		if err != nil {
			t.Fatal(err)
		}
	}
}
