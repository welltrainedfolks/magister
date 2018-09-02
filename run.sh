#!/bin/bash

version="0.0.1"
build=$(git rev-list --all --count)
builddate=$(date +'%Y-%m-%d %H:%M:%S')
revision=$(git rev-list --all | head -n 1)
branch=$(git branch | grep "*" | awk {' print $2 '})

OS=$(uname -s)

if [ "${OS}" != "Linux" -a "${OS}" != "Darwin" ]; then
    echo "Running under anything except linux and macos is not yet supported. Patches welcome!"
    exit 1
else
    READLINK="/bin/readlink"
    if [ "${OS}" == "Darwin" ]; then
        READLINK="/usr/local/bin/greadlink"
    fi
    echo "Try to use readlink from ${READLINK}"

    if [ ! -f "${READLINK}" ]; then
        echo "GNU coreutils should be installed."
        exit 2
    fi
    SCRIPT_PATH=$(dirname "`${READLINK} -f "${BASH_SOURCE}"`")
fi

GO=$(which go)
if [ "${#GO}" -eq 0 ]; then
    echo "Golang compiler is not installed. Cannot continue."
    exit 2
fi

function generate_version() {
    echo -e "package common
    
const (
    VERSION = \"${version}\"
    BUILD = ${build}
    BUILDDATE = \"${builddate}\"
    REVISION = \"${revision}\"
    BRANCH = \"${branch}\"
)" > "${SCRIPT_PATH}/common/version.go"
}

function ctl() {
    echo "Starting magisterctl version ${version} with CLI parameters: \"$@\""
    generate_version
    $GO run -race cmd/magisterctl/main.go $@
}

function run() {
    echo "Starting MAGISTER version ${version} with CLI parameters: \"$@\""
    generate_version
    $GO run -race cmd/magister/main.go $@
}

case $1 in
    ctl)
        shift
        ctl $@
    ;;
    run)
        shift
        run $@
    ;;
esac