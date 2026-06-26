package sections

import (
	// "math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SkillsModel struct {
	Width  int
	Height int

	phase int
}

func NewSkillsModel(w, h int) SkillsModel {
	return SkillsModel{
		Width:  w,
		Height: h,
	}
}

var (
	skillBright = lipgloss.Color("#9FEBDE")
	skillMid    = lipgloss.Color("#58B6A5")
	skillDim    = lipgloss.Color("#2F7A70")

	categoryBg = lipgloss.Color("#21D4B3")
	categoryFg = lipgloss.Color("#000000")

	boxBorder = lipgloss.Color("#045846")
)

var skillColors = []lipgloss.Color{
	skillBright,
	skillMid,
	skillDim,
}

func skillText(text string, phase int) string {

	hash := 0

	for i, c := range text {
		hash += int(c) * (i + 1)
	}

	shade := (hash + phase) % 3

	color := skillBright

	switch shade {
	case 1:
		color = skillMid
	case 2:
		color = skillDim
	}

	return lipgloss.NewStyle().
		Foreground(color).
		Render(text)
}

func categoryPill(title string) string {
	return lipgloss.NewStyle().
		Background(categoryBg).
		Foreground(categoryFg).
		Bold(true).
		Padding(0, 2).
		Render(title)
}

type SkillsTickMsg time.Time

func skillsTick() tea.Cmd {
	return tea.Tick(
		600*time.Millisecond,
		func(t time.Time) tea.Msg {
			return SkillsTickMsg(t)
		},
	)
}

func (m SkillsModel) Init() tea.Cmd {
	return skillsTick()
}

func (m SkillsModel) Update(msg tea.Msg) (SkillsModel, tea.Cmd) {
	switch msg.(type) {

	case SkillsTickMsg:
		m.phase++
		return m, skillsTick()
	}

	return m, nil
}

func (m SkillsModel) View(height int) string {

	software := lipgloss.NewStyle().
		Width(44).
		Height(14).
		Render(
			strings.Join([]string{
				"",
				"                 " + skillText("React.js", m.phase),
				"",
				"      " + skillText("C/C++", m.phase),
				"              " + skillText("Javascript", m.phase),
				"",
				"            " + categoryPill("SOFTWARE DEV"),
				"",
				"                          " + skillText("Electron.js", m.phase),
				"",
				skillText("MongoDB", m.phase),
				"                 " + skillText("SQL", m.phase),
				"                         " + skillText("Python", m.phase),
				"       " + skillText("Node.js", m.phase),
			}, "\n"),
		)

	graphics := lipgloss.NewStyle().
		Width(36).
		Height(12).
		Render(
			strings.Join([]string{
				"                  " + skillText("Blender", m.phase),
				"",
				"      " + skillText("Godot", m.phase),
				"",
				"",
				"        " + categoryPill("GRAPHICS"),
				"",
				"",
				"    " + skillText("Adobe Illustrator", m.phase),
				"                              " + skillText("Figma", m.phase),
			}, "\n"),
		)

	tools := lipgloss.NewStyle().
		Width(62).
		Height(14).
		Render(
			strings.Join([]string{
				"            " + skillText("Git/Github", m.phase),
				"                        " + skillText("NumPy", m.phase),
				"                              " + skillText("Pandas", m.phase),
				"",
				skillText("MATLAB", m.phase),
				"                                          " + skillText("SciPy", m.phase),
				"",
				"                  " + categoryPill("TOOLS & ENGINEERING"),
				"",
				"       " + skillText("SolidWorks", m.phase),
				"",
				"                                        " + skillText("Matplotlib", m.phase),
				"",
				skillText("Google Colab", m.phase) +
					"               " +
					skillText("Jupyter Notebook", m.phase),
			}, "\n"),
		)

	topRow := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().
			MarginLeft(8).
			Render(software),
		"      ",
		lipgloss.NewStyle().
			MarginTop(7).
			Render(graphics),
	)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		topRow,
		"",
		lipgloss.NewStyle().
			MarginTop(1).
			MarginLeft(14).
			Render(tools),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		content,
		"",
		"",
	)
}
