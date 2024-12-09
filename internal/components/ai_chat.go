package components

import (
	g "github.com/maragudk/gomponents"
)

func AIChat() g.Node {
	return g.Group([]g.Node{
		g.El("div",
			g.Attr("class", "ai-chat-container"),
			g.Attr("aria-expanded", "true"),
			g.El("div",
				g.Attr("class", "ai-chat-window"),
				g.El("div",
					g.Attr("class", "ai-chat-header"),
					g.El("span", g.Text("AI Assistant")),
					g.El("button",
						g.Attr("class", "ai-chat-close"),
						g.Attr("aria-label", "Close chat"),
						g.Text("×"),
					),
				),
				g.El("div",
					g.Attr("class", "ai-chat-messages"),
					g.Attr("id", "chat-messages"),
				),
				g.El("form",
					g.Attr("class", "ai-chat-input"),
					g.Attr("hx-post", "/api/chat"),
					g.Attr("hx-target", "#chat-messages"),
					g.Attr("hx-swap", "beforeend"),
					g.El("input",
						g.Attr("type", "text"),
						g.Attr("name", "message"),
						g.Attr("placeholder", "Type your message..."),
						g.Attr("autocomplete", "off"),
					),
					g.El("button",
						g.Attr("type", "submit"),
						g.Text("Send"),
					),
				),
			),
		),
		g.El("div",
			g.Attr("class", "ai-chat-backdrop"),
		),
	})
}
