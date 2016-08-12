#!/bin/bash
gz_files=$(find src/gologrotate -name "*.gz" | wc -l)
if [[ $gz_files -eq 3 ]]
then
	echo "gologrotate created the files correctly"
else
	echo "Files not found, test FAILED"
	exit 1
fi