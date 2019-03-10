package sleeker

import "github.com/gnanakeethan/sleeker/templates/html/layouts"

type Sleeker struct {
	layouts.HTML
}

func (sleeker *Sleeker) SetTitle(title string) {
	sleeker.HTML.Head.Title = title
}

func (sleeker *Sleeker) Render() string {
	return layouts.RenderHtml(sleeker.HTML)
}
