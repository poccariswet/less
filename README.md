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

## using ncurses

- WINDOW *iniscr(void); <br>
端末の初期化．
環境変数TERM，標準入出力を利用し，端末を初期化します．

- int endwin(void); <br>
スクリーンをリセットし，カーソルを左下に移動します．
スクリーン関連のメモリなどは解放しないため，必要に応じてdelscreenを呼び出す必要があります．

- int scrl(int n); <br>
nが正の数の場合，上方向にn行スクロールします．
負の数の場合，下方向にn行スクロールします．
この関数を使う場合，scrollokでスクロールが有効になっている必要があります．

- void getmaxyx(WINDOW *win, int y, int x); <br>
指定されたウィンドウのサイズを取得します．
これは，関数ではなくマクロで実装されています．

- int raw(void); <br>
raw modeにする。

- int getch(void); [w,mv,mvw] <br>
1文字読み込みます．
keypadがTRUEに設定されている場合は，キーコードが帰ってくる可能性があります．

- int refresh(void); [w] <br>
端末をウィンドウデータに基づき再描画します．
この関数は，wnoutrefreshとdoupdateを連続で呼び出すのとほぼ等価です．

- int printw(char *fmt [, arg] ...); [w,mv,mvw] <br>
printfに従い書式化した文字列を，addstrに渡した場合と同等です．

- int keypad(WINDOW *win, bool bf); <br>
bfにTRUEを指定した場合，wgetch等において，ファンクションキーを押したときにファンクションキーのコードを返すようになります．
FALSEの場合，エスケープシーケンスのまま取り出されます．

- nt echo(void); <br>
キーをタイプしたときに，スクリーンにエコーするモードにします


## Usage

```sh
$ go build
$ less <file name>
```

 - <kbd>q</kbd> Quits
 - <kbd>↑</kbd>, <kbd>k</kbd> scroll up
 - <kbd>↓</kbd>,<kbd>j</kbd>, <kbd>Enter</kbd> scroll down


## Reference

- [ncurses](http://www.kis-lab.com/serikashiki/man/ncurses.html)
- [goncurses](https://github.com/rthornton128/goncurses)