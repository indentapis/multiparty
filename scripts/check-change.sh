#!/bin/bash -e

GITHUB_SHA=${GITHUB_SHA:-$(git rev-parse HEAD)}

# Check that file matches the version in Git
function main() {
	local files=${@}
	for f in ${files}; do
		compare_git "${GITHUB_SHA}" ${f} || (echo "Uncommitted changes in ${f}." >&2; exit 1)
	done
}

function compare_git() {
	local sha="${1}"
	local path="${2}"
	
	echo "Checking ${path} is up to date for ${sha}"
	git show ${sha}:${path} | diff ${path} -
}

main "$@"
