default:
	@just --choose

bin := "./bin"
app_name := "my-go-app"
app := bin / app_name
src := "."

build:
	mkdir -p {{bin}}
	go build -o {{app}} {{src}}

test:
	go test -v ./...

coverage:
	go test -coverprofile cover.cov ./...

open-coverage:
	go tool cover -html=cover.cov

# golangci-lint fix linting errors and format if possible
fix:
    golangci-lint run --fast --fix

upgrade:
    git status --porcelain | grep -q . && echo "Repository is dirty, commit changes before upgrading." && exit 1 || exit 0
    go get -u ./...
    go mod tidy

release-snapshot:
	goreleaser release --snapshot --clean

# release with goreleaser
# Needs GITHUB_TOKEN defined
release:
	goreleaser release --clean

clean:
	rm -rf {{bin}}
