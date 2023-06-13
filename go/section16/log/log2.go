package main

import (
    "log"
    "os"
)

func main() {
    // ログの出力先
    log.SetOutput(os.Stdout)

    // フォーマットを指定
    // 標準フォーマット
    log.SetFlags(log.LstdFlags)
    log.Println("A")

    // マイクロ秒を追加
    log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
    log.Println("B")

    //　時刻とファイルの行番号（短縮形）
    log.SetFlags(log.Ltime | log.Lshortfile)
    log.Println("C")

    log.SetFlags(log.LstdFlags)
    // ログのプリフィックスを設定
    log.SetPrefix("[LOG]")
    log.Println("E")


}