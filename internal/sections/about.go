package sections

import (
	"fmt"

	"strings"

	"github.com/charmbracelet/lipgloss"
)

var blockPortrait = []string{
	`⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠋⠉⠉⠉⠉⠉⠉⠉⠙⠛⠻⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⣿⣿⣿⣿⣿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⣿⣿⡿⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣼⣿⣶⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⢀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⡇⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⡇⠀⠀⠀⠀⠀⠘⢿⣿⣿⣿⣿⣿⣿⣿⡟⠏⠁⠀⢠⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⠇⠀⠀⠀⠀⠀⠀⠸⠿⠿⡿⡿⢿⣿⣿⣇⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢿⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⢠⣤⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⣿⣀⠀⣀⣀⣀⣤⣶⣿⣿⣿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⢿⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⠅⠀⠀⠀⠀⠠⡀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⡀⠀⠀⠀⠀⠀⠀⠀⠸⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⡆⠀⠀⠀⠀⠈⣽⠀⠀⠀⠀⠀⠀⠀⠀⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⡇⠀⠀⠀⠀⠠⣿⡀⠀⢀⣴⣤⡄⠀⠀⠘⣿⣿⣿⡿⠟⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⢹⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⣇⠀⠀⠀⠀⠀⣿⡇⠀⠸⣿⣿⣿⣧⡀⠀⠛⢿⣿⠓⢾⣿⣿⣿⣿⡿⠿⢿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⣿⠀⠀⠀⢠⣀⢻⣷⣄⠀⠻⣿⣿⣿⡇⠀⠀⠀⠀⠂⠘⠻⠛⠉⠁⣠⣴⣿⣿⣿⣿⠇⠀⠀⠀⠀⠀⠀⠀⠀⢿⢿⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⣿⣆⠀⠀⠙⣿⡟⢿⣿⣦⡄⠀⠉⠋⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣿⣿⣿⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢤⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⣿⣿⣷⡄⠀⠹⣧⡀⠙⢿⣿⣿⣿⣆⡀⠀⠀⠀⠀⠀⠀⠀⠀⠘⢿⣿⣿⣿⣿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⣿⣿⣿`,
	`⣿⣿⣿⣿⣿⠟⠋⠀⠀⠹⣿⣶⣶⣿⣭⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⣴⣾⣿⣿⣿⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣿⣿⣿`,
	`⣿⣿⣿⠏⢀⡀⠀⠀⠀⠀⠈⠙⠻⠿⢿⣿⣿⣿⣿⣶⣄⠀⠀⠀⠀⠛⠿⠿⠏⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⣿`,
	`⣿⣿⠇⢀⣾⠟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⢿⣿⣿⣿⣿⣶⣤⣀⢠⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⠷⠄⠍`,
	`⣿⣿⠀⠘⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀⠀⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠹⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠀⠀`,
	`⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣿⣿⣿⣿⣿⣄⣀⣀⣀⠀⢀⣠⣴⡞⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠀`,
	`⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⡿⣿⣿⣿⡆⢺⣿⡿⣶⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	`⣿⠀⠀⠀⠀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⠇⠀⠛⢛⣿⡇⢸⡁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠁⠀⠀`,
	`⣿⡀⢀⣴⣿⣿⣿⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣿⣿⣿⣿⡿⠀⠀⢰⣿⣿⠤⣼⡇⠀⠀⠀⠀⠀⠀⠀⠀⡀⠀⠀⠀⠀⠀⠀⠀`,
	`⡿⠇⢼⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⢠⡄⠀⠀⠀⢿⣿⣿⣿⣿⡇⠀⠀⣿⣿⠟⠀⣿⡄⠀⠀⠀⠀⠀⠀⢰⣿⡿⠀⠀⠀⠀⠀⠀⠀`,
}

func makeGradient() []lipgloss.Color {
	rows := len(blockPortrait)
	gradient := make([]lipgloss.Color, rows)

	r1, g1, b1 := 10, 220, 174 // #0ADCAE
	r2, g2, b2 := 1, 111, 88   // #016F58

	for i := 0; i < rows; i++ {
		t := float64(i) / float64(rows-1)

		r := int(float64(r1)*(1-t) + float64(r2)*t)
		g := int(float64(g1)*(1-t) + float64(g2)*t)
		b := int(float64(b1)*(1-t) + float64(b2)*t)

		gradient[i] = lipgloss.Color(
			fmt.Sprintf("#%02X%02X%02X", r, g, b),
		)
	}

	return gradient
}

var portraitGradient = makeGradient()

var nameBanner = []string{
	`                                  ____      __              `,
	` ____  ___  ___  ____ ___  ____  / __/_  __/ /______ _____ _`,
	`/_  / / _ \/ _ \/ __  __ \/ __ \/ /_/ / / / //_/ __  / __  /`,
	` / /_/  __/  __/ / / / / / /_/ / __/ /_/ / ,< / /_/ / /_/ / `,
	`/___/\___/\___/_/ /_/ /_/\____/_/  \__,_/_/|_|\__,_/\__,_/  `,
	`                                                            `,
}

var (
	aboutText   = lipgloss.Color("#057059")
	bannerColor = lipgloss.Color("#9FEBDE")
	panelBorder = lipgloss.Color("#045846")
)

type AboutModel struct {
	Width  int
	Height int
}

func NewAboutModel(w, h int) AboutModel {
	return AboutModel{Width: w, Height: h}
}

func (m AboutModel) FullBio() string {
	return `I am a Metallurgical and Materials 
Engineering student at IIT Kharagpur.

I build websites, desktop applications,
3D animations, and interactive experiences.

My interests span software development,
computer graphics, design, and game
development. Rather than focusing on a
single discipline, I enjoy exploring
the spaces where different fields overlap.

Driven by curiosity, I like learning
new tools, solving problems, and turning
ideas into something tangible.`
}

func (m AboutModel) View(bio string) string {
	promptStyle := lipgloss.NewStyle().Foreground(aboutText)
	bioStyle := lipgloss.NewStyle().Foreground(aboutText)

	var portraitBuilder strings.Builder

	for i, line := range blockPortrait {
		color := portraitGradient[i]

		portraitBuilder.WriteString(
			lipgloss.NewStyle().
				Foreground(color).
				Render(line),
		)

		if i != len(blockPortrait)-1 {
			portraitBuilder.WriteByte('\n')
		}
	}

	portrait := portraitBuilder.String()

	bioContent := lipgloss.JoinVertical(lipgloss.Left,
		promptStyle.Render("$ zoya --about"),
		"",
		bioStyle.Render(bio),
	)

	bioPanel := lipgloss.NewStyle().
		Width(46).
		Height(30).
		Padding(0, 1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(panelBorder).
		Render(bioContent)

	portraitPanel := lipgloss.NewStyle().
		Width(50).
		Height(30).
		Border(lipgloss.NormalBorder()).
		BorderForeground(panelBorder).
		Render(portrait)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		bioPanel,
		// "  ",
		portraitPanel,
	)
}

func (m AboutModel) RenderNameBanner(width int) string {
	bannerStyle := lipgloss.NewStyle().
		Foreground(bannerColor).
		Bold(true)

	banner := bannerStyle.Render(strings.Join(nameBanner, "\n"))

	return lipgloss.NewStyle().
		Width(width).
		Align(lipgloss.Right).
		Render(banner)
}
