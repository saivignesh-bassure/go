package foodchain

import (
	"fmt"
	"strings"
	"testing"
)

var text = []string{``,

	`I know an old lady who swallowed a fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a spider.
It wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a bird.
How absurd to swallow a bird!
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a cat.
Imagine that, to swallow a cat!
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a dog.
What a hog, to swallow a dog!
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a goat.
Just opened her throat and swallowed a goat!
She swallowed the goat to catch the dog.
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a cow.
I don't know how she swallowed a cow!
She swallowed the cow to catch the goat.
She swallowed the goat to catch the dog.
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a horse.
She's dead, of course!`,
}

// diff compares two multi-line strings and returns a helpful comment
func diff(got, want string) string {
	g := strings.Split(got, "\n")
	w := strings.Split(want, "\n")
	for i := 0; ; i++ {
		switch {
		case i < len(g) && i < len(w):
			if g[i] == w[i] {
				continue
			}
			return fmt.Sprintf("-- first difference in line %d:\n"+
				"-- got : %q\n-- want: %q\n", i+1, g[i], w[i])
		case i < len(g):
			return fmt.Sprintf("-- got %d extra lines after line %d:\n"+
				"-- first extra line: %q\n", len(g)-len(w), i, g[i])
		case i < len(w):
			return fmt.Sprintf("-- got %d correct lines, want %d more lines:\n"+
				"-- want next: %q\n", i, len(w)-i, w[i])
		default:
			return "no differences found"
		}
	}
}

func TestVerse(t *testing.T) {
	for v := 1; v <= 8; v++ {
		t.Run(fmt.Sprintf("verse %d", v), func(t *testing.T) {
			if got := Verse(v); got != text[v] {
				t.Fatalf("Verse(%d)\ngot:%s\nwant:%s\nhelp: %s", v, got, text[v], diff(got, text[v]))
			}
		})
	}
}

func TestVerses(t *testing.T) {
	if got, want := Verses(1, 3), strings.Join(text[1:4], "\n\n"); got != want {
		t.Fatalf("Verses(1, 3) =\n%s\n  want:\n%s\n%s", got, want, diff(got, want))
	}
}

func TestSong(t *testing.T) {
	if got, want := Song(), strings.Join(text[1:], "\n\n"); got != want {
		t.Fatalf("Song() =\n%s\n  want:\n%s\n%s", got, want, diff(got, want))
	}
}

func BenchmarkSong(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Song()
	}
}
