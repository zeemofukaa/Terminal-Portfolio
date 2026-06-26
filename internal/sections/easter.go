package sections

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ── axolotl braille art ──────────────────────────────────────────────────────

var idleArt = []string{
	`⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⣀⣀⡤⠤⠤⠤⣄⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⢴⠋⠙⠳⣤⡀⣠⠖⠋⠁⠀⠀⠀⠀⠀⠀⠀⠉⠓⠤⡀⣠⡴⠟⠛⣷⠀⠀`,
	`⠀⠀⠈⠳⢤⣀⢈⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⢏⣀⣠⡴⠋⠀⠀`,
	`⢀⣠⣤⣄⣄⣉⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣇⣡⣤⡴⣦⣀`,
	`⢹⣅⡀⠀⠀⠈⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡏⠁⠀⢀⣠⠏`,
	`⠀⠈⠙⠉⢋⣉⣇⠀⠀⣾⣷⠄⠀⠀⠀⠀⠀⠀⠀⢴⣿⡆⠀⢀⣿⡙⠋⠋⠀⠀`,
	`⠀⠀⢤⠶⠋⠉⢈⣦⡀⠈⠉⠀⠀⠀⠉⠉⠉⠀⠀⠀⠉⠀⣠⣎⠈⠉⠛⢷⡀⠀`,
	`⠀⠀⠻⣤⣤⠶⠋⠀⠈⠑⠠⠄⣀⣀⣀⣀⣀⣀⡀⠤⠐⠉⠀⠈⠻⠶⠖⠶⠃⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠃⠤⡀⠀⡠⠤⢱⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡎⠀⠉⠁⠀⠉⠉⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⢀⠤⡀⠀⡠⢄⢀⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠢⣀⣀⣀⣈⡠⠊⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡌⡇⠀⣎⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢡⡇⣸⠜⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠓⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
}
var eatArt = []string{
	`⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⣀⣀⡤⠤⠤⠤⣄⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⢴⠋⠙⠳⣤⡀⣠⠖⠋⠁⠀⠀⠀⠀⠀⠀⠀⠉⠓⠤⡀⣠⡴⠟⠛⣷⠀⠀`,
	`⠀⠀⠈⠳⢤⣀⢈⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⢏⣀⣠⡴⠋⠀⠀`,
	`⢀⣠⣤⣄⣄⣉⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣇⣡⣤⡴⣦⣀`,
	`⢹⣅⡀⠀⠀⠈⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡏⠁⠀⢀⣠⠏`,
	`⠀⠈⠙⠉⢋⣉⣇⠀⠀⣾⣷⠄⠀⠀⠀⠀⠀⠀⠀⢴⣿⡆⠀⢀⣿⡙⠋⠋⠀⠀`,
	`⠀⠀⢤⠶⠋⠉⢈⣦⡀⠈⠉⠀⠀⠀⡔⠒⡄⠀⠀⠀⠉⠀⣠⣎⠈⠉⠛⢷⡀⠀`,
	`⠀⠀⠻⣤⣤⠶⠋⠀⠈⠑⠠⠄⣀⣀⣑⣊⣁⣀⡀⠤⠐⠉⠀⠈⠻⠶⠖⠶⠃⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠃⠤⡀⠀⡠⠤⢱⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡎⠀⠉⠁⠀⠉⠉⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⢀⠤⡀⠀⡠⢄⢀⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠢⣀⣀⣀⣈⡠⠊⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡌⡇⠀⣎⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢡⡇⣸⠜⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠓⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
}
var playArt = []string{
	`⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⣀⣀⡤⠤⠤⠤⣄⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⢴⠋⠙⠳⣤⡀⣠⠖⠋⠁⠀⠀⠀⠀⠀⠀⠀⠉⠓⠤⡀⣠⡴⠟⠛⣷⠀⠀`,
	`⠀⠀⠈⠳⢤⣀⢈⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⢏⣀⣠⡴⠋⠀⠀`,
	`⢀⣠⣤⣄⣄⣉⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣇⣡⣤⡴⣦⣀`,
	`⢹⣅⡀⠀⠀⠈⡇⠀⠀⢀⣀⡀⠀⠀⠀⠀⠀⠀⠀⣀⣀⠀⠀⠀⡏⠁⠀⢀⣠⠏`,
	`⠀⠈⠙⠉⢋⣉⣇⠀⠀⠁⠀⠈⠀⠀⠀⠀⠀⠀⠈⠀⠀⠁⠀⢀⣿⡙⠋⠋⠀⠀`,
	`⠀⠀⢤⠶⠋⠉⢈⣦⡀⠀⠀⠀⠀⠀⠑⠒⠊⠀⠀⠀⠀⠀⣠⣎⠈⠉⠛⢷⡀⠀`,
	`⠀⠀⠻⣤⣤⠶⠋⠀⠈⠑⠠⠄⣀⣀⣀⣀⣀⣀⡀⠤⠐⠉⠀⠈⠻⠶⠖⠶⠃⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠃⠤⡀⠀⡠⠤⢱⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡎⠀⠉⠁⠀⠉⠉⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⢀⠤⡀⠀⡠⢄⢀⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠢⣀⣀⣀⣈⡠⠊⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡌⡇⠀⣎⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢡⡇⣸⠜⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠓⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
}
var sleepArt = []string{
	`⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⣀⣀⡤⠤⠤⠤⣄⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⢴⠋⠙⠳⣤⡀⣠⠖⠋⠁⠀⠀⠀⠀⠀⠀⠀⠉⠓⠤⡀⣠⡴⠟⠛⣷⠀⠀`,
	`⠀⠀⠈⠳⢤⣀⢈⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⢏⣀⣠⡴⠋⠀⠀`,
	`⢀⣠⣤⣄⣄⣉⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣇⣡⣤⡴⣦⣀`,
	`⢹⣅⡀⠀⠀⠈⡇⠀⢀⠀⠀⡀⠀⠀⠀⠀⠀⠀⢀⠀⠀⡀⠀⠀⡏⠁⠀⢀⣠⠏`,
	`⠀⠈⠙⠉⢋⣉⣇⠀⠈⠉⠁⠀⠀⠀⠀⠀⠀⠀⠈⠉⠁⠀⠀⢀⣿⡙⠋⠋⠀⠀`,
	`⠀⠀⢤⠶⠋⠉⢈⣦⡀⠀⠀⠀⠀⠀⠉⠉⠉⠀⠀⠀⠀⠀⣠⣎⠈⠉⠛⢷⡀⠀`,
	`⠀⠀⠻⣤⣤⠶⠋⠀⠈⠑⠠⠄⣀⣀⣀⣀⣀⣀⡀⠤⠐⠉⠀⠈⠻⠶⠖⠶⠃⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠃⠤⡀⠀⡠⠤⢱⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡎⠀⠉⠁⠀⠉⠉⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⢀⠤⡀⠀⡠⢄⢀⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠢⣀⣀⣀⣈⡠⠊⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡌⡇⠀⣎⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢡⡇⣸⠜⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠓⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
}
var dedArt = []string{
	`⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⣀⣀⡤⠤⠤⠤⣄⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⢴⠋⠙⠳⣤⡀⣠⠖⠋⠁⠀⠀⠀⠀⠀⠀⠀⠉⠓⠤⡀⣠⡴⠟⠛⣷⠀⠀`,
	`⠀⠀⠈⠳⢤⣀⢈⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⢏⣀⣠⡴⠋⠀⠀`,
	`⢀⣠⣤⣄⣄⣉⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣇⣡⣤⡴⣦⣀`,
	`⢹⣅⡀⠀⠀⠈⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡏⠁⠀⢀⣠⠏`,
	`⠀⠈⠙⠉⢋⣉⣇⠀⠀⠀⠈⡢⡊⠀⠀⠀⠀⠀⡱⢎⠀⠀⠀⢀⣿⡙⠋⠋⠀⠀`,
	`⠀⠀⢤⠶⠋⠉⢈⣦⡀⠀⠀⠀⠀⠀⠈⠉⠁⠀⠀⠀⠀⠀⣠⣎⠈⠉⠛⢷⡀⠀`,
	`⠀⠀⠻⣤⣤⠶⠋⠀⠈⠑⠠⠄⣀⣀⣀⣀⣀⣀⡀⠤⠐⠉⠀⠈⠻⠶⠖⠶⠃⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠃⠤⡀⠀⡠⠤⢱⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡎⠀⠉⠁⠀⠉⠉⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⢀⠤⡀⠀⡠⢄⢀⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠢⣀⣀⣀⣈⡠⠊⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡌⡇⠀⣎⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢡⡇⣸⠜⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠓⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
}

// ── model ────────────────────────────────────────────────────────────────────

type EasterModel struct {
	Width  int
	Height int

	hunger    int
	happiness int
	energy    int

	dead bool

	lastAction string
	status     string

	flashTimer int
	decayTimer int

	state PetState

	Active bool
}

type PetState int

const (
	Idle PetState = iota
	Eating
	Playing
	Sleeping
	Dead
)

type AnimTickMsg struct{}

func animTickCmd() tea.Cmd {
	return tea.Tick(150*time.Millisecond, func(t time.Time) tea.Msg {
		return AnimTickMsg{}
	})
}

func clampStat(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func NewEasterModel(w, h int) EasterModel {
	return EasterModel{
		Width:     w,
		Height:    h,
		hunger:    65,
		happiness: 75,
		energy:    60,
		status:    "Waiting...",
		state:     Idle,
	}
}

func (m EasterModel) Init() tea.Cmd {
	return animTickCmd()
}

func (m EasterModel) Update(msg tea.Msg) (EasterModel, tea.Cmd) {
	switch msg := msg.(type) {

	case AnimTickMsg:
		if !m.Active {
			return m, animTickCmd()
		}
		if m.flashTimer > 0 {
			m.flashTimer--
			if m.flashTimer == 0 {
				m.lastAction = ""
				if m.state != Sleeping {
					m.state = Idle
				}
			}
		}

		m.decayTimer++

		if m.decayTimer >= 10 {
			m.decayTimer = 0

			m.hunger--
			m.happiness--
			m.energy--

			m.hunger = clampStat(m.hunger, 0, 100)
			m.happiness = clampStat(m.happiness, 0, 100)
			m.energy = clampStat(m.energy, 0, 100)

			if m.hunger == 0 && m.happiness == 0 && m.energy == 0 {
				m.dead = true
				m.status = "your pet died :("
				m.state = Dead
			}
		}

		return m, animTickCmd()

	case tea.KeyMsg:

		switch msg.String() {

		case "r":
			m = NewEasterModel(m.Width, m.Height)
			return m, animTickCmd()
		}

		if m.dead {
			return m, nil
		}

		switch msg.String() {

		case "f":
			m.hunger = clampStat(m.hunger+25, 0, 100)
			m.energy = clampStat(m.energy-5, 0, 100)
			m.lastAction = "feed"
			m.flashTimer = 10
			m.status = "fed axolotl"
			m.state = Eating

		case "p":
			m.happiness = clampStat(m.happiness+25, 0, 100)
			m.hunger = clampStat(m.hunger-15, 0, 100)
			m.energy = clampStat(m.energy-15, 0, 100)
			m.lastAction = "play"
			m.flashTimer = 10
			m.status = "played with axolotl"
			m.state = Playing

		case "s":
			m.energy = clampStat(m.energy+35, 0, 100)
			m.hunger = clampStat(m.hunger-5, 0, 100)
			m.lastAction = "sleep"
			m.flashTimer = 10
			m.status = "axolotl is taking a nap..."
			m.state = Sleeping
		}
	}
	return m, nil
}

// ── view ─────────────────────────────────────────────────────────────────────

const easterContentW = 96

func (m EasterModel) View(height int) string {
	teal := lipgloss.Color("#21D4B3")
	tealMid := lipgloss.Color("#057059")
	tealDim := lipgloss.Color("#045846")
	black := lipgloss.Color("#000000")

	labelStyle := lipgloss.NewStyle().Foreground(tealMid)
	barFillStyle := lipgloss.NewStyle().Foreground(teal)
	barEmptyStyle := lipgloss.NewStyle().Foreground(tealDim)
	artStyle := lipgloss.NewStyle().Foreground(teal)
	hintStyle := lipgloss.NewStyle().Foreground(teal)

	pageStyle := lipgloss.NewStyle().
		MarginLeft(2).
		MarginRight(2)

	btnNormal := lipgloss.NewStyle().
		Background(teal).
		Foreground(black).
		Padding(1, 5).
		Width(15).
		Align(lipgloss.Center)

	btnActive := lipgloss.NewStyle().
		Background(teal).
		Foreground(black).
		Padding(1, 5).
		Width(15).
		Align(lipgloss.Center).
		Bold(true)

	btnDisabled := lipgloss.NewStyle().
		Background(tealDim).
		Foreground(black).
		Padding(1, 5).
		Width(15).
		Align(lipgloss.Center)

	// ── stat bars ──
	renderBar := func(val int) string {
		const bw = 20
		filled := int(float64(val) / 100.0 * bw)
		empty := bw - filled
		return "[" +
			barFillStyle.Render(strings.Repeat("█", filled)) +
			barEmptyStyle.Render(strings.Repeat("░", empty)) +
			"]"
	}

	center := func(s string) string {
		return lipgloss.NewStyle().Width(easterContentW).Align(lipgloss.Center).Render(s)
	}

	var sb strings.Builder

	// blank row
	sb.WriteString("\n")

	descStyle := lipgloss.NewStyle().
		Foreground(tealMid)

	cmdStyle := lipgloss.NewStyle().
		Foreground(teal)
		// Bold(true)

	leftCol := lipgloss.NewStyle().Width(62)
	rightCol := lipgloss.NewStyle().Width(34).
		Align(lipgloss.Right)

	renderStat := func(label string, val int) string {
		line := labelStyle.Render(label) + "  " + renderBar(val)
		return rightCol.Render(line)
	}

	renderRow := func(leftText string, leftStyle lipgloss.Style, label string, val int) string {
		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			leftCol.Render(leftStyle.Render(leftText)),
			renderStat(label, val),
		)
	}

	sb.WriteString(renderRow("$ ./axotchi --demo", descStyle, "hunger", m.hunger))
	sb.WriteString("\n")
	sb.WriteString(renderRow("A miniature version of Axotchi.", descStyle, "happy ", m.happiness))
	sb.WriteString("\n")
	sb.WriteString(renderRow("Take care of your pet!", descStyle, "energy", m.energy))
	sb.WriteString("\n\n")
	sb.WriteString(cmdStyle.Render("> " + m.status))
	sb.WriteString("\n\n")

	// Choose art based on current state
	var art []string

	switch m.state {
	case Idle:
		art = idleArt
	case Eating:
		art = eatArt
	case Playing:
		art = playArt
	case Sleeping:
		art = sleepArt
	case Dead:
		art = dedArt
	}

	// Draw axolotl
	for _, line := range art {
		sb.WriteString(center(artStyle.Render(line)))
		sb.WriteString("\n")
	}

	sb.WriteString("\n")

	// buttons
	feedBtn := btnNormal.Render("feed")
	playBtn := btnNormal.Render("play")
	sleepBtn := btnNormal.Render("sleep")

	if m.dead {
		feedBtn = btnDisabled.Render("feed")
		playBtn = btnDisabled.Render("play")
		sleepBtn = btnDisabled.Render("sleep")
	} else {
		// highlight last action only when alive
		switch m.lastAction {
		case "feed":
			feedBtn = btnActive.Render("feed")
		case "play":
			playBtn = btnActive.Render("play")
		case "sleep":
			sleepBtn = btnActive.Render("sleep")
		}
	}

	buttons := lipgloss.JoinHorizontal(lipgloss.Top,
		feedBtn,
		"    ",
		playBtn,
		"    ",
		sleepBtn,
	)
	sb.WriteString(center(buttons))

	sb.WriteString("\n\n\n\n")

	// internal hint
	hint := hintStyle.Render("[ f to feed  ·  p to play  ·  s to sleep  ·  r to restart ]")
	sb.WriteString(center(hint))

	return pageStyle.Render(sb.String())
}
