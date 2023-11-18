#!/bin/zsh

# Creates the challenge directory and golang files for me.

if [[ -z $1 || $1 = "-h" ]];then
    echo "./init.sh <challenge URL> \n./init.sh https://adventofcode.com/2015/day/9"
    exit
fi

if [[ -z "${AOC_COOKIE}"  ]];then
    echo -n "Paste the AOC cookie:"
    read -s cookie
    export AOC_COOKIE=$cookie
fi

IFS='/' read -A URL <<< $1

year=${URL[4]}
day=${URL[6]}

dir=$year/day$day
mkdir -p $dir
cd $dir

touch test.txt
cp ../../main.go.template main.go

curl $1/input --cookie "session=$AOC_COOKIE" -o input.txt
