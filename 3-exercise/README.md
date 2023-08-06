# Exercise3: CI/CDパイプラインのセキュア化
## Goal
- いくつかのCI/CDパイプラインのセキュア化方法を試して、効果と限界を理解する

## Exercises
- Branch Protectionの設定
  - 注意: GitHub Freeではできない
- OIDC認証を通した長期鍵を利用しないGitHub ActionsとGoogle Cloudの連携

## Exercises procedure
### 1. Branch Protectionの設定
- レポジトリ: `seccamp-2023-b2-<name>-iac`

<kbd> <img src="https://github.com/rung/seccamp-2023-b2/assets/1150301/ccbddaf1-99c0-488d-8416-aa1bf1dcc7ed" height="400"> </kbd>


### 2. OIDC認証を通した長期鍵を利用しないGitHub ActionsとGoogle Cloudの連携
- レポジトリ: `seccamp-2023-b2-<name>-iac`
- `google_actions_oidc.tf_`を`google_actions_oidc.tf`　にリネームする (`seccamp-2023-b2-<name>-iac/terraform/training-project/`)
  - 下記の部分を、自分のレポジトリに変更する
      ```
        # You need to modify this value
        locals {
          app_repo_name = "<github org or name>/seccamp-2023-b2-app"
          iac_repo_name = "<github org or name>/seccamp-2023-b2-iac"
        }
      ```
  - このTerraformでOIDC認証設定が有効化される


<kbd> <img src="https://user-images.githubusercontent.com/1150301/183426987-2ba5d9ce-2d9d-4e33-882e-e0e732f3568c.png" height="200"> </kbd>

- Iac'sのActions Workflowsを編集する
  - `seccamp-2023-b2-iac/.github/workflows/apply.yaml`, `seccamp-2023-b2-iac/.github/workflows/plan.yaml`
    - アンコメント：`id-token: 'write'`
    - コメント化：`credentials_json`
    - アンコメント：`workload_identity_provider: 'projects/<Project Number>/locations/global/workloadIdentityPools/training-pool/providers/training-provider'`
    - アンコメント `service_account: 'iac-actions-cd@<Project ID>.iam.gserviceaccount.com'`
- Appに関してもOIDC認証をActionsの編集で実施可能

### オプショナル: SBOM
- 今回のCI/CD環境では、アプリケーションのビルドの際にSBOMを生成しています。SBOMがどのような"対策"だと言えるのか、改めて考えてみてください

## オプショナル2 (実施不要)
- GitHub Actions上からリバースシェルを貼り、OIDC認証において用いられる認証情報を窃取する
