package prompts

import (
	_ "embed"
	"strings"
)

//go:embed solve-issue.md
var solveIssue string

//go:embed review-pr.md
var reviewPr string

//go:embed respond.md
var respond string

//go:embed release-notes.md
var releaseNotes string

func Default(mode string) string {
	switch mode {
	case "solve-issue":
		return solveIssue
	case "review-pr":
		return reviewPr
	case "respond":
		return respond
	case "release-notes":
		return releaseNotes
	}
	return ""
}

func Compose(mode, promptFileContent string) string {
	base := strings.TrimSpace(Default(mode))
	extra := strings.TrimSpace(promptFileContent)
	switch {
	case base == "":
		return extra
	case extra == "":
		return base
	default:
		return base + "\n\n## Additional instructions from the workflow\n\n" + extra
	}
}
