terraform {
  backend "gcs" {
    # You need to modify bucket
    bucket = "tmp-seccamp2023-b2-rung"
    prefix = "terraform/state"
  }
}
