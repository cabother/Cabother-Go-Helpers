#!/bin/sh

find . -name mocks -exec rm -rf {} \;

rootPath=$PWD
echo $rootPath

for filePath in $(find . -name "*.go")
do  
    path=`echo $filePath | rev | cut -d"/" -f2- | rev`
    cd $path
    mockery --all --recursive=true
    cd $rootPath
done
