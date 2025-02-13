terraform {
  required_providers {
    k6 = {
      version = "0.0.1"
      source = "k6.io/loadimpact/k6"
    }
  }
}

provider "k6" {
  token = "936c2433aca0f641425980e3c20b4871fb3a4291a6b3c34805dd08fc0e8afc4f"
}

data "k6_organizations" "all" {}
