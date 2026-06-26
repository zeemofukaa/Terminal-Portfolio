package sections

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	cardHeaderActiveBg = lipgloss.Color("#1FD0B5")
	cardHeaderActiveFg = lipgloss.Color("#000000")

	cardHeaderInactiveBg = lipgloss.Color("#066E60")
	cardHeaderInactiveFg = lipgloss.Color("#B8FFF4")

	cardBorder = lipgloss.Color("#045846")

	shortDescColor = lipgloss.Color("#067363")
	descColor      = lipgloss.Color("#00A08B")
	tagColor       = lipgloss.Color("#05806C")
)

type Project struct {
	Name      string
	Tags      []string
	ShortDesc string
	Desc      string
	Link      string
	LinkLabel string
}

var projects = []Project{
	{
		Name:      "Omori Weather App",
		Tags:      []string{"electron", "js", "css", "openweather api"},
		ShortDesc: `A desktop weather app that feels like an RPG`,
		Desc: `Styled after the RPG game OMORI — White Space and Black Space
themes that shift based on time of day. Pixel-art walking
animations, adaptive background music, sound effects for every
interaction, and a settings menu for audio control.

Real-time data via OpenWeatherMap API. Packaged as a standalone
desktop app, downloadable for any OS from GitHub.`,
		Link:      "github.com/zeemofukaa/omori-weather",
		LinkLabel: "→ download on github",
	},
	{
		Name:      "Axotchi",
		Tags:      []string{"godot 4", "gdscript", "itch.io"},
		ShortDesc: `A cozy desktop virtual pet starring a pink axolotl`,
		Desc: `Feeding, sleeping, playing — full stat system with hunger,
energy, and happiness decay. Day/night UI transitions,
mood-responsive background tinting, dialogue bubbles,
offline progression, day streaks, and background music
with SFX ducking.

Shipped v1.0 and v1.1 on itch.io. Open source on GitHub.`,
		Link:      "zeemofukaa.itch.io/axotchi",
		LinkLabel: "→ play on itch.io",
	},
	{
		Name:      "Paimon Restaurant",
		Tags:      []string{"node.js", "express", "mongodb", "jwt"},
		ShortDesc: `A full-stack restaurant website inspired by Genshin Impact`,
		Desc: `JWT authentication, persistent cart, order history, and
dynamic menu rendering across Starters, Mains, Desserts,
and Drinks. Editorial-grid homepage with Genshin artwork.

Built as a real full-stack project with backend API,
database, and auth — not just a static mockup.`,
		Link:      "github.com/zeemofukaa/paimon-restaurant",
		LinkLabel: "→ view on github",
	},
	{
		Name:      "Terminal Portfolio",
		Tags:      []string{"go", "bubbletea", "lipgloss"},
		ShortDesc: "A retro terminal-inspired portfolio built in Go.",
		Desc: `Interactive TUI portfolio inspired by command-line
interfaces. Keyboard navigation, expandable project cards,
custom layouts, and a handcrafted retro aesthetic built
using Bubble Tea and Lipgloss.`,

		Link:      "github.com/zeemofukaa/portfolio",
		LinkLabel: "→ view on github",
	},
}

type ProjectsModel struct {
	Width    int
	Height   int
	cursor   int
	expanded bool
}

func NewProjectsModel(w, h int) ProjectsModel {
	return ProjectsModel{Width: w, Height: h}
}

func (m ProjectsModel) Update(msg tea.Msg) (ProjectsModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "a":
			if m.cursor%2 == 1 {
				m.cursor--
			}

		case "d":
			if m.cursor%2 == 0 && m.cursor+1 < len(projects) {
				m.cursor++
			}

		case "w":
			if m.cursor >= 2 {
				m.cursor -= 2
			}

		case "s":
			if m.cursor+2 < len(projects) {
				m.cursor += 2
			}

		case "enter":
			m.expanded = !m.expanded

		case "esc":
			m.expanded = false
		}
	}
	return m, nil
}

func renderClosedCard(p Project, selected bool) string {
	bg := cardHeaderInactiveBg
	fg := cardHeaderInactiveFg

	if selected {
		bg = cardHeaderActiveBg
		fg = cardHeaderActiveFg
	}

	header := lipgloss.NewStyle().
		Background(bg).
		Foreground(fg).
		Bold(true).
		Width(48).
		Padding(1, 1).
		Render("◇ " + strings.ToUpper(p.Name))

	shortDesc := lipgloss.NewStyle().
		Foreground(shortDescColor).
		Padding(1, 2).
		Width(44).
		Render(p.ShortDesc)

	return lipgloss.NewStyle().
		Width(48).
		Border(lipgloss.NormalBorder()).
		BorderForeground(cardBorder).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				header,
				shortDesc,
			),
		)
}

func renderOpenCard(p Project, selected bool) string {
	bg := cardHeaderInactiveBg
	fg := cardHeaderInactiveFg

	if selected {
		bg = cardHeaderActiveBg
		fg = cardHeaderActiveFg
	}
	const contentWidth = 44

	header := lipgloss.NewStyle().
		Background(bg).
		Foreground(fg).
		Bold(true).
		Width(48).
		Padding(1, 1).
		Render("◆ " + strings.ToUpper(p.Name))

	shortDesc := lipgloss.NewStyle().
		Foreground(shortDescColor).
		PaddingLeft(2).
		Width(contentWidth).
		Render(p.ShortDesc)

	desc := lipgloss.NewStyle().
		Foreground(descColor).
		PaddingLeft(2).
		Width(contentWidth).
		Render(p.Desc)

	tags := lipgloss.NewStyle().
		Foreground(tagColor).
		PaddingLeft(2).
		Width(contentWidth).
		Render("• " + strings.Join(p.Tags, " • "))

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		"",
		shortDesc,
		"",
		desc,
		"",
		tags,
	)

	return lipgloss.NewStyle().
		Width(48).
		Height(13).
		Border(lipgloss.NormalBorder()).
		BorderForeground(cardBorder).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				header,
				content,
			),
		)
}

func (m ProjectsModel) View(height int) string {
	cards := make([]string, len(projects))

	for i, p := range projects {

		selected := i == m.cursor

		if selected && m.expanded {
			cards[i] = renderOpenCard(p, true)
		} else {
			cards[i] = renderClosedCard(p, selected)
		}
	}

	row1 := lipgloss.JoinHorizontal(
		lipgloss.Top,
		cards[0],
		cards[1],
	)

	row2 := lipgloss.JoinHorizontal(
		lipgloss.Top,
		cards[2],
		// " ",
		cards[3],
	)

	navHint := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1FD0B5")).
		Width(100).
		Align(lipgloss.Center).
		Render("[ wasd to navigate  ·  enter to expand/collapse ]")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		row1,
		"",
		row2,
	)

	content = lipgloss.NewStyle().
		Height(height - 2).
		Render(content)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		content,
		navHint,
	)
}
