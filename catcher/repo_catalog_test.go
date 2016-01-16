package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMarkdownHeaderParsing(t *testing.T) {
	mustParse := func(what string, level int, input string, expected string) {
		header, lvl, ok := MdHeaderItem(input)
		expect(t, "Must parse "+what, ok, true)
		expect(t, "Must parse "+what, lvl, level)
		expect(t, "Must parse "+what, header, strings.TrimSpace(expected))
	}

	mustNotParse := func(what string, input string) {
		header, lvl, ok := MdHeaderItem(input)
		expect(t, "Must not parse "+what, ok, false)
		expect(t, "Must not parse "+what, header, "")
		expect(t, "Must not parse "+what, lvl, 0)
	}

	mustNotParse("empty string", "")
	mustNotParse("empty header", "# ")
	mustNotParse("empty header with newline", "# \n")

	testHeaders := []string{"A header", "Nospaces", "spaces    "}
	for _, h := range testHeaders {
		mustNotParse("normal text", h)
		mustParse("normal header", 1, "# "+h, h)
		mustParse("normal header with a lot of spaces", 4, "####    "+h, h)
		mustNotParse("header with no space", "#"+h)
		mustParse("normal 3rd level header", 3, "### "+h, h)
		mustParse("normal 6th level header", 6, "###### "+h, h)
		mustNotParse("3rd level header with no space", "###"+h)
		mustParse("normal header with newline", 1, "# "+h+"\n", h)
		mustParse("normal 3rd level header with newline", 3, "### "+h+"\n", h)
		mustParse("normal 3rd level header with newline and spaces before newline", 3, "### "+h+"  \n", h)
		mustParse("normal 3rd level header with newline and spaces after newline", 3, "### "+h+"\n  ", h)

		mustParse("bold text as 7th level header", 7, "*"+h+"*", h)
		mustNotParse("bold text with no leading asterisk", "*"+h)
		mustParse("astreisk list item as 8th level header", 8, "* "+h+"*", h+"*")
		mustParse("hyphen list item as 8th level header", 8, "- "+h, h)
		mustParse("plus list item as 8th level header", 8, "+ "+h, h)
	}
}

func TestRepoItemParsing(t *testing.T) {
	mustParse := func(what string, input string, username string, reponame string) {
		un, rn, ok := MdRepoItem(input)
		expect(t, "Must parse "+what, ok, true)
		expect(t, "Must parse "+what, un, username)
		expect(t, "Must parse "+what, rn, reponame)
	}

	mustNotParse := func(what string, input string) {
		un, rn, ok := MdRepoItem(input)
		expect(t, "Must not parse "+what, ok, false)
		expect(t, "Must not parse "+what, un, "")
		expect(t, "Must not parse "+what, rn, "")
	}

	mustNotParse("empty string", "")
	mustNotParse("string with no repo link", "   Some string   ")
	mustParse("string with repo link", "(http://github.com/moi/monrepo)", "moi", "monrepo")
	mustNotParse("string with wrong link", "(http://github.org/moi/monrepo)")
	mustParse("string with repo link and label", "[Some label](http://github.com/moi/monrepo)", "moi", "monrepo")
	mustParse("string with https repo link", "(https://github.com/moi/monrepo)", "moi", "monrepo")
	mustParse("string with repo link and text", "Text before (https://github.com/moi/monrepo) text after", "moi", "monrepo")
	mustParse("string with repo link, label and text", "* [a label](https://github.com/moi/monrepo) text after", "moi", "monrepo")
	mustParse("string with several links", "* [wrong link](https://example.com/path/) text after [a label](https://github.com/moi/monrepo)", "moi", "monrepo")
	mustParse("first repo if several", "* [wrong link](https://github.com/moi/monrepo) text after [a label](https://github.com/toi/tonrepo)", "moi", "monrepo")
	mustNotParse("no repo github link (less path segments)", "(http://github.com/moi)")
	mustNotParse("no repo github link (more path segments)", "(http://github.com/moi/monrepo/extra)")
}

func TestRepoListParsing(t *testing.T) {
	file, err := os.Open("repo_list_test_sample.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	list, err := ReadRepoCatalog(file)

	assert(t,
		"Error must be nil",
		err == nil,
	)
	expectFatal(t,
		"Block list must have correct length",
		len(list),
		11,
	)

	assertBlock := func(what string, b *RepoEntryBlock, entriesCount int, titles ...string) {
		expectFatal(t,
			fmt.Sprintf("%v must have correct number of entries", what),
			len(b.Entries), entriesCount,
		)
		for i, _ := range b.Entries {
			expect(t,
				fmt.Sprintf("%v's %dth entry must be correct", what, i),
				*b.Entries[i], RepoEntry{"user", fmt.Sprintf("package%v", i+1)},
			)
		}
		expectFatal(t,
			fmt.Sprintf("%v must have correct number of titles", what),
			len(b.Titles), len(titles),
		)
		for i, _ := range titles {
			assert(t,
				fmt.Sprintf("%v's %dth title must be %v", what, i, titles[i]),
				b.Titles[i] == titles[i],
			)
		}
	}

	assertBlock(
		"first block", list[0], 3,
		"Site Header",
		"Simple category",
		"", "", "", "",
		"Sample description",
		"",
	)

	assertBlock(
		"block under subheader", list[1], 3,
		"Site Header",
		"Category with subheaders",
		"Subheader 1",
		"", "", "",
		"Description",
		"",
	)

	assertBlock(
		"block under third category", list[3], 3,
		"Site Header",
		"Simple category after complex one",
		"", "", "", "",
		"Description",
		"",
	)

	assertBlock(
		"block under section", list[5], 3,
		"Site Header",
		"Category with two sections",
		"", "", "", "",
		"Section 2 as description",
		"",
	)

	assertBlock(
		"block under list section", list[6], 3,
		"Site Header",
		"Category with list sections",
		"", "", "", "",
		"Common description",
		"Section 1 as list item",
	)

	assertBlock(
		"block under second simple cat", list[8], 3,
		"Site Header",
		"Simple category after complex one",
		"", "", "", "",
		"Description",
		"",
	)

	assertBlock(
		"block with broken indentation", list[9], 4,
		"Site Header",
		"Category with broken indentation",
		"", "", "", "",
		"Description",
		"",
	)

	assertBlock(
		"block after unparsed entry", list[10], 3,
		"Another FIRST LEVEL header",
		"Category",
		"Subcategory",
		"", "", "", "", "",
	)

}
