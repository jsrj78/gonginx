package dumper

var (
	//NoIndentStyle default style
	NoIndentStyle = &Style{
		SortDirectives: false,
		StartIndent:    0,
		Indent:         0,
	}
)

//Style dumping style
type Style struct {
	SortDirectives    bool
	SpaceBeforeBlocks bool
	StartIndent       int
	Indent            int
}

//NewStyle create new style
func NewStyle() *Style {
	style := &Style{
		SortDirectives: false,
		StartIndent:    0,
		Indent:         4,
	}
	return style
}

//Iterate interate the indentation for sub blocks
func (s *Style) Iterate() *Style {
	newStyle := &Style{
		SortDirectives:    s.SortDirectives,
		SpaceBeforeBlocks: s.SpaceBeforeBlocks,
		StartIndent:       s.StartIndent + s.Indent,
		Indent:            s.Indent,
	}
	return newStyle
}