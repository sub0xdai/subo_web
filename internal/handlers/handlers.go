package handlers

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/maragudk/gomponents"
	"github.com/sub0x/subo_web/internal/components"
)

func renderPage(node gomponents.Node) (string, error) {
	var buf bytes.Buffer
	buf.WriteString("<!DOCTYPE html>\n")
	if err := node.Render(&buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func HandleHome(c *fiber.Ctx) error {
	props := components.BaseLayoutProps{
		Title:       "SUBO - Home",
		Description: "Personal portfolio and blog of SUBO",
		Children: []gomponents.Node{
			gomponents.El("div",
				gomponents.Attr("class", "hero"),
				gomponents.El("h1", 
					gomponents.Attr("class", "hero-title"),
					gomponents.Text("Welcome to SUBO's Portfolio"),
				),
				gomponents.El("p", 
					gomponents.Attr("class", "hero-subtitle"),
					gomponents.Text("Software Engineer & Creative Developer"),
				),
			),
		},
	}

	html, err := renderPage(components.BaseLayout(props))
	if err != nil {
		return err
	}
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(html)
}

func HandleProjects(c *fiber.Ctx) error {
	projects := []components.Project{
		{
			Title:       "Project 1",
			Description: "Description of project 1",
			ImageURL:    "/static/images/project1.jpg",
			Link:        "/projects/1",
			Tags:        []string{"Go", "HTMX", "Fiber"},
		},
		// Add more projects as needed
	}

	props := components.BaseLayoutProps{
		Title:       "SUBO - Projects",
		Description: "View my latest projects and work",
		Children:    []gomponents.Node{components.ProjectsGrid(projects)},
	}

	html, err := renderPage(components.BaseLayout(props))
	if err != nil {
		return err
	}
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(html)
}

func HandleBlog(c *fiber.Ctx) error {
	props := components.BaseLayoutProps{
		Title:       "SUBO - Blog",
		Description: "My thoughts and writings on technology",
		Children: []gomponents.Node{
			gomponents.El("div",
				gomponents.Attr("class", "blog-container"),
				gomponents.El("h1", gomponents.Text("Blog")),
				gomponents.El("p", gomponents.Text("Coming soon...")),
			),
		},
	}

	html, err := renderPage(components.BaseLayout(props))
	if err != nil {
		return err
	}
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(html)
}

func HandleAbout(c *fiber.Ctx) error {
	props := components.BaseLayoutProps{
		Title:       "SUBO - About",
		Description: "Learn more about me and my work",
		Children: []gomponents.Node{
			gomponents.El("div",
				gomponents.Attr("class", "about-container"),
				gomponents.El("h1", gomponents.Text("About Me")),
				gomponents.El("p", gomponents.Text("Software engineer passionate about building beautiful and functional web applications.")),
			),
		},
	}

	html, err := renderPage(components.BaseLayout(props))
	if err != nil {
		return err
	}
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(html)
}
