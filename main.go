package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	filePath := os.Args[1]
	dat, err := os.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("xcode-dark"),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert(dat, &buf); err != nil {
		panic(err)
	}

	output := buf.String()
	output = fmt.Sprintf("<div id=\"container\">\n%s\n</div>", output)
	output = fmt.Sprintf("<body>\n%s\n</body>", output)
	output = fmt.Sprintf("<header><style>%s%s</style></header>\n%s", getCSSReset(), getCSSStyles(), output)

	err = os.WriteFile("index.html", []byte(output), 0644)
	if err != nil {
		panic(err.Error())
	}
}

func getCSSReset() string {
	// Define CSS reset styles here
	return `
    /* CSS Reset */
    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }
    `
}

func getCSSStyles() string {
	// Define your CSS styles here
	return `
	/* Import Inconsolata font */
    @import url('https://fonts.googleapis.com/css2?family=Inconsolata:wght@200..900&display=swap');

	:root {
		--text-color: #333; /* Default text color */
		--background-color: #fff; /* Default background color */
	}

	@media (prefers-color-scheme: dark) {
		/* Dark theme */
		:root {
		  --text-color: #ddd; /* Dark text color */
		  --background-color: #222; /* Dark background color */
		}
	}

	/* For screens wider than 1200px (desktops and large screens) */
	@media only screen and (min-width: 1200px) {
		#container {
			max-width: 800px;
			margin: 0 auto; /* Center the div horizontally */
		}
	}
	
	/* For screens wider than 992px but less than or equal to 1200px (large desktops) */
	@media only screen and (max-width: 1200px) {
		#container {
			max-width: 600px;
			margin: 0 auto; /* Center the div horizontally */
		}
	}
	
	/* For screens wider than 576px but less than or equal to 992px (tablets and smaller desktops) */
	@media only screen and (max-width: 992px) {
		#container {
			max-width: calc(100% - 20px);
			margin: 0 auto; /* Center the div horizontally */
		}
	}
	
	/* For screens less than or equal to 576px (mobile devices) */
	@media only screen and (max-width: 576px) {
		#container {
			width: calc(100% - 20px); /* Full width minus 20px margin */
			margin-left: 10px; /* Add left margin to align with the body */
			margin-right: 10px; /* Add right margin to align with the body */
		}
	}

	body {
		color: var(--text-color);
		background-color: var(--background-color);
	}

	#container {
		padding: 10px;
	}

	p {
        margin-bottom: 2rem;
		text-align: justify;
		line-height: 1.5;
    }

	pre {
		display: inline-block;
		white-space: nowrap;
		overflow-x: auto;
		overflow-y: clip;
		width: 100%;
		padding: 1rem;
		border-radius: 0.5rem;
	}

	/* Apply Inconsolata font to code elements within p and pre elements */
	p code,
	pre code {
		font-family: 'Inconsolata', monospace;
	}
	
	p code {
		padding: 0.1rem 0.5rem;
		border-radius: 0.25rem;
		color: var(--code-text-color);
		background-color: var(--code-background-color);
	}
	
	/* Dark and light mode for code elements within p and pre elements */
	@media (prefers-color-scheme: dark) {
		p code {
			--code-text-color: #eee; /* Dark code text color */
			--code-background-color: #444; /* Dark code background color */
		}
	}
	
	@media (prefers-color-scheme: light) {
		p code {
			--code-text-color: #333; /* Light code text color */
			--code-background-color: #f0f0f0; /* Light code background color */
		}
	}

	img {
		width: 100%;
	}
	`
}
