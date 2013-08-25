// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package renderer

import (
	"fmt"
	"github.com/andreaskoch/allmark/repository"
	"io"
	"path/filepath"
)

var (

	// sort the items by date and folder name
	dateAndFolder = func(item1, item2 *repository.Item) bool {

		if item1.MetaData.Date.Equal(item2.MetaData.Date) {
			// ascending by directory name
			return filepath.Base(item1.Directory()) < filepath.Base(item2.Directory())
		}

		// descending by date
		return item1.MetaData.Date.After(item2.MetaData.Date)
	}
)

func (renderer *Renderer) XMLSitemap(w io.Writer, host string) {

	fmt.Fprintln(w, `<?xml version="1.0" encoding="UTF-8"?>`)
	fmt.Fprintln(w, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)

	// get all child items
	items := repository.GetAllChilds(renderer.root)

	// sort the items by date and folder name
	dateAndFolder := func(item1, item2 *repository.Item) bool {

		if item1.MetaData.Date.Equal(item2.MetaData.Date) {
			// ascending by directory name
			return filepath.Base(item1.Directory()) < filepath.Base(item2.Directory())
		}

		// descending by date
		return item1.MetaData.Date.After(item2.MetaData.Date)
	}

	repository.By(dateAndFolder).Sort(items)

	for _, item := range items {
		fmt.Fprintln(w, `<url>`)
		fmt.Fprintln(w, fmt.Sprintf(`<loc>%s</loc>`, getItemLocation(host, item)))
		fmt.Fprintln(w, fmt.Sprintf(`<lastmod>%s</lastmod>`, getItemDate(item)))
		fmt.Fprintln(w, `</url>`)
	}

	fmt.Fprintln(w, `</urlset>`)

}