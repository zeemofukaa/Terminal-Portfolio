package sections

import "github.com/charmbracelet/lipgloss"

type ContactModel struct {
	Width  int
	Height int
}

func NewContactModel(w, h int) ContactModel {
	return ContactModel{
		Width:  w,
		Height: h,
	}
}

var contactColor = lipgloss.Color("#057059")

func (m ContactModel) View(height int) string {
	style := lipgloss.NewStyle().
		Padding(1, 2).
		Foreground(contactColor)

	return style.Render(`$ contact --list

email     : alamzoya2004@gmail.com
github    : zeemofukaa
itch.io   : zeemofukaa.itch.io
linkedin  : zoya-alam

status    : available for opportunities`)
}
