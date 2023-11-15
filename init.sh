#!/bin/zsh

# Creates the challenge directory and golang files for me.

# if [[ $# -eq 0  ||  $1 -eq "-h" ]];then
#     echo "./init.sh <challenge URL> \n./init.sh https://adventofcode.com/2015/day/9"
# fi
#
#

IFS='/' read -A URL <<< $1

year=${URL[4]}
day=${URL[6]}

dir=$year/day$day
mkdir -p $dir

cd $dir
go mod init day$day
touch main.go sol1.go sol2.go test.txt input.txt

cat << EOF > main.go
package main

func main(){
    sol1()
    //sol2()
}
EOF

cat << EOF > sol1.go
package main

func sol1(){
    
}
EOF
cat << EOF > sol2.go
package main

func sol2(){

}
EOF
