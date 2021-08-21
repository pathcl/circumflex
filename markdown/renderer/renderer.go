package renderer

import (
	"clx/indent"
	"clx/markdown"
	"clx/syntax"
	"regexp"
	"strings"

	termtext "github.com/MichaelMure/go-term-text"
	. "github.com/logrusorgru/aurora/v3"
)

const (
	indentLevel1 = "  "
	indentLevel2 = "    "
)

func ToString(blocks []*markdown.Block, lineWidth int, altIndentBlock bool) string {
	output := ""

	for _, block := range blocks {
		switch block.Kind {
		case markdown.Text:
			output += renderText(block.Text, lineWidth, indentLevel1) + "\n\n"

		case markdown.Image:
			output += renderImage(block.Text, lineWidth) + "\n\n"

		case markdown.Code:
			output += renderCode(block.Text) + "\n\n"

		case markdown.Quote:
			output += renderQuote(block.Text, lineWidth, altIndentBlock) + "\n\n"

		case markdown.H1:
			output += h1(block.Text) + "\n\n"

		case markdown.H2:
			output += h2(block.Text) + "\n\n"

		case markdown.H3:
			output += h3(block.Text) + "\n\n"

		case markdown.H4:
			output += h4(block.Text) + "\n\n"

		case markdown.H5:
			output += h5(block.Text) + "\n\n"

		case markdown.H6:
			output += h6(block.Text) + "\n\n"

		case markdown.List:
			output += renderText(block.Text, lineWidth, indentLevel2) + "\n\n"

		default:
			output += renderText(block.Text, lineWidth, indentLevel1) + "\n\n"
		}
	}

	return output
}

func renderText(text string, lineWidth int, indentLevel string) string {
	text = it(text)
	text = bld(text)

	text = syntax.RemoveUnwantedNewLines(text)
	text = syntax.HighlightBackticks(text)

	padding := termtext.WrapPad(indentLevel)
	text, _ = termtext.Wrap(text, lineWidth, padding)

	return text
}

func renderImage(text string, lineWidth int) string {
	magenta := "\u001B[35m"
	italic := "\u001B[3m"
	faint := "\u001B[2m"
	normal := "\u001B[0m"

	exp := regexp.MustCompile(`!\[(.*?)\]\(.*?\)`)
	image := magenta + faint + "Image: " + normal + faint + italic

	// imageLabel := image+italic+faint+`$1.`+"### "

	text = exp.ReplaceAllString(text, image+`$1.`)

	lines := strings.Split(text, image)
	output := ""

	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}

		output += image + line + "\n\n"
	}

	// Remove 'Image: .' for images without captions
	output = strings.ReplaceAll(output, image+".", magenta+faint+"Image ")

	// output = strings.ReplaceAll(output, "###", "")
	// output = strings.ReplaceAll(output, "%%%", "")
	output = strings.TrimSuffix(output, "\n\n")
	output += normal

	output = it(output)
	output = bld(output)

	padding := termtext.WrapPad(indentLevel2)
	output, _ = termtext.Wrap(output, lineWidth, padding)

	return output
}

func renderCode(text string) string {
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimPrefix(text, "\n")

	return text
}

func renderQuote(text string, lineWidth int, altIndentBlock bool) string {
	text = Italic(text).Faint().String()

	text = removeUnwantedNewLines(text)

	indentSymbol := " " + indent.GetIndentSymbol(false, altIndentBlock)
	text = itReversed(text)
	text = bldInQuote(text)

	padding := termtext.WrapPad(indentLevel2 + Faint(indentSymbol).String())
	text, _ = termtext.Wrap(text, 70, padding)

	// text = strings.TrimSuffix(text, "\n")
	// text = strings.TrimPrefix(text, "\n")

	return text
}

func removeUnwantedNewLines(text string) string {
	paragraphSeparator := "\n\n"
	paragraphs := strings.Split(text, paragraphSeparator)
	output := ""

	for _, paragraph := range paragraphs {
		paragraph = syntax.RemoveUnwantedNewLines(paragraph)

		output += paragraph + paragraphSeparator

	}

	output = strings.TrimSuffix(output, paragraphSeparator)

	return output
}

func it(text string) string {
	italic := "\u001B[3m"
	noItalic := "\u001B[23m"

	text = strings.ReplaceAll(text, markdown.ItalicStart, italic)
	text = strings.ReplaceAll(text, markdown.ItalicStop, noItalic)

	return text
}

func itReversed(text string) string {
	italic := "\u001B[3m"
	noItalic := "\u001B[23m"

	text = strings.ReplaceAll(text, markdown.ItalicStart, noItalic)
	text = strings.ReplaceAll(text, markdown.ItalicStop, italic)

	return text
}

func bld(text string) string {
	bold := "\033[31m"
	noBold := "\033[0m"

	text = strings.ReplaceAll(text, markdown.BoldStart, bold)
	text = strings.ReplaceAll(text, markdown.BoldStop, noBold)

	return text
}

func bldInQuote(text string) string {
	// bold := "\033[31m"
	// noBold := "\033[0m"

	text = strings.ReplaceAll(text, markdown.BoldStart, "")
	text = strings.ReplaceAll(text, markdown.BoldStop, "")

	return text
}

func h1(text string) string {
	text = strings.TrimPrefix(text, "# ")

	return Bold(text).String()
}

func h2(text string) string {
	text = strings.TrimPrefix(text, "## ")

	return Bold(text).Blue().String()
}

func h3(text string) string {
	text = strings.TrimPrefix(text, "### ")

	return Bold(text).Yellow().String()
}

func h4(text string) string {
	text = strings.TrimPrefix(text, "#### ")

	return Bold(text).Green().String()
}

func h5(text string) string {
	text = strings.TrimPrefix(text, "#### ")

	return Bold(text).Cyan().String()
}

func h6(text string) string {
	text = strings.TrimPrefix(text, "#### ")

	return Bold(text).Blue().String()
}
