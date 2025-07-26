package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := map[string]struct {
		html     string
		inputURL string
		expected []string
	}{
		"absolute and relative urls": {
			html: `
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
			inputURL: "https://blog.boot.dev",
			expected: []string{
				"https://blog.boot.dev/path/one",
				"https://other.com/path/one",
			},
		},
		"self contained url": {
			html: `
			<html>
				<body>
					<a href="https://blog.boot.dev/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			inputURL: "https://blog.boot.dev",
			expected: []string{
				"https://blog.boot.dev/path/one",
			},
		},
		"urls nested into list and other tags": {
			html: `
			<html>
				<body>
					<div>
						<ul>
							<li>
								<article>
									<header>
										<h2>
											<a href="/post/1">
												<span><strong>Read</strong> this post</span>
											</a>
										</h2>
									</header>
								</article>
							</li>
							<li>
								<section>
									<aside>
										<p>
											Check this:
											<a href="https://external.dev/page">
												<em><span>External Page</span></em>
											</a>
										</p>
									</aside>
								</section>
							</li>
							<li>
								<footer>
									<small>
										<a href="/contact">
											<span><i>Contact us</i></span>
										</a>
									</small>
								</footer>
							</li>
							<li>
								<nav>
									<ol>
										<li>
											<a href="/help/docs/start">
												<span>
													<i>
														<code>Start Here</code>
													</i>
												</span>
											</a>
										</li>
									</ol>
								</nav>
							</li>
						</ul>
					</div>
				</body>
			</html>
			`,
			inputURL: "https://blog.boot.dev",
			expected: []string{
				"https://blog.boot.dev/post/1",
				"https://external.dev/page",
				"https://blog.boot.dev/contact",
				"https://blog.boot.dev/help/docs/start",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := getURLsFromHTML(tc.html, tc.inputURL)
			if err != nil {
				t.Errorf("Test '%s' FAIL: unexpected error: %v", name, err)
				return
			}
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}
