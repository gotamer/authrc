#!/usr/bin/env bash

# Tag first -> git tag -a v0.0.10 -m "my tag 10"

set -e

APPNAME='authrc'
PACKAGE="github.com/gotamer/authrc/bin/authrc"


cd $( dirname -- "$0"; );
ROOTDIR=$PWD;

# We don't delete files, we back them up to the trash bin
mkdir -p '/tmp/trash'

# Remember Go env settings, so we can set them back
GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
GOMOD=$(go env GO111MODULE)
GOBIN=$(go env GOBIN)

clear
echo "Hello ${USER}"

VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"

COMMIT_HASH="$(git rev-parse --short HEAD)"

BUILD_TIME=$(date '+%Y-%m-%dT%H:%M:%S')

# Set date and time of the latest commit unless already set.
COMMIT_TIME="${SOURCE_DATE_EPOCH:-$( git log -1 --pretty=%ct )}"

BUILD_USER=$(id -u -n)

readonly COMMIT_HASH
readonly APPNAME
readonly PACKAGE
readonly VERSION
readonly BUILD_TIME
readonly COMMIT_TIME

echo "Current folder: ${ROOTDIR}"
echo "VERSION:" $VERSION

# echo "FLAGS: " ${LDFLAGS[*]}
echo "[INF] make & install ${APPNAME}"

echo "[INF] setting build env"
go env -w GO111MODULE=off CGO_ENABLED=0

echo "[INF] building ${APPNAME}"
cd "${ROOTDIR}/bin/${APPNAME}"

go fmt .
wait

LDFLAGS=("-s -w "
  "-X '${PACKAGE}/build.AppName=${APPNAME}'"
  "-X '${PACKAGE}/build.Version=${VERSION}'"
  "-X '${PACKAGE}/build.CommitHash=${COMMIT_HASH}'"
  "-X '${PACKAGE}/build.BuildTime=${BUILD_TIME}'"
  "-X '${PACKAGE}/build.UserName=${BUILD_USER}'"
)

go build -ldflags="${LDFLAGS[*]}"

wait

	if [[ $USER == "root" ]]
	then
		mv "${APPNAME}" "/sbin/"
		echo "[INF] installed at:"
		which ${APPNAME}
	elif [[ -d $GOBIN ]]
	then
		mv "${APPNAME}" "${GOBIN}/"
		echo "[INF] installed at:"
		which ${APPNAME}
	else
		echo "[ERR] could not move to bin"
	fi

# executes the function $1 $2 $3 $4 $5 are the arguments
$@

wait

echo "[INF] setting Go env back to:"
echo "[INF] GOARCH=${GOARCH} GOOS=${GOOS} GO111MODULE=${GOMOD}"
go env -w GOARCH=${GOARCH} GOOS=${GOOS}
go env -w GO111MODULE=${GOMOD}

echo "[INF] all done"

