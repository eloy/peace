package peace

import (
	"io/ioutil"
	"net/http"
	"log"
	"mime"
	"path"
)


// Returns the content of the given file
func assetContentAndStatusFromFile(path string) ([]byte, bool) {
	// TODO: Cache
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return content, true
}


func Parse(path string) []string {
	collection := newAssetsCollection(path)
	return collection.files
}

func Content(fileName string) ([]byte, bool) {
	asset, found := getAsset(fileName)
	if found {
		return asset.content, true
	}
	return nil, false
}

func WriteResponse(response http.ResponseWriter, fileName string ) {
	asset, found := getAsset(fileName)
	if !found {
		log.Println("Not found", fileName)
		response.WriteHeader(404)
		return
	}

	ext := path.Ext(fileName)
	contentType := mime.TypeByExtension(ext) + "; charset=utf-8"
	response.Header().Set("Content-Type", contentType)
	response.Write(asset.content)
}

const JS = "js"
const CSS = "css"

var sources map[string][]string
var vendorSources map[string][]string

func AddSource(t string, path string) {
	sources[t] = append(sources[t], path)
}

func AddVendorSource(t string, path string) {
	vendorSources[t] = append(vendorSources[t], path)
}



func init() {
	sources = make(map[string][]string)
	vendorSources = make(map[string][]string)

	sources[JS] = make([]string, 0)
	sources[CSS] = make([]string, 0)

	vendorSources[JS] = make([]string, 0)
	vendorSources[CSS] = make([]string, 0)
}
