BASEURL ?= http://localhost
run:
	hugo server --bind=0.0.0.0 --enableGitInfo --theme=hugo-theme-techdoc --baseURL=$(BASEURL)

build:
	hugo --enableGitInfo

deps:
	git submodule add -f https://github.com/moul/hugo-theme-techdoc themes/hugo-theme-techdoc
