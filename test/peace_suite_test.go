package peace_test

import (
	"github.com/harlock/peace"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"path"
	// "os"
	// "path/filepath"
	"testing"
)

var TEST_ROOT = getRootPath()
var TEST_APP_ROOT = path.Join(TEST_ROOT, "fixtures", "app")

func TestPeace(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Peace Suite")
}

func getRootPath() string {
	// p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// return p
	return "/Users/harlock/src/go/src/github.com/harlock/peace/test"
}

func init() {
	peace.AddSource("js", path.Join(TEST_APP_ROOT, "assets", "js"))
	peace.AddVendorSource("js", path.Join(TEST_APP_ROOT, "vendor", "assets", "js"))
}
