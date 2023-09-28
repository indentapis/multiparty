#!/usr/bin/env bash

set -euo pipefail

SUMS="$(dirname "$(readlink -f "$0")")/get.sha256"
OUTDIR="$(dirname "$(readlink -f "$0")")/../_proto"

arch="$(uname -m)"
kernel="$(uname -s)"
kernelLower="$(tr '[:upper:]' '[:lower:]' <<< "$kernel")"

os="$kernelLower"
if [[ "$kernelLower" == "darwin" ]]; then
  os="osx"
fi

# write to temp file and cleanup later
tmpDir="$(mktemp -d)"
function cleanup() {
	rm -rf ${tmpDir}
}
trap cleanup EXIT

function main() {
  for name in "$@"; do
    case $name in
      buf)
        version=1.26.1
        target="buf-${kernel}-${arch}"
        url="https://github.com/bufbuild/buf/releases/download/v${version}/${target}"
        download "$url" "$tmpDir/$target"
        chmod +x "$tmpDir/$target"

        mv "$tmpDir/$target" "${OUTDIR}/buf"
        ;;

      protoc)
        if [[ "$arch" == "arm64" ]]; then
          arch="aarch_64"
        fi

        version=24.3
        target="protoc-${version}-${os}-${arch}.zip"
        url="https://github.com/protocolbuffers/protobuf/releases/download/v${version}/${target}"
        download "$url" "$tmpDir/$target"

        mkdir -p "${OUTDIR}/protoc"
        unzip -qq -d "${OUTDIR}/protoc" "$tmpDir/$target"
        ;;

      protoc-gen-go)
        if [[ "$kernelLower" == "darwin" ]]; then
          os="darwin"
        fi

        version=v1.31.0
        target="protoc-gen-go.${version}.${os}.${arch}.tar.gz"
        url="https://github.com/protocolbuffers/protobuf-go/releases/download/${version}/${target}"
        download "$url" "$tmpDir/$target"

        tar xvf "$tmpDir/$target" -C "${OUTDIR}" protoc-gen-go
        ;;

      *)
        echo "Unknown: ${name}"
        exit 1
        ;;

    esac
    echo "got ${name}!"

  done
}

function download() {
	local url="${1}"
	local dst="${2}"

	if [[ -z ${url} || -z ${dst} ]]; then
		echo "URL and Destination must be provided like: ${0} <url> <dst> [checksum]"
		exit 3
	elif [[ -f ${dst} ]]; then
		# check if already have copy
		check ${dst} && exit 0
	fi

	echo "Downloading ${url}..."
	curl -L -o ${dst} "${url}"

	echo "Verifying checksum of ${dst}"
	check ${dst} || (echo "Failed checksum!" && exit 2)
	echo "Checksum passed!"
}

function check() {
  local target="${1}"
	local checksum=$(grep $(basename ${target}) ${SUMS})

  # if no checksum provided, add to file
	if [[ -z ${checksum} ]]; then
    checksum=$(cd $(dirname ${target}) && sha256sum $(basename ${target}))

    if [[ -t 0 ]]; then
      printf "Checksum not found for ${target}.\nChecksum: ${checksum}\n"
      read -p "Do you want to add it? (Y/n) " response
      case ${response} in
        [nN]|[nN][oO]) return false ;;
      esac
    else
      return false
    fi

    echo "Adding checksum: ${checksum}"
    local newSums=$(echo ${checksum} | cat ${SUMS} - | sort -k2 | uniq)
    echo "${newSums}" > ${SUMS}
	fi

	cd $(dirname ${target}) && sha256sum --check --status <<< "${checksum}"
}

main "$@"
