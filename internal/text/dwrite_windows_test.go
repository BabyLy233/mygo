//go:build windows

package text

import "testing"

// Tabs reach their stops, those starting a line too, though DirectWrite
// draws no run of them: a line indented by tabs keeps its indent.
func TestDWriteTabsReachTheirStops(t *testing.T) {
	s := newSystem()
	style := Style{Family: "monospace", Size: 13}
	// after is where the last glyph of text starts.
	after := func(text string) float32 {
		l := s.Layout(Params{Text: text, Style: style})
		g := l.Lines[0].Glyphs
		return g[len(g)-1].X
	}
	stop := after("\tx")
	if stop <= after("x") {
		t.Fatalf("x after a tab is at %g", stop)
	}
	if x := after("ab\tx"); x != stop {
		t.Errorf("x after ab and a tab is at %g, after a tab at %g", x, stop)
	}
	if x := after("\t\tx"); x != 2*stop {
		t.Errorf("x after two tabs is at %g, after one at %g", x, stop)
	}
	if w := s.Layout(Params{Text: "\t", Style: style, KeepSpaces: true}).Lines[0].Width; w != stop {
		t.Errorf("a tab is %g wide, its stop at %g", w, stop)
	}
}
