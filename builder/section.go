package builder

import (
	"github.com/covrom/bulmactl/assets"
)

type SectionView struct {
	Id     string
	wantId bool
	parent ViewGroup
	tpl    string
	Items  []View
}

func NewSectionView(setId bool, parent ViewGroup) *SectionView {
	return &SectionView{
		Id:     "",
		wantId: setId,
		parent: parent,
		tpl:    assets.SectionTpl,
	}
}

// TODO: Bulma Hero