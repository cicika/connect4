#!/bin/bash

columns=$@
echo "Starting..."
if [ ${#columns} -ge 0 ]; 
  then ./connect4 --columns "$columns" ; 
else ./connect4
fi

