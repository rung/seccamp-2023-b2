# What is this

このレポジトリはセキュリティ・キャンプ全国大会2023のトレーニング「B2: 開発プロセスを攻撃者の視点で捉える(後半)」の演習の一部です

## 🎫 Requirements
- Google Cloud Account
- GitHub Account (Pro Version)

## 💻 Terminal
- Mac: Terminal
- Windows: WSL

- CloudShellでも実行可能です。
[![Open in Cloud Shell](http://gstatic.com/cloudssh/images/open-btn.png)](https://console.cloud.google.com/cloudshell/open?git_repo=https://github.com/rung/seccamp-2023-b2)

## ✍️ Exercises
### [Preparation: Setup Google Cloud and GitHub](0-preparation/README.md)
- ゴール: Google CloudおよびGitHubのセットアップ

### [Exercise1: Make and Try continuous deployment and Infrastructure as code](./3-exercise3/README.md)
- Goal: Understand the concept of Continuous Deployment and Infrastructure as code(Terraform)
- Exercise: Modify Go code and see automatic deployment, Add configuration via Terraform

### [Exercise2: Attack against CI/CD](./4-exercise4/README.md)
- Goal: Attack on CI/CD pipelines and understanding the attack surface
- Exercise: Overwrite source code without any review, Steal secrets from a non-protected branch, Try Supply-Chain attacks via Actions the repository uses

### [Exercise3: Secure your CI/CD pipeline](./5-exercise5/README.md)
- Goal: Try to secure CI/CD pipeline from attacks
- Exercise: Configure Branch Protection, Configure OIDC, then try keyless between GitHub actions and Google Cloud
