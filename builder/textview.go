package builder

import (
	"html"
	"html/template"
)

type TextViewType int

const (
	TextViewRegular TextViewType = iota
	TextViewP
	TextViewH1
	TextViewH2
	TextViewH3
	TextViewH4
	TextViewH5
	TextViewH6
)

type TextView struct {
	Id     string
	wantId bool
	parent ViewGroup
	Text   string
	Type   TextViewType
}

func NewTextView(text string, typ TextViewType, setId bool, parent ViewGroup) *TextView {
	return &TextView{
		Id:     "",
		wantId: setId,
		parent: parent,
		Text:   text,
		Type:   typ,
	}
}

func (v *TextView) Parent() ViewGroup {
	return v.parent
}

func (v *TextView) Render(ctx Context) (template.HTML, error) {
	if v.wantId && v.Id == "" {
		v.Id = ctx.NextId()
	}
	pr := "span"
	switch v.Type {
	case TextViewP:
		pr = "p"
	case TextViewH1:
		pr = "h1"
	case TextViewH2:
		pr = "h2"
	case TextViewH3:
		pr = "h3"
	case TextViewH4:
		pr = "h4"
	case TextViewH5:
		pr = "h5"
	case TextViewH6:
		pr = "h6"
	}
	if v.Id != "" {
		return template.HTML(
			"<" + pr + " id=\"" + v.Id + "\">" +
				html.EscapeString(v.Text) +
				"</" + pr + ">"), nil
	}
	return template.HTML(
		"<" + pr + ">" +
			html.EscapeString(v.Text) +
			"</" + pr + ">"), nil
}
