.PHONY: clone dependencies echo-open prepare

UNAME := $(shell uname)

help:
	@grep -E '(^[0-9a-zA-Z_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-25s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

clone: ## Clone the repository to the right place
	mkdir -p $GOPATH/src/github.com/darkweak/golang-starter || true
	git clone https://github.com/Darkweak/golang-starter $GOPATH/src/github.com/darkweak/golang-starter

dependencies: ## Install and check dependencies
    ifeq ($(UNAME), Linux)
		wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash
    endif
    ifeq ($(UNAME), Darwin)
		curl https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash
    endif

echo-open: ## Echo open the repository
	echo 'Please open the repository cloned at "$(echo "$GOPATH")/src/github.com/darkweak/golang-starter" '

echo-source: ## Echo source the bashrc or zshrc
	echo 'Please run `source ~/.bashrc` or `source ~/.zshrc` if using zsh'

prepare: dependencies clone ## Prepare to work
	echo 'Setup is done!'
	$(MAKE) echo-source
	$(MAKE) echo-open
