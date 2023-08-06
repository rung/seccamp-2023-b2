# Exercise1: CI/CDとInfrastructure as Codeを試す
## Goal
- Continuous DeploymentとInfrastructure as code(Terraform)のコンセプトを理解する

## 準備
- [Lab Setup Procedure](lab-setup.md) (受講者は実施不要)

## 現在の状態

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183403529-505390f7-ecfe-4cd4-9356-c7fcedfa5801.png" height="300"> </kbd>

## 実施手順
- Goコードを修正し、自動デプロイを確認する
- Terraformを使ってみる
- (optional) コードを確認し、何をやっているかを理解する

## Exercises Procedure
### 1. Goコードを修正し、自動デプロイを確認する
- Cloud Runにアクセスする
  - https://console.cloud.google.com/run
  - コンソールからURLを入手する
    - `https://training-helloworld-*****.a.run.app`
- seccamp2023-b2-<name>-app/main.goのHello Worldのメッセージを別のものに変える
- Cloud Runにもう一度アクセスし、自動デプロイが反映されたことを確認する

### 2. Terraformを使ってみる
- Document: https://registry.terraform.io/providers/hashicorp/google/latest/docs
- たとえば下記のコードを追加する
```terraform
resource "google_service_account" "test" {
  account_id = "service-account-id"
}
```

- 追加結果をGoogle Cloudのコンソールから確認する: https://console.cloud.google.com/iam-admin/serviceaccounts

### 3. (optional) コードを確認し、何をやっているかを理解する
- Actionsの設定などを確認し、何をやっているかを理解する
