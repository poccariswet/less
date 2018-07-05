# less

this repo for acl's study 


## Installation

``` sh
$ go get github.com/rthornton128/goncurses
```

## Each file

- buffer.go: ファイル操作
- view.go: レンダリング処理、スクロール系の処理
- main.go: initの処理、rawmodeでの1文字待ち

## Usage

```sh
$ go build
$ less <file name>
```

 - <kbd>q</kbd> Quits
 - <kbd>↑</kbd>, <kbd>k</kbd> scroll up
 - <kbd>↓</kbd>,<kbd>j</kbd>, <kbd>Enter</kbd> scroll down
