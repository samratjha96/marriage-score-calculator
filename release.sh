#!/usr/bin/env bash

set -euxo pipefail

# Based on https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
platforms=("windows/amd64" "darwin/amd64" "linux/amd64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	goos=${platform_split[0]}
	goarch=${platform_split[1]}
	output_name=marriage-$goos-$goarch
	if [ $goos = "windows" ]; then
		output_name=$output_name.exe
	fi

	CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -ldflags '-w -extldflags "-static"' -o binaries/$output_name .
done

