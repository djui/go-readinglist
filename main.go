package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"sort"
	"strings"
	"time"

	plist "github.com/DHowett/go-plist"
)

const defaultBookmarkPath = "~/Library/Safari/Bookmarks.plist"

func main() {
	var inputFlag string
	var jsonFlag bool
	flag.StringVar(&inputFlag, "input", defaultBookmarkPath, "Path to bookmark file")
	flag.BoolVar(&jsonFlag, "json", false, "Output as JSON")
	flag.Parse()

	bookmarksPath, err := expandTilde(inputFlag)
	assertNoError(err)

	f, err := os.Open(bookmarksPath)
	assertNoError(err)

	var bookmarks plistBookmarks
	decoder := plist.NewDecoder(f)
	err = decoder.Decode(&bookmarks)
	assertNoError(err)

	var readinglist []readinglistItem
	for _, child := range bookmarks.Children {
		if child.Title == "com.apple.ReadingList" {
			for _, readinglistChild := range child.Children {
				readinglist = append(readinglist, fromPlist(readinglistChild))
			}
		}
	}

	unreadlist := latestUnread(readinglist)

	if jsonFlag {
		j, err := json.Marshal(unreadlist)
		assertNoError(err)
		fmt.Println(string(j))
	} else {
		for _, item := range unreadlist {
			fmt.Println(item.URL)
		}
	}
}

func latestUnread(readinglist []readinglistItem) []readinglistItem {
	var res []readinglistItem
	for _, item := range readinglist {
		if !item.Read {
			res = append(res, item)
		}
	}
	sort.Sort(byAdded(res))
	return res
}

// fromPlist converts the Plist item to a reading list item. Prefers offline'd
// content field values.
func fromPlist(item plistChild) readinglistItem {
	preview := item.ReadingListNonSync.PreviewText
	if preview == "" {
		preview = item.ReadingList.PreviewText
	}

	title := item.ReadingListNonSync.Title
	if title == "" {
		title = item.URIDictionary.Title
	}

	return readinglistItem{
		Title:   title,
		Preview: preview,
		URL:     item.URLString,
		UUID:    item.WebBookmarkUUID,
		Source:  item.ReadingList.SourceLocalizedAppName,
		Read:    !item.ReadingList.DateLastViewed.IsZero(),
		Added:   item.ReadingList.DateAdded,
		Viewed:  item.ReadingList.DateLastViewed,
		fetched: item.ReadingListNonSync.DateLastFetched,
	}
}

type readinglistItem struct {
	URL     string    `json:"url,omitempty"`
	Title   string    `json:"title,omitempty"`
	Source  string    `json:"source,omitempty"`
	Preview string    `json:"-"`
	Read    bool      `json:"-"`
	UUID    string    `json:"uuid,omitempty"` // ~/Library/Safari/ReadingListArchives
	Added   time.Time `json:"date,omitempty"`
	Viewed  time.Time `json:"-"`
	fetched time.Time // For sorting
}

type byAdded []readinglistItem

func (r byAdded) Len() int {
	return len(r)
}
func (r byAdded) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r byAdded) Less(i, j int) bool {
	if r[i].Added != r[j].Added {
		return r[i].Added.After(r[j].Added)
	}
	return r[i].fetched.After(r[j].fetched)
}

func expandTilde(path string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	if path == "~" || path[:2] == "~/" {
		return strings.Replace(path, "~", usr.HomeDir, 1), nil
	}

	return path, nil
}

func assertNoError(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
