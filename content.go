package peace

import(
	"path"
	"os"
	"io/ioutil"
	"time"
	"os/exec"
)

type Provider interface {
	Ext() string
	Compile(string) []byte
}

var providers map[string][]Provider

type cachedAsset struct {
	content []byte
	updated time.Time
	path string
}

func newCachedAsset(p string) *cachedAsset {
	a := new(cachedAsset)
	a.path = p
	a.updated = time.Now()
	return a
}


func newCachedAssetFromFile(p string) *cachedAsset {
	a := newCachedAsset(p)
	a.content, _ = ioutil.ReadFile(p)
	return a
}

var cache = make(map[string]cachedAsset)


func getAsset(fileName string) (*cachedAsset, bool){
	return readAsset(fileName)
}

func readAsset(fileName string) (*cachedAsset, bool){
	// try vendor first
	ext := path.Ext(fileName)
	if ext == "" {
		return nil, false
	}
	t := ext[1:]
	if t != JS && t != CSS {
		// Only configured extensions
		return nil, false
	}

	if cached, found := readAssetFromVendor(fileName, t); found {
		return cached, true
	}

	for _, base := range sources[t] {
		fullName := path.Join(base, fileName)

		// try with the raw name
		if _, err := os.Stat(fullName); err == nil {
			return newCachedAssetFromFile(fullName), true
		}

		// Try with the providers
		for _, provider := range providers[t] {
			providerName := fullName + provider.Ext()
			if _, err := os.Stat(providerName); err == nil {
				asset := newCachedAsset(providerName)
				asset.content = provider.Compile(providerName)
				return asset, true
			}
		}
	}


	return nil, false
}


func readAssetFromVendor(fileName string, t string) (*cachedAsset, bool) {
	for _, base := range vendorSources[t] {
		fullName := path.Join(base, fileName)
		if _, err := os.Stat(fullName); err == nil {
			return newCachedAssetFromFile(fullName), true
		}
	}
	return nil, false
}




type CoffeeProvider struct {
}

func (this CoffeeProvider) Ext() string {
	return ".coffee"
}

func (this CoffeeProvider) Compile(fileName string) []byte {
	out, err := exec.Command("coffee","-cp", fileName).Output()
	if err != nil {
		panic(err)
	}
	return out
}


func init() {
	providers = make(map[string][]Provider)
	providers[JS] = []Provider{CoffeeProvider{}}
	providers[CSS] = []Provider{}
}
