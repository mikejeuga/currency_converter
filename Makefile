repo=$(shell basename "`pwd`")

gopher:
	@git init
	@touch .gitignore
	@touch README.md
	@touch .envrc
	@direnv allow
	@mkdir cmd
	@touch ./cmd/main.go
	@go mod init github.com/mikejeuga/$(repo)
	@go get github.com/google/uui
	@go get github.com/adamluzsi/testcase
	@go get github.com/gorilla/mux
	@go mod tidy

run:
	@go run ./cmd/main.go

t: test
test:
	@make at ut


ut: unit-test
unit-test:
	@go test -v --tags=unit ./...

at: acceptance-test
acceptance-test:
	@go test -v -tags=acceptance ./blackboxtests/...

ic: init
init:
	@gh repo create ${repo} --private
	@git add .
	@git commit -m "Initial commit"
	@git remote add origin git@github.com:mikejeuga/${repo}.git
	@git branch -M main
	@git push -u origin main

c: commit
commit:
	@git add .
	@git commit -m "$m"
	@git pull --rebase
	git push

privacy:
	@gh repo edit --visibility=private

public:
	@gh repo edit --visibility=public

destroy:
	rm -rf .git
	@gh repo delete ${repo}