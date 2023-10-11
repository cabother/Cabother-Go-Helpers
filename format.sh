#!/bin/sh

# Fix all errors from code quality
# Run with command: "sh fix.sh"
find . -name mocks -exec rm -rf {} \;

rootPath=$PWD
echo $rootPath

for filePath in $(find . -name "*.go")
do  
    path=`echo $filePath | rev | cut -d"/" -f2- | rev`
    cd $path
    mockery --all --recursive=true --with-expecter --case=underscore
    cd $rootPath
done

cont=0
for filePath in $(find . -name "*.go")
do
    # echo $filePath | rev | cut -d"/" -f2- | rev
    ((cont=$cont+1))
    file_name="${filePath##*/}"
    gofumpt -w $filePath
    echo "$cont - $file_name -> DONE!"
done
