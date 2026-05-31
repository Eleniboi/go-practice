package animation

import "strings"

type Animation struct{ frames []string }

func NewAnimation(_ string, n int) *Animation {
	return &Animation{make([]string, n)}
}

func (a *Animation) gen() {
	for i := range a.frames {
		ch := string([]rune{'-', '\\', '|', '/'}[i%4])
		a.frames[i] = strings.Repeat(strings.Repeat(ch, 10)+"\n", 10)
	}
}

func (a *Animation) GenerateSpinFrames()   { a.gen() }
func (a *Animation) GenerateWaveFrames()   { a.gen() }
func (a *Animation) GenerateZoomFrames()   { a.gen() }
func (a *Animation) GetFrame(i int) string { return a.frames[i%len(a.frames)] }
func (a *Animation) Play() string          { return "---\n" + strings.Join(a.frames, "---\n") }
