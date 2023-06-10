#!/bin/sh

cp -rf res/* temp/
cp -f go.mod temp/
cp -f go.sum temp/
cp -f LICENSE.md temp/
cp -f README.md temp/
git checkout master
cp -rf temp/* ./
rm -rf temp
mkdir temp
git add .
git commit -a -m "build binary"
latest=$(git describe --tags `git rev-list --tags --max-count=1`)
echo -e "Latest tag is "$latest
read -p "Please input version(ex: v1.0.0) for publising: " version
git tag ${version}
git push --tags
git checkout encrypt
read -p "publish done."