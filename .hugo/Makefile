run: symlinks
	hugo server --bind=0.0.0.0 --enableGitInfo --theme=hugo-theme-techdoc

symlinks:
	cd content; ln -s ../../* . 2>/dev/null || true

build: symlinks
	hugo --enableGitInfo
