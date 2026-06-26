package ui

import (
	"strings"

	"zeemofukaa/internal/sections"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Tab int

const (
	TabAbout Tab = iota
	TabProjects
	TabSkills
	TabContact
	TabEasterEgg
	tabCount
	navbarWidth  = 100 // ~974px
	navbarHeight = 1

	mainWidth  = 100 // ~974px
	mainHeight = 30  // ~681px

	hintsWidth  = 100
	hintsHeight = 1
)

var tabNames = []string{"about", "projects", "skills", "contact", "???"}

var (
	frameBorder = lipgloss.Color("#045846")
	frameBright = lipgloss.Color("#0ADCAE")
)

type Model struct {
	width     int
	height    int
	activeTab Tab

	aboutModel    sections.AboutModel
	projectsModel sections.ProjectsModel
	skillsModel   sections.SkillsModel
	contactModel  sections.ContactModel
	easterModel   sections.EasterModel
}

func NewModel() Model {
	return NewModelWithSize(120, 40)
}

func NewModelWithSize(w, h int) Model {
	m := Model{
		width:     w,
		height:    h,
		activeTab: TabAbout,
	}
	m.aboutModel = sections.NewAboutModel(w, h)
	m.projectsModel = sections.NewProjectsModel(w, h)
	m.skillsModel = sections.NewSkillsModel(w, h)
	m.contactModel = sections.NewContactModel(w, h)
	m.easterModel = sections.NewEasterModel(w, h)
	m.easterModel.Active = (m.activeTab == TabEasterEgg)
	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.skillsModel.Init(), m.easterModel.Init())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.aboutModel.Width = msg.Width
		m.aboutModel.Height = msg.Height
		m.projectsModel.Width = msg.Width
		m.projectsModel.Height = msg.Height
		m.skillsModel.Width = msg.Width
		m.skillsModel.Height = msg.Height
		m.contactModel.Width = msg.Width
		m.contactModel.Height = msg.Height
		m.easterModel.Width = msg.Width
		m.easterModel.Height = msg.Height

	case sections.SkillsTickMsg:
		var cmd tea.Cmd
		m.skillsModel, cmd = m.skillsModel.Update(msg)
		return m, cmd

	case sections.AnimTickMsg:
		var cmd tea.Cmd
		m.easterModel, cmd = m.easterModel.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "tab", "right", "l":
			m.activeTab = (m.activeTab + 1) % tabCount
			m.easterModel.Active = (m.activeTab == TabEasterEgg)

		case "shift+tab", "left", "h":
			m.activeTab = (m.activeTab - 1 + tabCount) % tabCount
			m.easterModel.Active = (m.activeTab == TabEasterEgg)

		default:
			switch m.activeTab {
			case TabProjects:
				updated, cmd := m.projectsModel.Update(msg)
				m.projectsModel = updated
				cmds = append(cmds, cmd)
			case TabEasterEgg:
				updated, cmd := m.easterModel.Update(msg)
				m.easterModel = updated
				cmds = append(cmds, cmd)
			}
		}

	default:
		if m.activeTab == TabEasterEgg {
			updated, cmd := m.easterModel.Update(msg)
			m.easterModel = updated
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	parts := []string{
		m.renderTopNav(),
		m.renderMainContainer(),
		m.renderHints(),
	}

	if m.activeTab == TabAbout {
		parts = append(parts, m.aboutModel.RenderNameBanner(navbarWidth))
	}

	return lipgloss.NewStyle().
		Width(mainWidth + 2).
		Align(lipgloss.Center).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				parts...,
			),
		)
}

func (m Model) renderTopNav() string {
	labelStyle := lipgloss.NewStyle().
		Foreground(frameBright).
		Bold(true)

	cellW := navbarWidth / len(tabNames)

	cells := make([]string, len(tabNames))
	for i, name := range tabNames {
		cells[i] = lipgloss.NewStyle().
			Width(cellW).
			Align(lipgloss.Center).
			Render(labelStyle.Render(strings.ToUpper(name)))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, cells...)

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(frameBorder).
		Width(navbarWidth).
		Height(navbarHeight).
		Render(row)
}

func (m Model) renderHints() string {
	text := lipgloss.NewStyle().
		Foreground(frameBright).
		Render("[ ← → to switch pages  ·  q to quit ]")

	row := lipgloss.NewStyle().
		Width(hintsWidth).
		Align(lipgloss.Center).
		Render(text)

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(frameBorder).
		Width(hintsWidth).
		Height(hintsHeight).
		Render(row)
}

func (m Model) renderMainContainer() string {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(frameBorder).
		Width(mainWidth).
		Height(mainHeight).
		Render(m.renderSectionContent())
}

func (m Model) renderMenuBox() string {
	return m.renderSectionContent()
}

func (m Model) renderSectionContent() string {
	contentHeight := m.height - 8
	if contentHeight < 5 {
		contentHeight = 5
	}

	switch m.activeTab {
	case TabAbout:
		return m.aboutModel.View(m.aboutModel.FullBio())
	case TabProjects:
		return m.projectsModel.View(contentHeight)
	case TabSkills:
		return m.skillsModel.View(contentHeight)
	case TabContact:
		return m.contactModel.View(contentHeight)
	case TabEasterEgg:
		return m.easterModel.View(contentHeight)
	}
	return ""
}
