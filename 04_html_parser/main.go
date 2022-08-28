package main

import (
	"fmt"
	"strings"

	"github.com/mathiashandle/04htmlparser/links"
)

var exampleHtml = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Document</title>
	</head>
	<body>
		<h1>Hey this is the biggest title</h1>

		<nav>
			<ul>
				<li>
					<a href="/">Home</a>
					<a href="/blog">Blog</a>
					<a href="/about">About</a>
				</li>
			</ul>
		</nav>

		<section>
			<article>
				<h2>Smaller title</h2>
				<p>
					Lorem ipsum dolor sit, amet consectetur adipisicing elit. In odio facilis quaerat illum blanditiis veniam libero ea, magni magnam quos. Ad
					libero dicta maxime eaque unde tempora nisi iste quos!

					<a href="/article-slug">article name</a>
				</p>
				<p>
					Lorem ipsum dolor sit, amet consectetur adipisicing elit. In odio facilis quaerat illum blanditiis veniam libero ea, magni magnam quos. Ad
					libero dicta maxime eaque unde tempora nisi iste quos!
					<a href="/article-slug">
						<!-- sdada -->
						article name
					</a>
				</p>
			</article>
		</section>
	</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := links.Parse(r)
	if err != nil {
		fmt.Printf("error parsing file")
	}

	fmt.Print(links)
}
