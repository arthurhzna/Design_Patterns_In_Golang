package main
package main

import (
  "fmt"
  "strings"
  "unicode"
)

type FormattedText struct {
  plainText  string
  capitalize []bool
}

func NewFormattedText(plainText string) *FormattedText {
  return &FormattedText{plainText,
    make([]bool, len(plainText))}
}

// type FormattedText struct {
// 	plainText  string   ---> "This is a brave new world"
// 	capitalize []bool ---> [false, false,....]
// }

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
	  c := f.plainText[i]
	  if f.capitalize[i] {
		sb.WriteRune(unicode.ToUpper(rune(c)))
	  } else {
		sb.WriteRune(rune(c))
	  }
	}
	return sb.String() // "This is a BRAVE new world"
}


func (f *FormattedText) Capitalize(start, end int) {
  for i := start; i <= end; i++ {
    f.capitalize[i] = true 
	// 0..9 = false
	// 10..15 = true
	// 16..24 = false
  }
}

type TextRange struct {
  Start, End int
  Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
  return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
  plainText string
  formatting []*TextRange
}

func (b *BetterFormattedText) String() string {
  sb := strings.Builder{}

  for i := 0; i < len(b.plainText); i++ {
    c := b.plainText[i]
    for _, r := range b.formatting {
      if r.Covers(i) && r.Capitalize {
        c = uint8(unicode.ToUpper(rune(c)))
      }
    }
    sb.WriteRune(rune(c))
  }

  return sb.String()
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
  return &BetterFormattedText{plainText: plainText}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
  r := &TextRange{start, end, false, false, false}
  b.formatting = append(b.formatting, r)
  return r
} 



func main() {
  text := "This is a brave new world"

  ft := NewFormattedText(text)
  ft.Capitalize(10, 15) // brave  <--- normal, not flyweight
  fmt.Println(ft.String())

  bft := NewBetterFormattedText(text)  //<-- flyweight
  bft.Range(16, 19).Capitalize = true // new 
  fmt.Println(bft.String())
}

// Flyweight is used to save memory by storing shared 
// text data once and keeping only small formatting rules, 
// such as ranges, instead of storing formatting for every character.
// In this example, BetterFormattedText is more efficient than FormattedText 
// because it stores only the formatted ranges, not a boolean flag for the entire text.

// In this example, the text "This is a brave new world" is formatted in two different ways.
// FormattedText stores a capitalization flag for every character, while BetterFormattedText stores only the 
// formatting ranges, such as making "new" uppercase. This shows how the Flyweight pattern saves memory by keeping only 
// the necessary formatting information instead of storing data for the entire text.