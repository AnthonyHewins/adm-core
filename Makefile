targets := polynomial-regression
.PHONY: build clean deploy

$(targets):
	GOOS=linux go build -ldflags="-s -w" -o bin/$@ cmd/$@/main.go

all: $(targets)

clean:
	find . -iname "*.go" -exec gofmt -w -s {} \;
	go mod download
	rm -rf ./bin

deploy: clean $(targets)
	sls deploy --verbose -f $(targets)
