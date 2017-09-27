#! /sbin/bash


for i in `seq 1 24`; do
	echo $i hola
	sed 's/^/###/' $i.md > c-$i.txt

done

