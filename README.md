# Golang-REST-API
Go言語と標準ライブラリのみを用いてREST APIのサンプルを作成

## なぜ作ったか？
Rubyと異なり、Go言語にはフレームワークやORMに複数の選択肢が存在する（Rubyの場合は事実上一択）

そのため各フレームワークやORMが何をやってくれるのか？・何を便利にしてくれるのか？をしっかり理解しておかないと適切な技術選定ができない。

そこでまずはフレームワークやORMを使わず標準ライブラリだけでCRUD操作を行うREST APIのサンプルを作成した。

このリポジトリをベースに様々な技術を導入していくことで、フレームワークやORMが何をどう抽象化してくれてどのような特徴があるのかを理解していきたい。

## 実装上の工夫
クリーンアーキテクチャの採用や依存性の注入の活用を心がけた。

データベースへのアクセスはinfrastructure配下に隔離しているため、ORMを導入する場合の影響範囲を最小限に抑えられるからである。

GinやEchoなどフレームワークを導入する場合はinfrastructure以外への影響も大きくなるが、それでもdomainへの影響は与えない（今回はシンプルなタスク管理アプリの想定なので、そもそも大したドメインロジックはないが）

## 直したいところ
Goの勉強をしながら作っていたため、書き方が不統一だったり引っかかる点は多々あると思う。

必ずしも全ての箇所でベストプラクティスを追求できているわけではなく「調べたものを真似してとりあえず動いた」にとどまっている点もあるはず。

抜け・漏れもありそう。

いきなり完璧を求めすぎず徐々に改善していければと思う。

## APIの動作確認

### 1. Create(POST)

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"title":"New Task","description":"This is a new task"}' \
  http://localhost:8080/tasks
```

### 2. Read(GET)

#### (a)一覧取得

```bash
curl http://localhost:8080/tasks
```

#### (b)個別取得

```bash
curl http://localhost:8080/tasks/1
```

### 3. Update(PUTまたはPATCH)

```bash
curl -X PUT \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Title","description":"Updated Description"}' \
  http://localhost:8080/tasks/1
```

### 4. Delete(DELETE)

```bash
curl -X DELETE http://localhost:8080/tasks/1
```