#!/bin/bash -e

if [[ $# -eq 0 ]]; then
    echo "Expected 1 argument - day number (format: DD, where D is dec digit)"
    exit 1
fi

DAY=$1

mkdir $DAY
cd $DAY

touch sample.txt
touch input.txt

mkdir 01
cp -r ../templates/. 01

mkdir 02
cp -r ../templates/. 02
