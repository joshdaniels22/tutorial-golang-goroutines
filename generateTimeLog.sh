#!/bin/bash

for i in {1,10,100,1000,10000,100000,1000000,10000000};
do 
    echo -e "for $i iterations:\n" >> time-log
    go run . $i >> time-log
    echo -e "------------------------\n" >> time-log
done
