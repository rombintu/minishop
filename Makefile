run:
	go run server/cmd/main.go -config-path config/config.toml

build:
	go build server/cmd/main.go

push:
	git add .
	git commit -m "auto git push"
	git push -u git@github.com:rombintu/minishop.git master