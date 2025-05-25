package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
        {
			name:     "remove http scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "handle uppercase",
			inputURL: "https://BLOG.BOOT.DEV/PATH",
			expected: "blog.boot.dev/PATH",
		},
		{
			name:     "handle port",
			inputURL: "https://blog.boot.dev:8080/path",
			expected: "blog.boot.dev:8080/path",
		},
		{
			name:     "handle query params",
			inputURL: "https://blog.boot.dev/path?query=1",
			expected: "blog.boot.dev/path?query=1",
		},
		{
			name:     "handle fragment",
			inputURL: "https://blog.boot.dev/path#section",
			expected: "blog.boot.dev/path#section",
		},
		{
			name:     "handle root path",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev/",
		},
		{
			name:     "handle root path no slash",
			inputURL: "https://blog.boot.dev",
			expected: "blog.boot.dev",
		},
		{
			name:     "handle complex path",
			inputURL: "https://blog.boot.dev/path/to/resource/",
			expected: "blog.boot.dev/path/to/resource",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
		