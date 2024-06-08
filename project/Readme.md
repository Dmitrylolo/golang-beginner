Beginner's First Test Project
==========================

## Running Locally

To run the project locally, you'll need to have Go installed. Once you have Go installed, run the following commands:

```bash
$ git clone https://github.com/jmccann/beginners-first-test-project.git
$ cd beginners-first-test-project
$ go run cmd/cli/main.go
```

## Build Applicaiton

To build the application, run the following command:

Macos/Linux:
```bash
$ go build -o cli cmd/cli/main.go
```

Windows:
```bash
$ go build -o cli.exe cmd/cli/main.go
```
