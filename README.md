# EzBlogGo
A very simple and lightweight blogging software written in Go.

## What's EzBlog?

EzBlogGo is the implementation of EzBlog in the Go programming language. Why do you ask? Because Go is much faster than any Bash script could be, because it can be compiled.

Compiles markdown blog entries to html and presents it to the user. It can be configured with an own stylesheet. Some of it's features are:

* Serve a light weight blog
* Writing articles using the command line (for example via SSH)
* Using markdown syntax for writing articles
* Easy customization to the css stylesheet
* Many configuration options
* Generate an RSS feed

# Installation and Configuration

## Installing

To compile it yourself, clone the repo into a folder and run 

```
go build
```

And fix all the dependencies by running

```
go get PACKAGENAME
```

... or download a binary from the release page.
