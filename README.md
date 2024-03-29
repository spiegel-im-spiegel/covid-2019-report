# 日本における COVID-2019 確認発症者のレポート

**日本では2023年5月より5類に移行し以前と同じ条件で統計情報が取得できなくなったため，このリポジトリでの情報収集を終了します。WHOでは引き続き世界全体で情報提供を行っています。また日本国内でも自治体ごとに COVID-2019 を含む5類感染症の動向が公開されていますので，今後はそちらを参照してください。**

SARS-CoV-2 ウイルスによると見られる症状（COVID-2019）について整理しています。

WHO がパンデミック宣言を行った 2020-03-11 以降における国内の確認発症者数および新規死者数を7日毎に集計してグラフ化したものです。
2022年9月以降，日本国内の検査対象者の条件が変わったため，それ以前のデータとの整合性がとれなくなっています。あくまで推移を把握するための参考データとして見ていただければと思います。

![Confirmed COVID-2019 Cases in Japan](./covid-2019-new-cases-histgram-in-japan.png)

以下は国内の新規死者数のみをグラフ化したものです。同じく7日毎の集計になっています。上のグラフとは縦軸のスケールが全く異なるので注意してください。あくまで推移を見るということで...

![COVID-2019 deaths in Japan](./covid-2019-new-deaths-histgram-in-japan.png)

基となるデータは “[WHO Coronavirus Disease (COVID-19) Dashboard](https://covid19.who.int/)” から提供される CSV データを使っています。

- [`WHO-COVID-19-global-data.csv`](https://covid19.who.int/WHO-COVID-19-global-data.csv)

また，データの取得には拙作のパッケージを使用しています。

- [goark/cov19data: Importing WHO COVID-2019 Cases Global Data](https://github.com/goark/cov19data)

作成した情報はあくまで個人的な目的で作成したもので，データの正確性については保証しません（できません）し，これらを使って何かを主張するつもりもありません（私は医療関係者ではありません）ので，あらかじめご了承ください。

## 東京都のデータについて

東京都のデータは以下の Web ページから取得しています。

- [都内の最新感染動向 | 東京都 新型コロナウイルス感染症対策サイト](https://stopcovid19.metro.tokyo.lg.jp/)
  - [東京都 新型コロナウイルス陽性患者発表詳細 - データセット - 東京都オープンデータカタログサイト](https://catalog.data.metro.tokyo.lg.jp/dataset/t000010d0000000068)

東京都のデータは PCR 検査を行って陽性反応が出た人をカウントしているだけで WHO のデータとの整合性は考慮してません。あくまでも参考程度と考えてください。

**【2022-10-15 追記】** 東京都は陽性者の公表を 2022-09-26 を以って止めたようです。データが取得できなくなりましたのでグラフから外しました。

## COVID-2019 関連のリンク集

- [Coronavirus Disease (COVID-19) Situation Reports](https://www.who.int/emergencies/diseases/novel-coronavirus-2019/situation-reports) : Situation Reports 週単位の報告（PDF）になった模様
- [Flatten the curve | These guidelines are intended to help Flatten the Curve with the COVID19 outbreak, to help limit spread and reduce the load on hospitals and other healthcare.](https://www.flattenthecurve.com/)
  - [ブログ: コロナウイルス(COVID-19)へのアドバイス](https://okuranagaimo.blogspot.com/2020/03/covid-19_11.html)
- [We're in for 2 months - foobuzz](https://foobuzz.github.io/covid19/)
  - [ブログ: 私たちには2ヶ月必要です](https://okuranagaimo.blogspot.com/2020/04/2.html)
- [Mysterious Heart Damage, Not Just Lung Troubles, Befalling COVID-19 Patients | Kaiser Health News](https://khn.org/news/mysterious-heart-damage-not-just-lung-troubles-befalling-covid-19-patients/)
  - [ブログ: COVID-19患者に降りかかる不思議な心臓の損傷](https://okuranagaimo.blogspot.com/2020/04/covid-19_7.html)
- [COVID-19: The T Cell Story - Articles](https://berthub.eu/articles/posts/covid-19-t-cells/)
  - [ブログ: COVID-19: T細胞の話](https://okuranagaimo.blogspot.com/2020/06/covid-19-t.html)
- [2020年7月1日ニュース「国内初のコロナワクチンの治験を開始 創薬ベンチャーの『アンジェス』」 | SciencePortal](https://scienceportal.jst.go.jp/news/newsflash_review/newsflash/2020/07/20200701_01.html)
- [MIT Tech Review: 新型コロナとインフルの似ているところ、違うところ＝WHO報告](https://www.technologyreview.jp/nl/these-are-6-of-the-main-differences-between-flu-and-coronavirus/)
- [New Data on T Cells and the Coronavirus  |  In the Pipeline](https://blogs.sciencemag.org/pipeline/archives/2020/07/15/new-data-on-t-cells-and-the-coronavirus)
  - [ブログ: T細胞とコロナウイルスに関する新しいデータ](https://okuranagaimo.blogspot.com/2020/07/t.html)
- [COVID-19（新型コロナウイルス感染症）に関する情報とリソース - Google](https://www.google.com/intl/ja_jp/covid19/)
- [Characteristics of SARS-CoV-2 and COVID-19 | Nature Reviews Microbiology](https://www.nature.com/articles/s41579-020-00459-7?error=cookies_not_supported&code=70d81179-79f2-4810-afd8-4e9f9b6d57db)
  - [ブログ: SARS-CoV-2とCOVID-19の特徴](https://okuranagaimo.blogspot.com/2020/10/sars-cov-2covid-19.html)
- [感染症数理モデルとCOVID-19 | 日本医師会 COVID-19有識者会議](https://www.covid19-jma-medical-expert-meeting.jp/topic/3925)
- [Japan: COVID-19 Public Forecasts](https://datastudio.google.com/reporting/8224d512-a76e-4d38-91c1-935ba119eb8f/page/ncZpB?feature=opengraph)
- [COVID-19 rarely spreads through surfaces. So why are we still deep cleaning?](https://www.nature.com/articles/d41586-021-00251-4?error=cookies_not_supported&code=8208a01d-425d-4369-8169-c9d29038d2c1)
  - [ブログ: COVID-19が表面に広がることはほとんどない。では、なぜ私たちはまだディープクリーニングをしているのか?](https://okuranagaimo.blogspot.com/2021/02/covid-19.html)
- [Five reasons why COVID herd immunity is probably impossible](https://www.nature.com/articles/d41586-021-00728-2?error=cookies_not_supported&code=b270c063-5f42-45fc-bd8a-c4ed058b448c)
  - [ブログ: 新型コロナの集団免疫がおそらく不可能である5つの理由](https://okuranagaimo.blogspot.com/2021/03/5.html)
- [変異株と闘う世界で進むワクチン接種、3回目の追加接種を模索する動きも：新型コロナウイルスと世界のいま（2021年4月） | WIRED.jp](https://wired.jp/2021/05/05/covid-19-april-2021/)

## Dependency Graph

[![dependency.png](./dependency.png)](./dependency.png)
