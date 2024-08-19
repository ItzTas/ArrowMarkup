package reader

import (
	"strings"

	"github.com/spf13/afero"
)

type File struct {
	Filename string
	Filepath string
	Content  string
}

func normalizePath(path string) string {
	return strings.TrimPrefix(path, "./")
}

type Reader struct {
	Reader afero.Fs
	Files  []File
}

func NewReader() Reader {
	return Reader{
		Reader: afero.NewReadOnlyFs(afero.NewOsFs()),
		Files:  []File{},
	}
}

func (r *Reader) ReadDir(dir string) error {
	infos, err := afero.ReadDir(r.Reader, dir)
	if err != nil {
		return err
	}

	for _, info := range infos {
		path := dir + "/" + info.Name()
		if info.IsDir() {
			err := r.ReadDir(path)
			if err != nil {
				return err
			}
			continue
		}
		if len(info.Name()) < 3 {
			continue
		}
		if info.Name()[len(info.Name())-3:] != ".am" {
			continue
		}
		content, err := afero.ReadFile(r.Reader, path)
		if err != nil {
			return err
		}
		r.Files = append(r.Files, File{
			Filename: info.Name(),
			Filepath: normalizePath(path),
			Content:  string(content),
		})
	}
	return nil
}

func (r *Reader) GetFiles() []File {
	return r.Files
}
