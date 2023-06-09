#!/bin/sh

cp -f go.mod .git/temp/
cp -f go.sum .git/temp/
cp -f LICENSE.md .git/temp/
cp -f README.md .git/temp/
git checkout master
cp -rf .git/temp/* ./
rm -rf .git/temp
mkdir .git/temp
git add .
git commit -a -m "build binary"
latest=$(git describe --tags `git rev-list --tags --max-count=1`)
echo -e "Latest tag is "$latest
read -p "Please input version(ex: v1.0.0) for publising: " version
git tag ${version}
git push --tags
git checkout encrypt
read -p "publish done."