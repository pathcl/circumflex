package filter

import (
	ansi "clx/utils/strip-ansi"
	"strings"
)

type RuleSet struct {
	skipLineContains []string
	skipLineEquals   []string
	skipParContains  []string
	skipParEquals    []string
	endLineContains  []string
	endLineEquals    []string
}

func (rs *RuleSet) Filter(text string) string {
	paragraphs := strings.Split(text, "\n\n")
	output := ""

	output = filterByParagraph(paragraphs, output, rs)

	lines := strings.Split(output, "\n")
	output = ""

	output = filterByLine(lines, output, rs)

	output = strings.ReplaceAll(output, "\n\n\n\n", "\n\n\n")
	output = strings.ReplaceAll(output, "\n\n\n", "\n\n")
	output = strings.ReplaceAll(output, "\n\n\n", "\n\n")
	output = strings.ReplaceAll(output, "\n\n\n", "\n\n")

	return output
}

func filterByLine(lines []string, output string, rs *RuleSet) string {
	for i, line := range lines {
		isOnFirstOrLastLine := i == 0 || i == len(lines)-1
		lineNoLeadingWhitespace := strings.TrimLeft(line, " ")

		if len(lineNoLeadingWhitespace) == 1 {
			continue
		}

		if isOnFirstOrLastLine {
			output += line + "\n"

			continue
		}

		if equals(rs.skipLineEquals, line) ||
			contains(rs.skipLineContains, line) {
			continue
		}

		if IsOnLineBeforeTargetEquals(rs.endLineEquals, lines, i) ||
			IsOnLineBeforeTargetContains(rs.endLineContains, lines, i) {
			output += "\n"

			break
		}

		output += line + "\n"
	}

	return output
}

func filterByParagraph(paragraphs []string, output string, rs *RuleSet) string {
	for i, paragraph := range paragraphs {
		isOnFirstOrLastParagraph := i == 0 || i == len(paragraphs)-1
		parNoLeadingWhitespace := strings.TrimLeft(paragraph, " ")

		if len(parNoLeadingWhitespace) == 1 {
			continue
		}

		if isOnFirstOrLastParagraph {
			output += paragraph + "\n\n"

			continue
		}

		if equals(rs.skipLineEquals, paragraph) ||
			contains(rs.skipLineContains, paragraph) {
			continue
		}

		output += paragraph + "\n\n"
	}

	return output
}

func (rs *RuleSet) SkipLineContains(text string) {
	rs.skipLineContains = append(rs.skipLineContains, text)
}

func (rs *RuleSet) SkipLineEquals(text string) {
	rs.skipLineEquals = append(rs.skipLineEquals, text)
}

func (rs *RuleSet) SkipParContains(text string) {
	rs.skipParContains = append(rs.skipParContains, text)
}

func (rs *RuleSet) SkipParEquals(text string) {
	rs.skipParEquals = append(rs.skipParEquals, text)
}

func (rs *RuleSet) EndBeforeLineContains(text string) {
	rs.endLineContains = append(rs.endLineContains, text)
}

func (rs *RuleSet) EndBeforeLineEquals(text string) {
	rs.endLineEquals = append(rs.endLineEquals, text)
}

func equals(targets []string, line string) bool {
	for _, target := range targets {
		line = ansi.Strip(line)
		line = strings.TrimLeft(line, " ")

		if line == target {
			return true
		}
	}

	return false
}

func contains(targets []string, line string) bool {
	for _, target := range targets {
		if strings.Contains(line, target) {
			return true
		}
	}

	return false
}

func IsOnLineBeforeTargetEquals(targets []string, lines []string, i int) bool {
	for _, target := range targets {
		nextLine := lines[i+1]
		nextLine = ansi.Strip(nextLine)
		nextLine = strings.TrimLeft(nextLine, " ")

		if nextLine == target {
			return true
		}
	}

	return false
}

func IsOnLineBeforeTargetContains(targets []string, lines []string, i int) bool {
	for _, target := range targets {
		nextLine := lines[i+1]
		nextLine = ansi.Strip(nextLine)
		nextLine = strings.TrimLeft(nextLine, " ")

		if strings.Contains(nextLine, target) {
			return true
		}
	}

	return false
}
