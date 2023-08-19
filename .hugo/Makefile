BASEURL ?= http://localhost
run:
	hugo server --bind=0.0.0.0 --enableGitInfo --theme=hugo-theme-techdoc --baseURL=$(BASEURL) --buildDrafts --buildFuture --contentDir=..

build:
	hugo --enableGitInfo --contentDir=..

deps:
	git submodule add -f https://github.com/moul/hugo-theme-techdoc themes/hugo-theme-techdoc
