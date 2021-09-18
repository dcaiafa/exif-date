package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	exif "github.com/dsoprea/go-exif/v3"
)

func parseDateTime(v string) string {
	const layout = "2006:01:02 15:04:05"
	d, err := time.Parse(layout, v)
	if err != nil {
		log.Fatal(err)
	}
	return d.Format(time.RFC3339)
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		log.Fatal("Missing filename")
	}

	filename := flag.Arg(0)

	rawExif, err := exif.SearchFileAndExtractExif(filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	var tags []exif.ExifTag
	tags, _, err = exif.GetFlatExifData(rawExif, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	var dateTimeOriginalTag *exif.ExifTag
	for i := range tags {
		switch tags[i].TagName {
		case "DateTimeOriginal":
			dateTimeOriginalTag = &tags[i]
		default:
		}
	}

	if dateTimeOriginalTag != nil {
		fmt.Println(parseDateTime(dateTimeOriginalTag.Value.(string)))
	}
}
