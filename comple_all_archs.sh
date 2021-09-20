#!/usr/bin/bash
archs=(amd64 arm64 ppc64le ppc64 s390x)
oss=(linux windows android)

# shellcheck disable=SC2068
for os in ${oss[@]}
do
    echo " Attempting to compile for "${os}" on all architectures..."
    for arch in ${archs[@]}
    do
        env GOOS=${os} GOARCH=${arch} go build -o ./out/${os}/gotype_${arch} 2>/dev/null
    done
done
echo "Compilation finished!"