#!/usr/bin/bash
archs=(amd64 arm64 ppc64le ppc64 s390x)
oss=(linux windows android)

# shellcheck disable=SC2068
for arch in ${archs[@]}
do
    for os in ${oss[@]}
    do
        env GOOS=${os} GOARCH=${arch} go build -o ./out/${os}/gotype_${arch}
    done
done