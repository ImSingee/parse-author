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
			name:      `invalid`,
			in:        "()",
			willError: true,
		},
		{
			name: `all`,
			in:   `Jon Schlinkert <jon.schlinkert@sellside.com> (https://github.com/jonschlinkert)`,
			want: [3]string{"Jon Schlinkert", "jon.schlinkert@sellside.com", "https://github.com/jonschlinkert"},
		},
		{
			name: `only name`,
			in:   `Jon Schlinkert`,
			want: [3]string{"Jon Schlinkert", "", ""},
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

func TestAuthor_String(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{
			in:   `Jon Schlinkert <jon.schlinkert@sellside.com> (https://github.com/jonschlinkert)`,
			want: `Jon Schlinkert <jon.schlinkert@sellside.com> (https://github.com/jonschlinkert)`,
		},
		{
			in:   `Jon Schlinkert`,
			want: "Jon Schlinkert",
		},
		{
			in:   `Jon Schlinkert (https://github.com/jonschlinkert)`,
			want: "Jon Schlinkert (https://github.com/jonschlinkert)",
		},
		{
			in:   `Jon Schlinkert(https://github.com/jonschlinkert)`,
			want: `Jon Schlinkert (https://github.com/jonschlinkert)`,
		},
		{
			in:   `<jon.schlinkert@sellside.com> (https://github.com/jonschlinkert)`,
			want: `<jon.schlinkert@sellside.com> (https://github.com/jonschlinkert)`,
		},
		{
			in:   `Jon Schlinkert <jon.schlinkert@sellside.com>`,
			want: `Jon Schlinkert <jon.schlinkert@sellside.com>`,
		},
		{
			in:   ` Jon Schlinkert  <jon.schlinkert@sellside.com> `,
			want: `Jon Schlinkert <jon.schlinkert@sellside.com>`,
		},
		{
			in:   `Jon Schlinkert  <jon.schlinkert@sellside.com>(https://github.com/jonschlinkert)  `,
			want: `Jon Schlinkert <jon.schlinkert@sellside.com> (https://github.com/jonschlinkert)`,
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			got, err := Parse(c.in)
			tt.AssertIsNil(t, err)

			tt.AssertEqual(t, c.want, got.String())
		})
	}
}
