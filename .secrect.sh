#!/bin/sh

file=$(cd `dirname $0`; pwd)
file="${file##*/}"
file=~/.ssh/${file,,}.secrect
echo git-crypt: $file

if [ -f $file ]; then
	git-crypt unlock $file
else
	read -p "warning: git-crypt was invalid!" ret
	if [[ $ret == "init" ]]
	then
		git-crypt init
		git-crypt export-key $file
		echo "git-crypt has been created."
	fi
fi

read -p "git-crypt has been inited."