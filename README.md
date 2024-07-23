## LeapKit Template

<img width="300" alt="logo" src="https://leapkit.dev/assets/logo.svg">
<br><br>

This is the  LeapKit template for building web applications with Go, HTMX and Tailwind CSS. It integrates useful features such as hot code reload and css recompiling.

### Getting started


### Setup

Install dependencies:

```sh
go mod download
go run ./cmd/setup
```

### Running in dev mode

Install LeapKit CLI. You can run the following command
```bash
go install github.com/leapkit/leapkit/kit@latest
```

Now you can use the following command to run the app in dev mode

```bash
kit dev
```

And open `http://localhost:3001` in your browser.
