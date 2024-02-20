package genreq

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var ErrNoColor = errors.New("not a color")

// Color enumerates the set of valid color constants we can use to describe images
type Color string

const (
	ColorRed          Color = "red"
	ColorRedOrange    Color = "red-orange"
	ColorOrange       Color = "orange"
	ColorYellowOrange Color = "yellow-orange"
	ColorYellow       Color = "yellow"
	ColorChartreuse   Color = "chartreuse"
	ColorGreen        Color = "green"
	ColorCyan         Color = "cyan"
	ColorSkyBlue      Color = "sky-blue"
	ColorBlue         Color = "blue"
	ColorIndigo       Color = "indigo"
	ColorPurple       Color = "purple"
	ColorMagenta      Color = "magenta"
)

var Colors = []Color{
	ColorRed,
	ColorRedOrange,
	ColorOrange,
	ColorYellowOrange,
	ColorYellow,
	ColorChartreuse,
	ColorGreen,
	ColorCyan,
	ColorSkyBlue,
	ColorBlue,
	ColorIndigo,
	ColorPurple,
	ColorMagenta,
}

func (c Color) GetComplement() Color {
	switch c {
	case ColorRed:
		return ColorGreen
	case ColorRedOrange:
		return ColorCyan
	case ColorOrange:
		return ColorSkyBlue
	case ColorYellowOrange:
		return ColorBlue
	case ColorYellow:
		return ColorIndigo
	case ColorChartreuse:
		return ColorMagenta
	case ColorGreen:
		return ColorRed
	case ColorCyan:
		return ColorRedOrange
	case ColorSkyBlue:
		return ColorOrange
	case ColorBlue:
		return ColorYellowOrange
	case ColorIndigo:
		return ColorYellowOrange
	case ColorPurple:
		return ColorYellow
	case ColorMagenta:
		return ColorChartreuse
	}
	return ColorRed
}

// MatchColor takes a string and returns a canonical Color constant representing the
// color named at the start of that string, followed by the remainder of the string
// after whitespace; or ErrNoColor if no color name is detected. Matching on color name
// is case-insensitive.
//
// Examples:
// - MatchColor("red") => (ColorRed, "", nil)
// - MatchColor("red shoes") => (ColorRed, "shoes" nil)
// - MatchColor("yellow-orange") => (ColorYellowOrange, "", nil)
// - MatchColor("orange-yellow") => (ColorYellowOrange, "", nil)
// - MatchColor("Green Muscadine grapes") => (ColorGreen, "Muscadine grapes", nil)
// - MatchColor("a ripe orange on a tree") => (_, _, ErrNoColor)
func MatchColor(s string) (Color, string, error) {
	m := colorRegexp.FindStringSubmatch(s)
	if m == nil {
		return ColorRed, "", ErrNoColor
	}
	key := resolveLookupKey(m[1], m[2])
	color, ok := colorLookup[key]
	if !ok {
		return ColorRed, "", ErrNoColor
	}
	remainderPos := len(m[0])
	for remainderPos < len(s) && s[remainderPos] == ' ' {
		remainderPos++
	}
	return color, s[remainderPos:], nil
}

func resolveLookupKey(lhs string, rhs string) string {
	lhs = strings.ToLower(lhs)
	rhs = strings.ToLower(rhs)
	if rhs == "" || rhs == lhs {
		return lhs
	}
	if lhs < rhs {
		return fmt.Sprintf("%s-%s", lhs, rhs)
	}
	return fmt.Sprintf("%s-%s", rhs, lhs)
}

var colorRegexp *regexp.Regexp
var colorLookup map[string]Color

func init() {
	colorRegexp = makeColorRegexp()
	colorLookup = makeColorLookup()
}

func makeColorRegexp() *regexp.Regexp {
	// Identify each atomic color word from our set of valid Color constants, either
	// '%s' or '%s-%s'
	slugSet := make(map[string]struct{})
	for _, color := range Colors {
		s := string(color)
		hyphenPos := strings.IndexRune(s, '-')
		if hyphenPos > 0 && hyphenPos < len(s)-1 {
			slugSet[s[:hyphenPos]] = struct{}{}
			slugSet[s[hyphenPos+1:]] = struct{}{}
		} else {
			slugSet[s] = struct{}{}
		}
	}

	// Convert our set of slugs to a sorted slice, for the sake of determinism
	slugs := make([]string, 0, len(slugSet))
	for slug := range slugSet {
		slugs = append(slugs, slug)
	}
	sort.Strings(slugs)

	// Prepare a regex pattern that will match on:
	// - group 1 (required): any slug value
	// - group 2 (optional): any slug value, delimited with a space, slash, or hyphen
	slugsGroup := fmt.Sprintf("(%s)", strings.Join(slugs, "|"))
	delimChars := "[-/ ]"
	pattern := fmt.Sprintf("(?i)^%s(?:%s%s)?", slugsGroup, delimChars, slugsGroup)
	return regexp.MustCompile(pattern)
}

func makeColorLookup() map[string]Color {
	lookup := make(map[string]Color)
	for _, color := range Colors {
		var lhs, rhs string
		s := string(color)
		hyphenPos := strings.IndexRune(s, '-')
		if hyphenPos > 0 && hyphenPos < len(s)-1 {
			lhs = s[:hyphenPos]
			rhs = s[hyphenPos+1:]
		} else {
			lhs = s
		}
		key := resolveLookupKey(lhs, rhs)
		lookup[key] = color
	}
	return lookup
}
