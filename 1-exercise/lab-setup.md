# Lab Setup Procedure
## 受講者の環境はセットアップ済のため、Lab Setupは実施不要です
### Prepare GitHub repositories and automation flows
#### Prepare GitHub repositories
- CloudShell user needs to set Git Configuration
```
$ git config --global user.email "<Your email address>"
$ git config --global user.name "<Your name>"
```

- For application
  - Make an empty repository. 
    - Name: `seccamp2023-b2-app`
    - Visibility: `Private`

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183316472-fa7ad0f4-a7f5-4907-8232-1d190e8cb889.png" height="300"> </kbd>


```bash
# Go to "seccamp-2023-b2"
$ cd template/seccamp2023-b2-app/
$ git init
$ git add -A
$ git status
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   .github/workflows/deploy.yaml
        ...
$ git commit -m "first commit"
[master (root-commit) 215a574] first commit
 7 files changed, 159 insertions(+)
 create mode 100644 .github/workflows/deploy.yaml
 ...
$ git remote add origin https://github.com/<Your org or user>/seccamp2023-b2-app.git
$ git branch -M main
$ git push -u origin main
$
```
  - Access the repository

- For Infrastructure as Code
  - Make an empty repository.
    - Name: `seccamp2023-b2-iac`
    - Visibility: `Private`

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183316487-f54ea2f8-62e3-405c-a5da-8ce7419b6339.png" height="300"> </kbd>


```
# Go to Go to "seccamp-2023-b2"
$ cd template/seccamp2023-b2-iac/
$ git init
$ git add -A
$ git status
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   .github/workflows/apply.yaml
        ...

$ git commit -m "first commit"
[master (root-commit) 5ba5f2b] first commit
 10 files changed, 234 insertions(+)
 create mode 100644 .github/workflows/apply.yaml
 ...
$ git remote add origin https://github.com/<Your org or user>/seccamp2023-b2-iac.git
$ git branch -M main
$ git push -u origin main
$
```
  - Access the repository

#### Configure google cloud and GitHub for automation
- Google Cloud Configuration
  - Create a new service account, grant owner to the SA, then download a key of the SA
```bash
$ PROJECT_ID=<PROJECT_ID>
$ gcloud config set project $PROJECT_ID
$ gcloud iam service-accounts create actions-secret
$ gcloud projects add-iam-policy-binding ${PROJECT_ID} --member="serviceAccount:actions-secret@${PROJECT_ID}.iam.gserviceaccount.com" --role="roles/owner"
$ gcloud iam service-accounts keys create google_static_key.json --iam-account=actions-secret@${PROJECT_ID}.iam.gserviceaccount.com
$ cat google_static_key.json
```

  - Save google_static_key.json to the GitHub
    - `https://github.com/<Your org or name>/seccamp2023-b2-app/settings/secrets/actions`
    - KEY_NAME: `GCP_SA_KEY`
    - ![image](https://user-images.githubusercontent.com/1150301/183318694-7b57f92d-f95f-448d-8ec4-534f9a23de82.png)
  - Do same thing to `seccamp-2023-b2-iac` too

#### Configure google cloud for Terraform (IaC)
- Create bucket for terraform state file
  - Bucket Name: `tmp-seccamp2023-b2-<GitHub Name>`
```
$ gsutil mb -c Standard -l US --pap enforced -b on gs://tmp-seccamp2023-b2-<Your GitHub Name>
```

#### Modify GitHub repositories
##### seccamp2023-b2-iac/terraform/training-project/
- backend.tf
  - Modify `bucket` to correct name you created
  - `tmp-seccamp2023-b2-<Your GitHub Name>`
- variable.tf
  - Change the project name to correct one
```
variable "project" {
  default = "kinetic-dryad-354617"
}
 ```
- Actions should succeed

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183320761-be628735-f8c2-49f6-9883-a7a0896f488f.png" height="100"> </kbd>

* When you get error like `API has not been used in project `, please re-run the job



##### seccamp2023-b2-app/.github/workflows/deploy.yaml
- Change this line. (Replace `<Project ID>` to Your Project Name)
```
    env:
      # You need to modify these values
      PROJECT_ID: <Project ID>
      RUN_SA_ID: run-helloworld@<Project ID>.iam.gserviceaccount.com
```

** Now, You can access cloud run **

- Service URL is written in actions result.

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183321839-861268d6-3ba4-42c7-a3d4-fb1fc3118bc2.png" height="300"> </kbd>
  
<kbd> <img src="https://user-images.githubusercontent.com/1150301/183321810-3b6815a6-65ee-4858-9955-0d04d2d0c093.png" height="100"> </kbd>
