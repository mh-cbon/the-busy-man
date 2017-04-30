package wish

import "strings"

// Wish tells about a wish and it shades,
// wish:shade1+shade2
type Wish struct {
	Plugin string
	Shades *StringSlice
}

// Parse s and make a new Wish.
func Parse(s string) (Wish, error) {
	t := strings.Split(s, ":")
	x := strings.Split(t[1], "+") // can do better

	return Wish{
		Plugin: t[0],
		Shades: NewStringSlice().Push(x...),
	}, nil
}

// GetID implements lister
func (w *Wish) GetID() string {
	return w.Plugin
}

//FilterByPlugin filters wishes by plugin p
func FilterByPlugin(p ...string) func(Wish) bool {
	return func(w Wish) bool {
		for _, u := range p {
			if u == w.Plugin {
				return true
			}
		}
		return false
	}
}

//FilterByShade filters strings by value s
func FilterByShade(s string) func(string) bool {
	return func(w string) bool {
		return w == s
	}
}

//go:generate lister string_gen.go string:StringSlice
//go:generate lister wishs_gen.go Wish:*Wishes
