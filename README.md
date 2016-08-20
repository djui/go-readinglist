# go-readinglist

Export Safari's ReadingList as list of URLs or JSON.


## Installation

    $ go get github.com/djui/go-readinglist


## Usage

    $ go-readinglist -h
    Usage of go-readinglist:
      -input string
    Path to bookmark file (default "~/Library/Safari/Bookmarks.plist")
      -json
    Output as JSON


## Dependencies

- [github.com/DHowett/go-plist](https://github.com/DHowett/go-plist)
    $ go get github.com/DHowett/go-plist
