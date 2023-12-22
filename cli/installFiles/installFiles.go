package installFiles

import (
	"fmt"
	"os"
  "time"
  "strings"

  fns "templ-gen/fns"

  "github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	files []os.FileInfo
	index    int
	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
	done     bool
}

var (
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle           = lipgloss.NewStyle().Margin(1, 2)
	checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

func newModel() model {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
  
  src := fns.GetSrcPath()
	
  return model{
		files: fns.GetFiles(src),
		spinner:  s,
		progress: p,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(downloadAndInstall(m.files[m.index].Name()), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case installedPkgMsg:
		if m.index >= len(m.files)-1 {
			// Everything's been installed. We're done!
			m.done = true
			return m, tea.Quit
		}

		// Update progress bar
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.files)-1))

		m.index++
		return m, tea.Batch(
			progressCmd,
			tea.Printf("%s %s", checkMark, m.files[m.index].Name()), // print success message above our program
			downloadAndInstall(m.files[m.index].Name()),             // download the next package
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	n := len(m.files)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render(fmt.Sprintf("Done! Installed %d files/directories.\nRun \"air\" to serve.\n", n))
	}

	pkgCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n-1)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+pkgCount))

	pkgName := currentPkgNameStyle.Render(m.files[m.index].Name())
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Installing " + pkgName)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+pkgCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + pkgCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type installedPkgMsg string

func downloadAndInstall(file string) tea.Cmd {
  // Estamate the duration of time to install file based on file size
  src := fns.GetSrcPath()
  // Gives the size in bytes (int64)
  size, _ := fns.GetSize(fmt.Sprintf("%s/files/%s", src, file))

	d := time.Millisecond * time.Duration(size / 10) //nolint:gosec

	return tea.Tick(d, func(t time.Time) tea.Msg {
		fns.InstallFile()

    return installedPkgMsg(file)
	})
}

func InstallFiles() {
  if _, err := tea.NewProgram(newModel()).Run(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }
}
