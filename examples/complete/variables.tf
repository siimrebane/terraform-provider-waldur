# Waldur Connection
variable "waldur_endpoint" {
  description = "Waldur API endpoint URL"
  type        = string
}

variable "waldur_token" {
  description = "Waldur API token"
  type        = string
  sensitive   = true
}

# Project
variable "project_name" {
  description = "Name of your Waldur project"
  type        = string
}

# Environment
variable "environment" {
  description = "Environment name (used as prefix)"
  type        = string
  default     = "dev"
}

# UUIDs (find these via Waldur API)
variable "offering_uuid" {
  description = "Marketplace offering UUID for OpenStack.Instance"
  type        = string
}

variable "volume_type_uuid" {
  description = "Volume type UUID"
  type        = string
}

variable "subnet_uuid" {
  description = "Subnet UUID for VM network"
  type        = string
}

variable "ssh_key_uuid" {
  description = "SSH public key UUID registered in Waldur"
  type        = string
}
