package components

import (
	g "github.com/maragudk/gomponents"
)

type Project struct {
	Title       string
	Description string
	ImageURL    string
	Link        string
	Tags        []string
}

func ProjectsGrid(projects []Project) g.Node {
	var projectNodes []g.Node
	for _, project := range projects {
		projectNodes = append(projectNodes, projectCard(project))
	}
	
	return g.El("div", append([]g.Node{
		g.Attr("id", "projects-grid"),
		g.Attr("class", "projects-grid"),
	}, projectNodes...)...)
}

func projectCard(project Project) g.Node {
	var tagNodes []g.Node
	for _, tag := range project.Tags {
		tagNodes = append(tagNodes, 
			g.El("span", 
				g.Attr("class", "project-tag"),
				g.Text(tag),
			),
		)
	}

	return g.El("article", 
		g.Attr("class", "project-card"),
		g.El("a", 
			g.Attr("href", project.Link),
			g.El("img", 
				g.Attr("src", project.ImageURL),
				g.Attr("alt", project.Title),
				g.Attr("class", "project-image"),
			),
			g.El("div", 
				g.Attr("class", "project-content"),
				g.El("h3", 
					g.Attr("class", "project-title"),
					g.Text(project.Title),
				),
				g.El("p", 
					g.Attr("class", "project-description"),
					g.Text(project.Description),
				),
				g.El("div", 
					append([]g.Node{g.Attr("class", "project-tags")}, tagNodes...)...,
				),
			),
		),
	)
}
