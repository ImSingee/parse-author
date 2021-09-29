package pauthor

import (
	"github.com/ImSingee/tt"
	"testing"
)

func TestParser(t *testing.T) {
	cases := []struct {
		name      string
		in        string
		want      [3]string
		willError bool
	}{
		{
			name:      `empty`,
			in:        "",
			willError: true,
		},
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
			got, err := Parse(c.in)

			if !c.willError {
				tt.AssertIsNil(t, err)

				tt.AssertEqual(t, c.want[0], got.Name)
				tt.AssertEqual(t, c.want[1], got.Email)
				tt.AssertEqual(t, c.want[2], got.Url)
			} else {
				tt.AssertIsNotNil(t, err)
			}
		})
	}
}