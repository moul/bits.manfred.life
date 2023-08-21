BASEURL ?= http://localhost
run:
	hugo server --bind=0.0.0.0 --enableGitInfo --baseURL=$(BASEURL) --buildDrafts --buildFuture --contentDir=.. --disableFastRender

build:
	hugo --enableGitInfo --buildDrafts --contentDir=..

deps:
	git submodule add -f https://github.com/moul/hugo-theme-techdoc themes/hugo-theme-techdoc
