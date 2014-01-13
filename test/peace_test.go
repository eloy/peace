package peace_test

import (
	"github.com/harlock/peace"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"path"
	"io/ioutil"
)

var _ = Describe("Peace", func() {
	var jsIncludeFile = path.Join(TEST_APP_ROOT, "assets", "js", "application.js")

	// Parse
	//----------------------------------------------------------------------

	Describe("Parse(string)", func() {
		It("Should return an AssetCollection object", func() {
			res := peace.Parse(jsIncludeFile)
			expected := []string {"vendor.js", "javascript.js", "coffee.js", "source_tree/test1.js", "source_tree/test2.js", "source_tree/subdir/test3.js",}
			Expect(res).To(Equal(expected))
		})
	})

	// Content(string)
	//----------------------------------------------------------------------

	Describe("Content(string)", func() {
		Context("Javascript File", func() {
			It("Should return the file content if the file is in the vendor dir", func() {
				expected, _ := ioutil.ReadFile(path.Join(TEST_APP_ROOT, "vendor", "assets", "js", "vendor.js"))
				content, found := peace.Content("vendor.js")
				Expect(found).To(BeTrue())
				Expect(content).To(Equal(expected))
			})

			It("Should return the file content if the file is in the sources dir", func() {
				expected, _ := ioutil.ReadFile(path.Join(TEST_APP_ROOT, "assets", "js", "javascript.js"))
				content, found := peace.Content("javascript.js")
				Expect(found).To(BeTrue())
				Expect(content).To(Equal(expected))
			})
		})

		Context("Coffee File", func() {
			It("Should return the compiled js content", func() {
				expected, _ := ioutil.ReadFile(path.Join(TEST_APP_ROOT, "assets", "js", "_coffee.compiled"))
				content, found := peace.Content("coffee.js")
				Expect(found).To(BeTrue())
				Expect(content).To(Equal(expected))
			})
		})
	})


})
