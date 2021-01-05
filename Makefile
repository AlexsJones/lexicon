all:
	gox -output="build/lexicon_{{.OS}}_{{.Arch}}" -os="linux/amd64" -osarch="darwin/amd64"
