# go-frameproto

Package **frameproto** provides tools for the **Frame Protocol** — which is also known as **Farcaster Frames**, for the Go programming language.

The specification for the **Frame Protocol** is at:
https://docs.farcaster.xyz/reference/frames/spec

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-frameproto

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-frameproto?status.svg)](https://godoc.org/sourcecode.social/reiver/go-frameproto)

## Explanation

The **Frames Protocol**, also known **Farcaster Frames**, is a simple web-based technology used for making applications.

It uses HTML without really using HTML, so that **Frames Protocol** applications work with clients that don't support the **Frames Protocol**.
The fall-back being OpenGraph.

Really, a **Frames Protocol** application is mostly made up of **images** and **buttons** on the client-side (that are specified using HTML `<meta>` element) with a back-end that gets HTTP `POST`ed to, which can return a new "page" with an **image** and **buttons**, and so on and so on.

This choice of just being mostly **images** and **buttons**  actually makes the **Frames Protocol** simpler to create a viewer from scatch.
No need to implement all Web technologies.
No need to worry about security and privacy holes that Web technologies introduce.

Although the **Frames Protocol** _could_ be used outside of **Farcaster**, at the time of writing, **Farcaster** clients (such as **Warpcast**) are the only major (client-side) platform to support it.

(The server-side of the <strong>Frames Protocol</strong>, which is called a <strong>Frame Server</strong>, is an just HTTP resource — which some might loosely call an HTTP (or HTTPS) URL.)

Enough talking — let's look at some code.
Here is the client-side of a **Frames Protocol** application:

```html
<meta property="fc:frame" content="vNext" />
<meta property="fc:frame:image" content="https://example.com/path/to/image.png" />
<meta property="og:image" content="https://example.com/path/to/image.png" />
```

It is just HTML.

Although this would need to be embedded into an HTML document, so really it would be something more like this:

```html
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="utf-8" />

<meta property="fc:frame" content="vNext" />
<meta property="fc:frame:image" content="https://example.com/path/to/image.png" />
<meta property="og:image" content="https://example.com/path/to/image.png" />

</head>
<body>
</body>
</html>
```

This package provides you tools for creating this.

For example:

```golang
func ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	// ...

	frameproto.WriteFrame(responseWriter, frameproto.VersionVNext)
	frameproto.WriteFrameImage(responseWriter, frameImageURL)

	// ...
}


```

## Import

To import package **frameproto** use `import` code like the follownig:
```
import "sourcecode.social/reiver/go-frameproto"
```

## Installation

To install package **frameproto** do the following:
```
GOPROXY=direct go get https://sourcecode.social/reiver/go-frameproto
```

## Author

Package **frameproto** was written by [Charles Iliya Krempeaux](http://changelog.ca)
