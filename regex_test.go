package pauthor

import (
	"strconv"
	"testing"

	"github.com/ImSingee/tt"
)

func author(s string) [][]string {
	return _regex.FindAllStringSubmatch(s, -1)
}

func oneAuthor(s string) []string {
	as := author(s)
	if len(as) != 1 {
		panic("Want 1 author but got " + strconv.Itoa(len(as)))
	}

	return as[0]
}

func TestRegex(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want [3]string
	}{
		{
			name: `all`,
			in:   `Jon Schlinkert <jon.schlinkert@sellside.com> (https://github.com/jonschlinkert)`,
			want: [3]string{"Jon Schlinkert", "jon.schlinkert@sellside.com", "https://github.com/jonschlinkert"},
		},
		{
			name: `name and url`,
			in:   `Jon Schlinkert (https://github.com/jonschlinkert)`,
			want: [3]string{"Jon Schlinkert", "", "https://github.com/jonschlinkert"},
		},
		{
			name: `name and url`,
			in:   `Jon Schlinkert(https://github.com/jonschlinkert)`,
			want: [3]string{"Jon Schlinkert", "", "https://github.com/jonschlinkert"},
		},
		{
			name: `email and url`,
			in:   `<jon.schlinkert@sellside.com> (https://github.com/jonschlinkert)`,
			want: [3]string{"", "jon.schlinkert@sellside.com", "https://github.com/jonschlinkert"},
		},
		{
			name: `name and email`,
			in:   `Jon Schlinkert <jon.schlinkert@sellside.com>`,
			want: [3]string{"Jon Schlinkert", "jon.schlinkert@sellside.com", ""},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := oneAuthor(c.in)

			tt.AssertEqual(t, 4, len(got))
			tt.AssertEqual(t, c.want[0], got[1])
			tt.AssertEqual(t, c.want[1], got[2])
			tt.AssertEqual(t, c.want[2], got[3])
		})
	}
}
