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
			name:     "only relative urls",
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<a href="/path/foo">
							<span>Boot.dev</span>
						</a>
						<a href="/path/bar">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			expected: []string{"https://blog.boot.dev/path/foo", "https://blog.boot.dev/path/bar"},
		},
		{
			name:     "only absolute urls",
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<a href="https://example.com">
							<span>Boot.dev</span>
						</a>
						<a href="https://foobar.com">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			expected: []string{"https://example.com", "https://foobar.com"},
		},
		{
			name:     "deeply nested urls",
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<div class="some-container">
							<a href="/path/foo">
								<span>Boot.dev</span>
							</a>
						</div>
						<div class="some-other-container">
							<div class="inner-container">
								<a href="https://example.com">
									<span>Boot.dev</span>
								</a>
							</div>
						</div>
						<a href="https://foobar.com">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			expected: []string{"https://blog.boot.dev/path/foo", "https://example.com", "https://foobar.com"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URLs: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
