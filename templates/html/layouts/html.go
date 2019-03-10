package layouts

import (
	"github.com/gnanakeethan/sleeker/templates/html/layouts/body"
	"github.com/gnanakeethan/sleeker/templates/html/layouts/head"
)

type HTML struct {
	Head HeadContainer
	Body BodyContainer
}

type HeadContainer struct {
	Title   string
	Metas   []*head.MetaContainer
	Styles  []*head.StyleContainer
	Scripts []*head.ScriptContainer
}

type BodyContainer struct {
	Scripts  []*body.ScriptContainer
	Contents string
}
