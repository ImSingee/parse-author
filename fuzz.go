//go:build gofuzz
// +build gofuzz

package pauthor

func Fuzz(data []byte) int {
	d := string(data)

	_, err := Parse(d)
	if err == nil {
		return 1
	}

	return 0
}
