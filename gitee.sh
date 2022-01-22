#!/usr/bin/env bash

set -eu

cd $(dirname $0)

rm -rf .temp/gitee
mkdir -p .temp/gitee
cd .temp/gitee
git clone git@gitee.com:k3x/crawl.git .

cd ../../
cp *.go *.md .temp/gitee
cp .gitignore .temp/gitee

cd .temp/gitee
find . \( -name '*.md' -o -name '*.go' -o -name 'go.mod' -o -name 'go.sum' \) -exec sed -i '' 's@github.com/cnk3x@gitee.com/k3x@g' {} \;
rm -f go.mod go.sum
go mod init gitee.com/k3x/crawl
go mod tidy

git add .
git commit -m 'sync'
git push

cd ../../
rm -rf .temp/gitee
