# jikkyo

アニメ放送の「実際の開始時刻」、「A」「B」「C」パートの開始時刻を自動検出する CLI ツールです。実況コメントファイル、.ts.program.txtも出力します。


## 使い方


### フラグ一覧

```
>jikkyo help
Flags:
  -e, --episode int         Episode number (required)
  -h, --help                help for jikkyo
  -l, --log-file string     Log file path (optional)
  -o, --output-dir string   Output directory for program info file (optional)
  -t, --title string        Anime title to search for (required)
  -v, --verbose             Enable verbose logging

Error: required flag(s) "episode", "title" not set
```

### 実行例

#### 1. タイトルとエピソードを指定して検索

```bash
> jikkyo -t フリーレン -e 10
```

**結果:**
```
Multiple titles found:
1. 葬送のフリーレン (TID: 6776)
2. 葬送のフリーレン(第2期) (TID: 7629)
```

#### 2. 複数マッチする場合は選択
```bash
Select (1-2): 1
```

**結果:**

```
(title) 葬送のフリーレン
(episode) 10
(subtitle) 強い魔法使い
(start) 2023-11-10 23:30:00
(real_start_time) 2023-11-10 23:30:03
(A) 2023-11-10 23:34:20
(B) 2023-11-10 23:45:30
(program_filename) 202311102330000102-葬送のフリーレン 第10話『強い魔法使い』[字].ts.program.txt
(program_content) 2023/11/10(金) 23:30～00:00
日テレ１
葬送のフリーレン 第10話『強い魔法使い』[字]
```


## 外部 API クレジット

- [Syoboi Calendar](http://cal.syoboi.jp/) - テレビ放送予定データベース
- [Jikkyo API](https://jikkyo.tsukumijima.net/) - 放送コメント API

