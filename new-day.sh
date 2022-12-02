#!/bin/bash -e

if [[ $# -eq 0 ]]; then
    echo "Expected 1 argument - day number (format: DD, where D is dec digit)"
    exit 1
fi

DAY=$1

mkdir $DAY
cd $DAY

mkdir 01
cd 01
go mod init aoc-2022/$DAY/01
cd -
cp -r ../templates/. 01

mkdir 02
cd 02
go mod init aoc-2022/$DAY/02
cd -
cp -r ../templates/. 02
