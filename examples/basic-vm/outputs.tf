output "vm_id" {
  description = "UUID of the created VM"
  value       = waldur_openstack_instance.vm.id
}

output "vm_internal_ips" {
  description = "Internal IP addresses"
  value       = waldur_openstack_instance.vm.internal_ips
}

output "vm_state" {
  description = "VM state"
  value       = waldur_openstack_instance.vm.state
}
