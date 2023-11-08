package main

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Program         lipgloss.Style
	TetriminoStyles map[byte]lipgloss.Style
}

func DefaultStyles() *Styles {
	return &Styles{
		Program: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0),
		TetriminoStyles: map[byte]lipgloss.Style{
			'I': lipgloss.NewStyle().Foreground(lipgloss.Color("#64C4EB")),
			'O': lipgloss.NewStyle().Foreground(lipgloss.Color("#F1D448")),
			'T': lipgloss.NewStyle().Foreground(lipgloss.Color("#A15398")),
			'S': lipgloss.NewStyle().Foreground(lipgloss.Color("#64B452")),
			'Z': lipgloss.NewStyle().Foreground(lipgloss.Color("#DC3A35")),
			'J': lipgloss.NewStyle().Foreground(lipgloss.Color("#5C65A8")),
			'L': lipgloss.NewStyle().Foreground(lipgloss.Color("#E07F3A")),
		},
	}
}
