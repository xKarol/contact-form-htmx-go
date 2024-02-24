# Contact Form

## Prerequisites

Ensure that you have the following installed on your system:

- Go
- Node.js
- npm
- Makefile

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/xkarol/contact-form-htmx-go.git
cd contact-form-htmx-go
```

2. Create .env file

3. Init project

```bash
make init
```

## Development

1. If you are running sendgrid service locally, start docker compose

```bash
docker compose up
```

2. Run the Go server and watch for changes

```bash
make dev
```

3. Watch for CSS changes:

```bash
make css-watch
```

4. Watch for .templ files changes:

```bash
make templ-watch
```
