# Personal Portfolio Web Application

A modern, minimalist portfolio website built with Go, HTMX, and Tailwind CSS.

## Tech Stack

- Go (Backend)
- HTMX (Interactive UI)
- Gomponents (Go HTML Components)
- Fiber (Web Framework)
- Tailwind CSS (Styling)
- SQLite (Database)

## Features

- AI Chatbot (Terminal-style)
- Project Showcase
- Personal Blog
- About Section
- Responsive Navigation
- Catpuccin color theme
- Dark/Light Mode Toggle

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Node.js and npm (for Tailwind CSS)

### Installation

1. Clone the repository
```bash
git clone https://github.com/yourusername/subo_web.git
cd subo_web
```

2. Install Go dependencies
```bash
go mod download
```

3. Run the application
```bash
go run cmd/main.go
```

The application will be available at `http://localhost:3000`

## Project Structure

```
subo_web/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── components/             # Reusable UI components
│   ├── database/               # Database operations
│   ├── handlers/               # Request handlers
│   ├── middleware/             # Custom middleware
│   ├── models/                 # Data models
│   └── utils/                  # Utility functions
└── web/
    ├── templates/              # HTML templates
    └── static/                 # Static assets
        ├── css/               # Stylesheets
        ├── js/                # JavaScript files
        └── images/            # Image assets
```

## Design Philosophy

- Minimal and clean design
- Gradient & geometric elements
- Fibonacci-inspired aesthetics
- Focus on performance and user experience
