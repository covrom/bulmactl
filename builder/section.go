package builder

import (
	"github.com/covrom/bulmactl/assets"
)

type SectionViewClass int

const (
	SectionViewSection SectionViewClass = iota
	SectionViewHero
	SectionViewHeroPrimary
	SectionViewHeroInfo
	SectionViewHeroSuccess
	SectionViewHeroWarning
	SectionViewHeroDanger
	SectionViewHeroLight
	SectionViewHeroDark
)

type SectionViewSize int

const (
	SectionViewSizeDefault SectionViewSize = iota
	SectionViewSizeMedium
	SectionViewSizeLarge
	SectionViewSizeFull
)

type SectionView struct {
	Id     string
	wantId bool
	parent ViewGroup
	tpl    string
	class  SectionViewClass
	isBold bool
	size   SectionViewSize
	items  []View
	header View
	body   View
	footer View
}

func NewSectionView(setId bool, parent ViewGroup) *SectionView {
	return &SectionView{
		Id:     "",
		wantId: setId,
		parent: parent,
		tpl:    assets.SectionTpl,
	}
}

// TODO: completing