package builder

import (
	"bytes"
	"html/template"
	"strconv"
	"sync/atomic"

	"github.com/covrom/bulmactl/assets"
)

type DocumentView struct {
	gid       uint32
	name, tpl string
	items     []View
}

func NewDocumentView(header string) *DocumentView {
	d := &DocumentView{
		gid:  0,
		name: header,
		tpl:  assets.BaseTemplate,
	}
	if header == "" {
		return d
	}
	d.items = []View{
		&TextView{
			Text: header,
			Type: TextViewH1,
		},
	}
	return d
}

func (v *DocumentView) Parent() ViewGroup {
	return nil
}

func (v *DocumentView) NextId() string {
	i := atomic.AddUint32(&v.gid, 1)
	return "n" + strconv.Itoa(int(i))
}

func (v *DocumentView) Items() []View {
	return v.items
}

func (v *DocumentView) AddItem(a View) {
	v.items = append(v.items, a)
}

func (v *DocumentView) Render(ctx Context) (template.HTML, error) {
	t, err := template.New("documentview").Parse(v.tpl)
	if err != nil {
		return "", err
	}

	data := struct {
		Title string
		Items []template.HTML
	}{
		Title: v.name,
	}

	for _, vi := range v.items {
		vis, err := vi.Render(v)
		if err != nil {
			return "", err
		}
		data.Items = append(data.Items, vis)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return template.HTML(buf.String()), nil
}
