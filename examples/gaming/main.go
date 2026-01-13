// Example gaming demonstrates game-style UI elements like health, mana, and XP bars.
//
// Run with: go run ./examples/gaming
//
// Controls:
//   - H: Take damage (reduce health)
//   - R: Use healing potion (restore health)
//   - M: Cast spell (use mana)
//   - Space: Gain experience
//   - L: Level up (if XP full)
//   - Tab: Toggle auto-regeneration
//   - q or Ctrl+C: Quit
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ohmymex/tuslide"
)

var (
	// Character info styles
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("236")).
			Padding(0, 1)

	levelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("226"))

	classStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("213"))

	// Bar label styles
	healthLabelStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("196")).
				Width(8)

	manaLabelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("39")).
			Width(8)

	xpLabelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("226")).
			Width(8)

	staminaLabelStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("46")).
				Width(8)

	// Status styles
	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("243"))

	warningStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196"))

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("87"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(1, 2)
)

type tickMsg time.Time

type model struct {
	health     *tuslide.SliderState
	mana       *tuslide.SliderState
	xp         *tuslide.SliderState
	stamina    *tuslide.SliderState
	level      int
	gold       int
	potions    int
	autoRegen  bool
	lastAction string
	quitting   bool
}

func initialModel() model {
	return model{
		health: tuslide.NewState(
			tuslide.WithMin(0),
			tuslide.WithMax(100),
			tuslide.WithValue(100),
			tuslide.WithStep(1),
		),
		mana: tuslide.NewState(
			tuslide.WithMin(0),
			tuslide.WithMax(80),
			tuslide.WithValue(80),
			tuslide.WithStep(1),
		),
		xp: tuslide.NewState(
			tuslide.WithMin(0),
			tuslide.WithMax(100),
			tuslide.WithValue(0),
			tuslide.WithStep(1),
		),
		stamina: tuslide.NewState(
			tuslide.WithMin(0),
			tuslide.WithMax(50),
			tuslide.WithValue(50),
			tuslide.WithStep(1),
		),
		level:      1,
		gold:       100,
		potions:    3,
		lastAction: "Welcome, adventurer!",
	}
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if m.autoRegen {
			// Regenerate health slowly
			if m.health.Value() < m.health.Max() {
				m.health.SetValue(m.health.Value() + 1)
			}
			// Regenerate mana
			if m.mana.Value() < m.mana.Max() {
				m.mana.SetValue(m.mana.Value() + 2)
			}
			// Regenerate stamina
			if m.stamina.Value() < m.stamina.Max() {
				m.stamina.SetValue(m.stamina.Value() + 3)
			}
		}
		return m, tickCmd()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "h":
			// Take damage
			damage := 15.0
			m.health.SetValue(m.health.Value() - damage)
			if m.health.Value() <= 0 {
				m.lastAction = "You died! Game over..."
			} else if m.health.Value() < 30 {
				m.lastAction = fmt.Sprintf("Critical hit! Took %.0f damage. Health low!", damage)
			} else {
				m.lastAction = fmt.Sprintf("Took %.0f damage from enemy attack.", damage)
			}

		case "r":
			// Use healing potion
			if m.potions > 0 {
				if m.health.Value() >= m.health.Max() {
					m.lastAction = "Health already full!"
				} else {
					m.potions--
					healAmount := 40.0
					m.health.SetValue(m.health.Value() + healAmount)
					m.lastAction = fmt.Sprintf("Used potion! Restored %.0f health. %d potions left.", healAmount, m.potions)
				}
			} else {
				m.lastAction = "No potions left!"
			}

		case "m":
			// Cast spell
			manaCost := 20.0
			if m.mana.Value() >= manaCost {
				m.mana.SetValue(m.mana.Value() - manaCost)
				m.lastAction = fmt.Sprintf("Cast Fireball! Used %.0f mana.", manaCost)
			} else {
				m.lastAction = "Not enough mana!"
			}

		case " ":
			// Gain XP
			xpGain := 15.0
			m.xp.SetValue(m.xp.Value() + xpGain)
			m.gold += 10
			if m.xp.Value() >= m.xp.Max() {
				m.lastAction = fmt.Sprintf("Gained %.0f XP! Ready to level up! (Press L)", xpGain)
			} else {
				m.lastAction = fmt.Sprintf("Defeated enemy! Gained %.0f XP and 10 gold.", xpGain)
			}

		case "l":
			// Level up
			if m.xp.Value() >= m.xp.Max() {
				m.level++
				m.xp.SetValue(0)
				// Increase max stats
				m.health.SetMax(m.health.Max() + 20)
				m.health.SetValue(m.health.Max())
				m.mana.SetMax(m.mana.Max() + 10)
				m.mana.SetValue(m.mana.Max())
				m.stamina.SetMax(m.stamina.Max() + 5)
				m.stamina.SetValue(m.stamina.Max())
				m.lastAction = fmt.Sprintf("LEVEL UP! Now level %d! All stats increased!", m.level)
			} else {
				m.lastAction = "Not enough XP to level up."
			}

		case "tab":
			m.autoRegen = !m.autoRegen
			if m.autoRegen {
				m.lastAction = "Auto-regeneration enabled."
			} else {
				m.lastAction = "Auto-regeneration disabled."
			}

		case "p":
			// Buy potion
			if m.gold >= 25 {
				m.gold -= 25
				m.potions++
				m.lastAction = fmt.Sprintf("Bought potion for 25 gold. Now have %d potions.", m.potions)
			} else {
				m.lastAction = "Not enough gold! Potions cost 25 gold."
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Thanks for playing!\n"
	}

	var b strings.Builder

	// Character header
	b.WriteString("\n")
	b.WriteString(titleStyle.Render(" HERO "))
	b.WriteString("  ")
	b.WriteString(levelStyle.Render(fmt.Sprintf("Level %d", m.level)))
	b.WriteString("  ")
	b.WriteString(classStyle.Render("Warrior"))
	b.WriteString("\n\n")

	// Health bar
	healthBar := tuslide.New(m.health,
		tuslide.WithWidth(40),
		tuslide.WithHandle(false),
		tuslide.WithStyle(tuslide.StyleHealth()),
	)
	healthPct := m.health.Percentage() * 100
	healthText := fmt.Sprintf("%.0f/%.0f", m.health.Value(), m.health.Max())
	healthLabel := healthLabelStyle.Render("HP")
	if healthPct < 30 {
		healthLabel = warningStyle.Render("HP")
	}
	b.WriteString(fmt.Sprintf("  %s %s %s\n", healthLabel, healthBar.View(), healthText))

	// Mana bar
	manaBar := tuslide.New(m.mana,
		tuslide.WithWidth(40),
		tuslide.WithHandle(false),
		tuslide.WithStyle(tuslide.StyleMana()),
	)
	manaText := fmt.Sprintf("%.0f/%.0f", m.mana.Value(), m.mana.Max())
	b.WriteString(fmt.Sprintf("  %s %s %s\n", manaLabelStyle.Render("MP"), manaBar.View(), manaText))

	// Stamina bar
	staminaBar := tuslide.New(m.stamina,
		tuslide.WithWidth(40),
		tuslide.WithHandle(false),
		tuslide.WithSymbols(tuslide.Symbols{
			Filled: "▓",
			Empty:  "░",
			Handle: "",
		}),
		tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("46"))),
		tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("236"))),
	)
	staminaText := fmt.Sprintf("%.0f/%.0f", m.stamina.Value(), m.stamina.Max())
	b.WriteString(fmt.Sprintf("  %s %s %s\n", staminaLabelStyle.Render("STA"), staminaBar.View(), staminaText))

	// XP bar
	xpBar := tuslide.New(m.xp,
		tuslide.WithWidth(40),
		tuslide.WithHandle(false),
		tuslide.WithStyle(tuslide.StyleExperience()),
	)
	xpText := fmt.Sprintf("%.0f/%.0f", m.xp.Value(), m.xp.Max())
	xpLabel := xpLabelStyle.Render("XP")
	if m.xp.Value() >= m.xp.Max() {
		xpLabel = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("226")).Blink(true).Render("XP MAX!")
	}
	b.WriteString(fmt.Sprintf("  %s %s %s\n", xpLabel, xpBar.View(), xpText))

	// Inventory
	b.WriteString("\n")
	b.WriteString(statusStyle.Render(fmt.Sprintf("  Gold: %d  |  Potions: %d", m.gold, m.potions)))
	if m.autoRegen {
		b.WriteString(infoStyle.Render("  [Regen ON]"))
	}
	b.WriteString("\n")

	// Last action
	b.WriteString("\n")
	actionStyle := statusStyle
	if strings.Contains(m.lastAction, "LEVEL UP") {
		actionStyle = levelStyle
	} else if strings.Contains(m.lastAction, "died") || strings.Contains(m.lastAction, "Critical") {
		actionStyle = warningStyle
	} else if strings.Contains(m.lastAction, "Not enough") || strings.Contains(m.lastAction, "No potions") {
		actionStyle = warningStyle
	}
	b.WriteString(fmt.Sprintf("  %s\n", actionStyle.Render(m.lastAction)))

	// Help
	b.WriteString(helpStyle.Render("\n  [H] Take damage  [R] Use potion  [M] Cast spell  [Space] Fight"))
	b.WriteString(helpStyle.Render("\n  [L] Level up  [P] Buy potion (25g)  [Tab] Toggle regen  [Q] Quit"))
	b.WriteString("\n")

	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
