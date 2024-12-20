# go-lambda-common-lib

共通ライブラリの構築により複数のLambdaに共通する処理をライブラリ経由で利用できるようにして、
新規/既存Lambdaの開発工数、Lambdaのバージョンアップの開発工数、手間やミスの発生リスクおよび作業における心理的負担の低減し、作業効率化を実現します。

共通化の他の方法として、LambdaレイヤーというLambda関数の依存関係を効率的に管理できる機能がありますが、
GoやRustでLambda関数を作る場合はレイヤーを使わない方が良いと公式が提唱しているため使いません。
https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/chapter-layers.html
