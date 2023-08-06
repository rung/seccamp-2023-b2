# Exercise2: CI/CDへの攻撃
## Goal
- CI/CDへの攻撃を通してアタックサーフェスを理解する
  - 開発者の開発端末が侵害されたことや、ソフトウェア・ツールの依存先が侵害されたことを想定する

## 事前準備(任意)
- 配られた侵入先の端末Cookieを使い、GitHubにアクセスする
  1. Chromeでgithub.comを開く
  1. Cookie Editorに、配られたCookieファイルをImportsする
  1. リロードし、Cookieの読み込みが行われ侵入したアカウントのセッションを奪えたことを確認する

## 演習
- ソースコードをレビューをバイパスして上書きする
- Non-protected branchからクレデンシャルを盗む
- Actionsの依存を侵害する（Supply-Chain attacks）

## Procedure
### 1. ソースコードをレビューをバイパスして上書きする
- main branch上のソースコード(seccamp2023-b2-app)を変更する

### 2. Non-protected branchからクレデンシャルを盗む
- `echo ${{ secrets.GCP_SA_KEY }} | base64` on CI  

### 3. Actionsの依存を侵害する（Supply-Chain attacks）
- （今回は一緒のレポジトリに配置しているが、seccamp2023-b2-app上のCDが外部の"supplychain" actionsに依存している想定）
  - seccamp2023-b2-app/supplychain/action.yml を編集する

## 追加演習
- IACのCI(read)で使われているトークンを盗み、Google Storage上に保管されたtfstateファイルを確認する
- AppのDeploymentで使われているトークンを盗み、次にオーナーロールに権限昇格する
  - Reference: https://rhinosecuritylabs.com/gcp/privilege-escalation-google-cloud-platform-part-1/
