## ■MEMO

### add command
<pre>
cobra add parent
cobra add child -p parentCmd
</pre>

## gox

https://github.com/mitchellh/gox

gox -os="linux darwin windows" -arch="amd64"

## ghr

https://github.com/tcnksm/ghr

git config --global github.token "....."

export GITHUB_API=http://github.company.com/api/v3/

ghr v0.1.0 dist/

## environment
$ go version

go version go1.9.4 linux/amd64

## exec(go run)
$ dep ensure

$ go run main.go --config .repertorium.yaml get