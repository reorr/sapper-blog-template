package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	begin := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <source> [destination]")
		return
	}

	source := os.Args[1]
	destination := "out"
	if len(os.Args) > 2 {
		destination = os.Args[2]
	}

	err := processFiles(source, destination)
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(begin).Microseconds())
}

func processFiles(source, destination string) error {
	err := filepath.Walk(source, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path == source {
			return nil // Skip non-directories and the source directory itself
		}

		if info.IsDir() {
			relDirPath, err := filepath.Rel(source, path)
			if err != nil {
				return err
			}

			err = createIndexHTML(path, destination, relDirPath)
			if err != nil {
				return err
			}
		}

		if filepath.Ext(path) == ".md" {
			err := convertMarkdownToHTML(path, source, destination)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func createIndexHTML(directory, destination, relDirPath string) error {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	var markdownFiles []fileInfo

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".md" {
			markdownFiles = append(markdownFiles, fileInfo{
				Name: file.Name(),
				Time: file.ModTime(),
			})
		}
	}

	sort.Slice(markdownFiles, func(i, j int) bool {
		return markdownFiles[i].Time.After(markdownFiles[j].Time)
	})

	var links []string
	for _, file := range markdownFiles {
		linkName, err := getHeading1Text(filepath.Join(directory, file.Name))
		if err != nil {
			continue
		}
		link := fmt.Sprintf("<li><a href=\"%s\">%s</a></li>", file.htmlFileName(), linkName)
		links = append(links, link)
	}

	indexContent := fmt.Sprintf("<ul>%s</ul>", strings.Join(links, ""))

	indexFilePath := filepath.Join(destination, relDirPath, "index.html")
	return os.WriteFile(indexFilePath, []byte(indexContent), 0644)
}

func getHeading1Text(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# "), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("heading 1 not found")
}

func convertMarkdownToHTML(source, root, destination string) error {
	dat, err := os.ReadFile(source)
	if err != nil {
		return err
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
		return err
	}

	output := buf.String()
	relPath, err := filepath.Rel(root, source)
	if err != nil {
		return err
	}
	outputPath := filepath.Join(destination, relPath[:len(relPath)-len(filepath.Ext(relPath))]+".html")
	outputDir := filepath.Dir(outputPath)

	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		return err
	}

	output = fmt.Sprintf("<div id=\"container\">\n%s\n</div>", output)
	output = fmt.Sprintf("<body>\n%s\n</body>", output)
	output = fmt.Sprintf("<header><style>%s%s</style></header>\n%s", getCSSReset(), getCSSStyles(), output)

	return os.WriteFile(outputPath, []byte(output), 0644)
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

type fileInfo struct {
	Name string
	Time time.Time
}

func (f fileInfo) htmlFileName() string {
	return fmt.Sprintf("%s.html", f.Name[:len(f.Name)-len(filepath.Ext(f.Name))])
}
