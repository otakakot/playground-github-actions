#!/bin/bash

count=0

while read line; do

    count=$(($count + 1))

    if [[ $count -eq 5 ]] ; then
        sed -i -e '5s/^#//' $1
    fi

    if [[ $count -eq 6 ]]; then
        sed -i -e '6s/^#//' $1
    fi

done < $1
