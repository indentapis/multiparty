#!/bin/bash -e

# Copyright 2023 Indent Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Check that file matches version committed in Git.
# Compared against the current HEAD, or the commit specified by GITHUB_SHA.
# Usage: scripts/check-change.sh <file>

GITHUB_SHA=${GITHUB_SHA:-$(git rev-parse HEAD)}

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
