#!/usr/bin/env bash

set -e

cd $( dirname -- "$0"; );
SCRIPTDIR=$PWD;

# We don't delete files, we back them up to the trash bin
mkdir -p '/tmp/trash'

# Remember Go env settings, so we can set them back
GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
GOMOD=$(go env GO111MODULE)
GOBIN=$(go env GOBIN)

clear
echo "Hello ${USER}"
APPNAME='authrc'
readonly APPNAME

PACKAGE="github.com/gotamer/authrc/bin/authrc"
readonly PACKAGE

VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
readonly VERSION

COMMIT_HASH="$(git rev-parse --short HEAD)"
readonly COMMIT_HASH

BUILD_TIME=$(date '+%Y-%m-%dT%H:%M:%S')
readonly BUILD_TIME

# Set date and time of the latest commit unless already set.
COMMIT_TIME="${SOURCE_DATE_EPOCH:-$( git log -1 --pretty=%ct )}"
readonly COMMIT_TIME

echo "VERSION:" $VERSION

# echo "FLAGS: " ${LDFLAGS[*]}
echo "[INF] make & install ${APPNAME}"

echo "[INF] setting build env"
go env -w GO111MODULE=off CGO_ENABLED=0

echo "[INF] building ${APPNAME}"
cd "${SCRIPTDIR}/bin/${APPNAME}"
go fmt .
sleep 0.5
go build -ldflags="${LDFLAGS[*]}"

cd "${SCRIPTDIR}"

sleep 2

APPEXE="${SCRIPTDIR}/bin/${APPNAME}/${APPNAME}"

install(){
	if [[ $USER == "root" ]]
	then
		mv "${APPEXE}" "/sbin/${APPNAME}"
	elif [[ -d $GOBIN ]]
	then
		mv "${APPEXE}" "${GOBIN}/${APPNAME}"
	fi
	echo "[INF] installed at:"
	which ${APPNAME}
}

echo "[INF] setting Go env back to:"
echo "[INF] GOARCH=${GOARCH} GOOS=${GOOS} GO111MODULE=${GOMOD}"
go env -w GOARCH=${GOARCH} GOOS=${GOOS}
go env -w GO111MODULE=${GOMOD}

echo "[INF] all done"

# executes the function $1 $2 $3 $4 $5 are the arguments
$@
