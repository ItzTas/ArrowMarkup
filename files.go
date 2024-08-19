package main

import (
	"fmt"
	"strings"

	"github.com/ItzTas/arrowmarkup/internal/reader"
)

func convertAMDirToHTML(dirPath string) ([]HTMLFile, error) {
	reader := reader.NewReader()
	err := reader.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	amParser := NewAmParser()
	htmlfs := []HTMLFile{}
	for _, f := range reader.GetFiles() {
		fmt.Println(f.Content)
		ams, err := amParser.parseAMs(f.Content)
		if err != nil {
			return nil, err
		}
		content := ""
		for _, a := range ams {
			cont, err := a.toHTML()
			if err != nil {
				return nil, err
			}
			content += cont + "\n"
		}
		content = upperTemplate + content + bottomTemplate
		htmlf := HTMLFile{
			filepath: f.Filepath,
			fileName: f.Filename,
			content:  content,
		}
		htmlfs = append(htmlfs, htmlf)
	}
	return htmlfs, nil
}

func purgeEmptyStrFromSlice(slc []string) []string {
	new := []string{}
	for _, s := range slc {
		if strings.Trim(s, " ") == "" {
			continue
		}
		new = append(new, s)
	}
	return new
}
