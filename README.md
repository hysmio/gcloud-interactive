# GCloud Interactive

**GCloud Interactive** is an interactive CLI Application for communicating with GCloud.

This was developed to speed up my own deployment to GCloud but feel free to contribute & make improvements to this tool.

## Installation

### Prerequisites
[Install and setup Go](https://golang.org/doc/install)

[Install and initialize gcloud](https://cloud.google.com/sdk/docs/quickstarts)

**Make sure to run `gcloud init` and sign in, this tool will not work without it.**

**This tool currently doesn't support OAuth2 connections and just relies on the `gcloud` binary**

### Install & build
Using `go get`

```bash
# Get and build the project
go get github.com/hysmio/gcloud-interactive
```

Using `git`
```bash
# Clone the repo
git clone https://github.com/hysmio/gcloud-interactive

# Change to the new directory
cd gcloud-interactive

# Build the project
go build
```

## Usage

```bash
gcloud-interactive
```

## Planned Features
- **OAuth2 Authentication**
- **Remove need for `gcloud` binary**

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)
