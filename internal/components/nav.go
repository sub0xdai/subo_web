package components

import (
	g "github.com/maragudk/gomponents"
)

func NavBar() g.Node {
	return g.El("nav",
		g.Attr("class", "nav"),
		g.El("div",
			g.Attr("class", "nav-container"),
			g.El("a",
				g.Attr("href", "/"),
				g.Attr("class", "nav-link"),
				g.Text("SUBO"),
			),
			g.El("div",
				g.Attr("class", "nav-links"),
				navLink("/", "Home"),
				navLink("/projects", "Projects"),
				navLink("/blog", "Blog"),
				navLink("/about", "About"),
			),
			g.El("div",
				g.Attr("class", "nav-controls"),
				g.El("button",
					g.Attr("type", "button"),
					g.Attr("class", "menu-toggle"),
					g.Raw(`<svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path></svg>`),
				),
				g.El("button",
					g.Attr("type", "button"),
					g.Attr("id", "theme-toggle"),
					g.Attr("class", "theme-toggle"),
					g.Raw(`<svg class="theme-toggle-icon" width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
						<path class="sun-icon" d="M12 2v2m0 16v2M4 12H2m20 0h-2m-2.05-6.95L16.24 7.76M7.76 16.24l-1.71 1.71M7.76 7.76L6.05 6.05m10.19 10.19l1.71 1.71M12 17a5 5 0 100-10 5 5 0 000 10z"/>
						<path class="moon-icon" d="M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z"/>
					</svg>`),
				),
			),
		),
	)
}

func navLink(href, text string) g.Node {
	return g.El("a",
		g.Attr("href", href),
		g.Attr("class", "nav-link"),
		g.Text(text),
	)
}
