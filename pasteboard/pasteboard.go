package pasteboard

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/progrium/darwinkit/macos/appkit"
)

const AppName = "Rich Paste"

func CLI(args []string) error {
	pboard := appkit.Pasteboard_GeneralPasteboard()
	html := pboard.StringForType(appkit.PasteboardTypeHTML)
	p := bluemonday.StrictPolicy()

	p.AllowElements("br", "ol", "ul", "li", "a", "b", "i", "strong",
		"em", "u", "s", "sub", "sup", "code", "pre", "blockquote",
		"h1", "h2", "h3", "h4", "h5", "h6")
	p.AllowAttrs("href").OnElements("a")
	p.AllowAttrs("style").OnElements("span")
	p.AllowStyles("font-weight", "font-style", "text-decoration").Globally()

	html = p.Sanitize(html)

	pboard.ClearContents()
	pboard.SetStringForType(html, appkit.PasteboardTypeHTML)
	return nil
}
