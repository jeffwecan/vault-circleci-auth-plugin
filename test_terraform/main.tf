
variable "circleci_auth_ttl" {
  type        = string
  description = "Time duration used as the default ttl when no ttl is specified with the login request. This value is larger than max_ttl, it will be capped at max_ttl."
  default     = "5m"
}

variable "circleci_auth_attempt_cache_expiry" {
  type        = string
  description = "Time duration that sets the largest ttl that a new token can be assigned in this plugin."
  default     = "18000s"
}

variable "circleci_auth_max_ttl" {
  type        = string
  description = <<-description
      Time duration that sets the largest ttl that a new token can be assigned in this plugin.

      The plugin caches login attempts for and an approximate duration of attempts_cache_time.
      The longer this duration, the greater the memory requirement will be.
      However, in order to prevent replay attacks, the plugin must cache the attempt for a duration that exceeds
      the maximum CircleCI build duration possible.
      description
  default     = "60m"
}

variable "circleci_auth_base_url" {
  type        = string
  description = "Alternate base URL where CircleCI API calls are sent. This parameter must end up with a slash."
  default     = "https://circleci.com/api/v1.1/"
}

variable "circleci_auth_vcs_type" {
  type        = string
  description = "Value indicating the Version Control System type of the project builds being looked up in CircleCI's API. Valid values are github and bitbucket."
  default     = "github"
}

variable "circleci_auth_owner" {
  type        = string
  description = "The username or organization that owns the project in the VCS (as reflected on CircleCI's side of things)."
  default     = "jeffwecan"
}

variable "circleci_auth_token" {
  type        = string
  description = "CircleCI personal API token that allows the associated Vault authentication plugin to make API calls to CircleCI."
  # sensitive   = true
}

terraform {
  required_version = ">= 0.13"
  required_providers {
    vault = {
      source  = "hashicorp/vault"
      version = "~> 2.15.0"
    }
  }
}

resource "vault_generic_endpoint" "circleci_config" {
  path = "auth/circleci/config"

  data_json = jsonencode({
    circleci_token       = var.circleci_auth_token
    base_url             = var.circleci_auth_base_url
    vcs_type             = var.circleci_auth_vcs_type
    owner                = var.circleci_auth_owner
    ttl                  = var.circleci_auth_ttl
    max_tll              = var.circleci_auth_max_ttl
    attempt_cache_expiry = var.circleci_auth_attempt_cache_expiry
  })
}
