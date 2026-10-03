package tui

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"

	"github.com/diisn/pidan/internal/provider"
	"github.com/diisn/pidan/internal/selfupdate"
)

// This file builds the startup splash shown at the top of the transcript: a
// minimal text panel with a "pidan" title, a one-line tagline, and the
// session's basic configuration (model, provider, protocol, thinking effort,
// directory) laid out in rows. It is seeded once by withSession so it scrolls
// up as the conversation grows, like a shell's login banner.

// renderBanner builds the startup config panel. Its only I/O is a single cheap
// read of the local update-check cache (no network — CachedLatest); it never
// panics, so it is safe to build eagerly at startup.
func renderBanner(theme Theme, opts Options, cwd string) string {
	title := lipgloss.NewStyle().Foreground(lipgloss.Color(colorAccent)).Bold(true)
	label := lipgloss.NewStyle().Foreground(lipgloss.Color(colorGray))
	value := lipgloss.NewStyle().Foreground(lipgloss.Color(colorUser)).Bold(true)

	rows := [][2]string{
		{"Version", firstNonEmpty(opts.Version, "dev")},
		{"Model", firstNonEmpty(opts.Model, "—")},
		{"Provider", firstNonEmpty(opts.ProviderName, "—")},
		{"Protocol", firstNonEmpty(provider.ProtocolLabel(opts.Protocol), "—")},
		{"Thinking", firstNonEmpty(string(opts.ThinkingLevel), "off")},
		{"Directory", firstNonEmpty(cwd, "—")},
	}

	// When the cached latest-release check says a newer version exists, append a
	// highlighted "→ vX.Y.Z" and an upgrade hint to the Version row. The check is
	// read from the local cache only (no network here); a background refresh keeps
	// it current for the next launch. dev/unparseable versions never trigger this.
	upgradeHint := ""
	if latest, _ := selfupdate.CachedLatest(); latest != "" {
		if avail, comparable := selfupdate.UpdateAvailable(opts.Version, latest); comparable && avail {
			newVer := lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Bold(true).Render(latest)
			rows[0][1] = rows[0][1] + "  →  " + newVer
			upgradeHint = label.Render(strings.Repeat(" ", 11)) +
				lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Render("Run pidan update to upgrade")
		}
	}

	var info strings.Builder
	info.WriteString(title.Render("pidan") + "  " + theme.System.Render("Terminal AI coding assistant") + "\n\n")
	for i, r := range rows {
		if i > 0 {
			info.WriteByte('\n')
		}
		info.WriteString(label.Render(fmt.Sprintf("%-10s ", r[0])) + value.Render(r[1]))
	}
	if upgradeHint != "" {
		info.WriteString("\n" + upgradeHint)
	}

	return info.String()
}

// firstNonEmpty returns s when it is non-empty, otherwise the fallback.
func firstNonEmpty(s, fallback string) string {
	if strings.TrimSpace(s) == "" {
		return fallback
	}
	return s
}
