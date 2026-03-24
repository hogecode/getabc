

受け取った引数を使ってCLIにA、B、Cの開始時間を表示する処理は以前書いた。
今回は、それに加えて以下のようなファイルを作成してほしい。

ファイル名は
202603192356000102-エリスの聖杯　第１１話『運命にあらがう者たち』[字].ts.program.txt
{year}{month}{day}{start}000102-{title}{ }第{episode_number}{subtitle}.ts.program.txt
のような形式である。


ファイルの内容は以下のようになる。
```
2026/03/19(木) 23:56～00:26
ＴＢＳ１
エリスの聖杯　第１１話『運命にあらがう者たち』[字]

```

```
{year}/{month}/{day}({week}) {start}～{end}
{channel}
{title}　第{episode}話『{subitle}』[字]

```

既に存在しているコードだが、このようにProgLookupコマンドを呼び出す処理が存在する。
http://cal.syoboi.jp/db?Command=ProgLookup&TID=1853&JOIN=SubTitles

レスポンスは以下のようになる。これもすでに型がコードに存在している。
```xml
<ProgLookupResponse>
<ProgItems>
<ProgItem id="157950">
<LastUpdate>2010-01-06 10:48:07</LastUpdate>
<PID>157950</PID>
<TID>1853</TID>
<StTime>2010-02-07 08:30:00</StTime>
<StOffset>0</StOffset>
<EdTime>2010-02-07 09:00:00</EdTime>
<Count>1</Count>
<SubTitle/>
<ProgComment/>
<Flag>2</Flag>
<Deleted>0</Deleted>
<Warn>1</Warn>
<ChID>67</ChID>
<Revision>0</Revision>
<STSubTitle>私、変わります！変わってみせます!!</STSubTitle>
</ProgItem>
```
Countは話数で、それに対応するサブタイトルはSTSubTitleで取得できる。