// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package templates

import (
	"fmt"
)

var masterTemplate = fmt.Sprintf(`<!DOCTYPE HTML>
<html lang="{{.LanguageTag}}">
<head>
	<meta charset="utf-8">
	<title>{{.Title}}</title>

	<link rel="schema.DC" href="http://purl.org/dc/terms/">
	<meta name="DC.date" content="{{.Date}}">

	<link rel="alternate" type="application/rss+xml" title="RSS" href="/rss.xml">

	<link rel="shortcut icon" href="/theme/favicon.ico" />

	<link rel="stylesheet" href="/theme/deck.css" media="screen">
	<link rel="stylesheet" href="/theme/screen.css" media="screen">
	<link rel="stylesheet" href="/theme/print.css" media="print">
	<link rel="stylesheet" href="/theme/codehighlighting/highlight.css" media="screen, print">

	<script src="/theme/modernizr.js"></script>
</head>
<body>

{{ if .ToplevelNavigation}}
<nav class="toplevel">
	<ul>
	{{range .ToplevelNavigation.Entries}}
	<li>
		<a href="{{.Path}}">{{.Title}}</a>
	</li>
	{{end}}
	</ul>
</nav>
{{end}}

{{ if .BreadcrumbNavigation}}
<nav class="breadcrumb">
	<ul>
	{{range .BreadcrumbNavigation.Entries}}
	<li>
		<a href="{{.Path}}">{{.Title}}</a>
	</li>
	{{end}}
	</ul>
</nav>
{{end}}

<article class="{{.Type}} level-{{.Level}}">
%s
</article>

<footer>
	<nav>
		<ul>
			<li><a href="/tags.html">Tags</a></li>
			<li><a href="/sitemap.html">Sitemap</a></li>
			<li><a href="/rss.xml">RSS Feed</a></li>
		</ul>
	</nav>
</footer>

<script src="/theme/jquery.js"></script>
<script src="/theme/autoupdate.js"></script>
<script src="/theme/pdf.js"></script>
<script src="/theme/pdf-preview.js"></script>
<script src="/theme/codehighlighting/highlight.js"></script>
<script>hljs.initHighlightingOnLoad();</script>
<script src="/theme/deck.js"></script>
<script src="/theme/presentation.js"></script>

</body>
</html>`, ChildTemplatePlaceholder)

const repositoryTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>

<div class="cleaner"></div>

<section class="childs">
<ol class="list">
{{range .Childs}}
<li class="child">
	<a href="{{.RelativeRoute}}" class="child-title child-link">{{.Title}}</a>
	<p class="child-description">{{.Description}}</p>
</li>
{{end}}
</ol>
</section>
`

const documentTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>

<div class="cleaner"></div>

{{ if .Tags }}
<section class="tags">
	<header>
		Tags:
	</header>

	<ul class="tags">
	{{range .Tags}}
	<li class="tag">
		<a href="{{.AbsoluteRoute}}" title="{{.Description}}">{{.Name}}</a>
	</li>
	{{end}}
	</ul>
</section>
{{end}}

<section class="childs">
<ol class="list">
{{range .Childs}}
<li class="child">
	<a href="{{.RelativeRoute}}" class="child-title child-link">{{.Title}}</a>
	<p class="child-description">{{.Description}}</p>
</li>
{{end}}
</ol>
</section>
`

const presentationTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<nav>
	<div class="nav-element pager deck-status">
		<span class="deck-status-current"></span> /	<span class="deck-status-total"></span>
	</div>

	<div class="nav-element controls">
		<button class="deck-prev-link" title="Previous">&#8592;</button>
		<button href="#" class="deck-next-link" title="Next">&#8594;</button>
	</div>

	<div class="nav-element jumper">
		<form action="." method="get" class="goto-form">
			<label for="goto-slide">Go to slide:</label>
			<input type="text" name="slidenum" id="goto-slide" list="goto-datalist">
			<datalist id="goto-datalist"></datalist>
			<input type="submit" value="Go">
		</form>
	</div>
</nav>

<section class="content">
{{.Content}}
</section>

<div class="cleaner"></div>
`

const messageTemplate = `
<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>
`

const errorTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>
`

var sitemapContentTemplate = fmt.Sprintf(`
<li>
	<a href="{{.AbsoluteRoute}}" {{ if .Description }}title="{{.Description}}"{{ end }}>{{.Title}}</a>

	{{ if .Childs }}	
	<ol>
	%s
	</ol>
	{{ end }}
</li>`, ChildTemplatePlaceholder)

const sitemapTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
<ol>
{{.Content}}
</ol>
</section>
`

var tagmapContentTemplate = `
<ul class="tags">
{{range .Tags}}
<li class="tag">
	<a name="{{.Name}}" title="{{.Description}}">{{.Name}}</a>
	<ol class="childs">
		{{range .Childs}}
		<li class="child">
			<a href="{{.RelativeRoute}}" class="child-title child-link">{{.Title}}</a>
			<p class="child-description">{{.Description}}</p>
		</li>
		{{end}}
	</ol>
</li>
{{end}}
</ul>
`

const tagmapTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>
`
