package components

import (
	g "github.com/maragudk/gomponents"
)

type BaseLayoutProps struct {
	Title       string
	Description string
	Children    []g.Node
}

func BaseLayout(props BaseLayoutProps) g.Node {
	charset := g.El("meta", g.Attr("charset", "utf-8"))
	viewport := g.El("meta", 
		g.Attr("name", "viewport"),
		g.Attr("content", "width=device-width, initial-scale=1"),
	)
	description := g.El("meta",
		g.Attr("name", "description"),
		g.Attr("content", props.Description),
	)
	title := g.El("title", g.Text(props.Title))
	faviconICO := g.El("link",
		g.Attr("rel", "icon"),
		g.Attr("type", "image/x-icon"),
		g.Attr("href", "/static/images/favicon.ico"),
	)
	faviconSVG := g.El("link",
		g.Attr("rel", "icon"),
		g.Attr("type", "image/svg+xml"),
		g.Attr("href", "/static/images/favicon.svg"),
	)
	stylesheet := g.El("link",
		g.Attr("rel", "stylesheet"),
		g.Attr("href", "/static/css/style.css"),
	)
	htmx := g.El("script",
		g.Attr("src", "https://unpkg.com/htmx.org@1.9.6"),
		g.Attr("defer", ""),
	)

	head := g.El("head",
		charset,
		viewport,
		description,
		title,
		faviconICO,
		faviconSVG,
		stylesheet,
		htmx,
		g.Raw(`<script>
			function setTheme(theme) {
				document.documentElement.setAttribute('data-theme', theme);
				localStorage.setItem('theme', theme);
			}
			const savedTheme = localStorage.getItem('theme') || 'light';
			setTheme(savedTheme);
		</script>`),
	)

	mainNodes := append([]g.Node{g.Attr("class", "container")}, props.Children...)
	body := g.El("body",
		NavBar(),
		g.El("main", mainNodes...),
		AIChat(),
		g.El("script",
			g.Raw(`
				document.getElementById('theme-toggle').addEventListener('click', function() {
					const currentTheme = localStorage.getItem('theme') || 'light';
					const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
					setTheme(newTheme);
				});
			`),
		),
	)

	return g.El("html", 
		g.Attr("lang", "en"),
		head,
		body,
	)
}
