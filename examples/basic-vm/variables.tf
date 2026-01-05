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

# VM Configuration
variable "vm_name" {
  description = "Name for the VM"
  type        = string
  default     = "my-terraform-vm"
}

variable "flavor_name" {
  description = "VM flavor (size)"
  type        = string
  default     = "m1.small"
}

variable "image_name" {
  description = "OS image name"
  type        = string
  default     = "Ubuntu 22.04"
}

variable "volume_size" {
  description = "System volume size in MiB"
  type        = number
  default     = 20480  # 20 GB
}

# UUIDs (find these via Waldur API or UI)
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
