package main

import "time"

type plistBookmarks struct {
	Children []plistChild `plist:"Children"`
}

type plistChild struct {
	Children               []plistChild            `plist:"Children"` // Tree style recursion
	Title                  string                  `plist:"Title"`
	ShouldOmitFromUI       bool                    `plist:"ShouldOmitFromUI"`
	WebBookmarkFileVersion int                     `plist:"WebBookmarkFileVersion"`
	WebBookmarkIdentifier  string                  `plist:"WebBookmarkIdentifier"`
	WebBookmarkType        string                  `plist:"WebBookmarkType"`
	WebBookmarkUUID        string                  `plist:"WebBookmarkUUID"`
	URLString              string                  `plist:"URLString"`
	URIDictionary          plistURIDictionary      `plist:"URIDictionary"`
	Sync                   plistSync               `plist:"Sync"`
	ReadingList            plistReadingList        `plist:"ReadingList"`
	ReadingListNonSync     plistReadingListNonSync `plist:"ReadingListNonSync"`
}

type plistSync struct {
	Key        string `plist:"Key"`
	ServerID   string `plist:"ServerID"`
	ServerData string `plist:"ServerData"`
}

type plistURIDictionary struct {
	Title string `plist:"title"`
}

type plistReadingList struct {
	DateAdded              time.Time `plist:"DateAdded"`
	DateLastViewed         time.Time `plist:"DateLastViewed"`
	DateLastFetched        time.Time `plist:"DateLastFetched"`
	PreviewText            string    `plist:"PreviewText"`
	SourceBundleID         string    `plist:"SourceBundleID"`
	SourceLocalizedAppName string    `plist:"SourceLocalizedAppName"`
}

type plistReadingListNonSync struct {
	DateLastFetched                     time.Time `plist:"DateLastFetched"`
	ArchiveOnDisk                       bool      `plist:"ArchiveOnDisk"`
	FetchResult                         int       `plist:"FetchResult"`
	Title                               string    `plist:"Title"`
	PreviewText                         string    `plist:"PreviewText"`
	URLStringsForAdditionalPages        []string  `plist:"URLStringsForAdditionalPages"`
	NumberOfFailedLoadsWithUnknownError int       `plist:"NumberOfFailedLoadsWithUnknownError"`
}
