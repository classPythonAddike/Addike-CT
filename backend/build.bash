#!/usr/bin/env bash

platforms=("windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64" "linux/amd64" "linux/386")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    executable='addike-ct'
	output_name='bin/'$platform'/'$executable''

	if [ $GOOS = "windows" ]; then
        output_name+='.exe'
        executable+='.exe'
    fi

    echo "Building for $platform"

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name .
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

	cp "$CTPath" "bin/$platform/public/" -r

    echo "Zipping for $platform"
    zipname="addikect-$GOOS-$GOARCH.zip"

	cd "bin/$platform/"
	zip -rq "$zipname" "./$executable" "public"
	cd ../../..
    cp "bin/$platform/$zipname" "bin/"

    echo "Cleaning up"

    rm "bin/$GOOS" -r

    echo "==================================================================="

done

echo "Done compiling project!"