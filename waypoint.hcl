project = "vault-circleci-auth-plugin"

app "vault-circleci-auth-plugin" {
  build {
    use "docker" {}
  }

  deploy {
    use "docker" {
      service_port = 8200
    }
  }
}
