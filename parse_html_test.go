package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
    tests := []struct {
        name      string
        inputURL  string
        inputBody string
        expected  []string
    }{
        {
            name:     "absolute and relative URLs",
            inputURL: "https://blog.boot.dev",
            inputBody: `
<html>
    <body>
        <a href="/path/one">
            <span>Boot.dev</span>
        </a>
        <a href="https://other.com/path/one">
            <span>Boot.dev</span>
        </a>
    </body>
</html>
`,
            expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
        },
        {
            name:     "relative URLs only",
            inputURL: "https://example.com",
            inputBody: `
<html>
    <body>
        <a href="/about">About</a>
        <a href="/contact">Contact</a>
        <a href="relative/path">Relative</a>
    </body>
</html>
`,
            expected: []string{"https://example.com/about", "https://example.com/contact", "https://example.com/relative/path"},
        },
        {
            name:     "absolute URLs only",
            inputURL: "https://blog.boot.dev",
            inputBody: `
<html>
    <body>
        <a href="https://google.com">Google</a>
        <a href="http://example.com">Example</a>
        <a href="https://github.com/user/repo">GitHub</a>
    </body>
</html>
`,
            expected: []string{"https://google.com", "http://example.com", "https://github.com/user/repo"},
        },
        {
            name:     "no anchor tags",
            inputURL: "https://blog.boot.dev",
            inputBody: `
<html>
    <body>
        <p>This is a paragraph with no links.</p>
        <div>Another div without links</div>
    </body>
</html>
`,
            expected: []string{},
        },
        {
            name:     "empty href attributes",
            inputURL: "https://blog.boot.dev",
            inputBody: `
<html>
    <body>
        <a href="">Empty href</a>
        <a href="/valid">Valid link</a>
        <a>No href attribute</a>
    </body>
</html>
`,
            expected: []string{"https://blog.boot.dev/valid"},
        },
        {
            name:     "base URL with trailing slash",
            inputURL: "https://blog.boot.dev/section/",
            inputBody: `
<html>
    <body>
        <a href="../parent">Parent directory</a>
        <a href="./current">Current directory</a>
        <a href="sub/page">Sub directory</a>
        <a href="/root">Root path</a>
    </body>
</html>
`,
            expected: []string{
                "https://blog.boot.dev/parent",
                "https://blog.boot.dev/section/current",
                "https://blog.boot.dev/section/sub/page",
                "https://blog.boot.dev/root",
            },
        },
        {
            name:     "mixed case and nested tags",
            inputURL: "https://example.org",
            inputBody: `
<HTML>
    <BODY>
        <A HREF="/Path/One">
            <span>Nested content</span>
        </A>
        <a href="https://EXTERNAL.COM/path">External</a>
        <A href="relative">Relative</A>
    </BODY>
</HTML>
`,
            expected: []string{
                "https://example.org/Path/One",
                "https://EXTERNAL.COM/path",
                "https://example.org/relative",
            },
        },
        {
            name:     "well formed HTML only",
            inputURL: "https://test.com",
            inputBody: `
<html>
    <body>
        <a href="/valid1">Valid 1</a>
        <a href="/valid2">Valid 2</a>
        <a href="/valid3">Valid 3</a>
    </body>
</html>
`,
            expected: []string{"https://test.com/valid1", "https://test.com/valid2", "https://test.com/valid3"},
        },
        {
            name:     "query parameters and fragments",
            inputURL: "https://site.com",
            inputBody: `
<html>
    <body>
        <a href="/search?q=test">Search</a>
        <a href="/page#section">Page section</a>
        <a href="https://external.com?param=value#frag">External with params</a>
    </body>
</html>
`,
            expected: []string{
                "https://site.com/search?q=test",
                "https://site.com/page#section",
                "https://external.com?param=value#frag",
            },
        },
        {
            name:     "multiple nested links",
            inputURL: "https://blog.example.com",
            inputBody: `
<html>
    <head><title>Test</title></head>
    <body>
        <nav>
            <a href="/home">Home</a>
            <a href="/about">About</a>
        </nav>
        <main>
            <article>
                <a href="https://external.com">External Link</a>
                <div>
                    <a href="/nested/deep">Deep Link</a>
                </div>
            </article>
        </main>
    </body>
</html>
`,
            expected: []string{
                "https://blog.example.com/home",
                "https://blog.example.com/about", 
                "https://external.com",
                "https://blog.example.com/nested/deep",
            },
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
            if err != nil {
                t.Errorf("Test '%s' FAIL: unexpected error: %v", tc.name, err)
                return
            }
            if len(actual) == 0 && len(tc.expected) == 0 {
                return 
            }
            if !reflect.DeepEqual(actual, tc.expected) {
                t.Errorf("Test '%s' FAIL:\nexpected: %v\nactual: %v", tc.name, tc.expected, actual)
            }
        })
    }
}