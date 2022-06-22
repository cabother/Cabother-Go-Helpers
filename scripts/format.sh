#!/bin/sh

# Fix all errors from code quality
# Run with command: "sh fix.sh"
rootPath=$PWD
echo $rootPath

cont=0
for filePath in $(find . -name "*.go")
do
    # echo $filePath | rev | cut -d"/" -f2- | rev
    ((cont=$cont+1))
    file_name="${filePath##*/}"
    gofumpt -w $filePath
    echo "$cont - $file_name -> DONE!"
done
