#!/bin/sh

git checkout master
mv .git/temp/* ./
git add .
git commit -a -m "build binary"
latest=$(git describe --tags `git rev-list --tags --max-count=1`)
echo -e "Latest tag is "$latest
read -p "Please input version(ex: v1.0) for publising: " version
git tag ${version}
git push origin ${version}
git checkout encrypt
read -p "publish done."