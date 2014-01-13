package peace

import (
	"regexp"
	"path"
	"io/ioutil"
)

// Regular expressions for match require in include files
//----------------------------------------------------------------------

var requireRegexp = regexp.MustCompile(`//=\s?require (.+)`)
var requireTreeRegexp = regexp.MustCompile(`//=\s?require_tree (.+)`)

type AssetsCollection struct {
	SourceFile string
	base string
	files []string
}

func newAssetsCollection(sourceFileName string) *AssetsCollection {
	collection := new(AssetsCollection)
	collection.files = make([]string, 0)

	collection.SourceFile = sourceFileName
	collection.base = path.Dir(sourceFileName)
	collection.load()
	return collection
}

func (this *AssetsCollection) addFile(file string) {
	this.files = append(this.files, file)
}


func (this *AssetsCollection) loadTree(sourceDir string, sufix string) {
	subdirs := make([]string, 0)

	dir := path.Join(this.base, sourceDir)

	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() == false {
			name := f.Name()
			if ext := path.Ext(name); ext != sufix {
				name = name[:len(name) - len(ext)]
			}
			this.addFile(sourceDir + "/" + name)
		} else {
			subdirs = append(subdirs, path.Join(sourceDir, f.Name()))
		}
	}

	for _, subdir := range subdirs {
		this.loadTree(subdir, sufix)
	}
}

func (this *AssetsCollection) load() {
	sufix := path.Ext(this.SourceFile)

	content, change := assetContentAndStatusFromFile(this.SourceFile)
	if change {
		// Include single files
		found := requireRegexp.FindAllStringSubmatch(string(content), -1)
		for _, match := range found {
			this.addFile(match[1] + sufix)
		}

		// Include Tree
		found = requireTreeRegexp.FindAllStringSubmatch(string(content), -1)
		for _, match := range found {
			this.loadTree(match[1], sufix)
		}
	}
}
