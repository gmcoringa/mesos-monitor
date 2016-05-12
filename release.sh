#!/bin/bash
set -e

increment () {
  declare -a part=( ${1//\./ } )
  declare    new
  declare -i carry=1

  for (( CNTR=${#part[@]}-1; CNTR>=0; CNTR-=1 )); do
    len=${#part[CNTR]}
    new=$((part[CNTR]+carry))
    [ ${#new} -gt $len ] && carry=1 || carry=0
    [ $CNTR -gt 0 ] && part[CNTR]=${new: -len} || part[CNTR]=${new}
  done
  new="${part[*]}"
  echo -e "${new// /.}"
} 

LAST_VERSION="${LAST_VERSION:-$(git describe --abbrev=0 --tags)}"
NEW_VERSION="`increment $LAST_VERSION`"

git fpm
git tag -a $NEW_VERSION -m "Version $NEW_VERSION release"
git push origin $NEW_VERSION

echo "Release complete"
 