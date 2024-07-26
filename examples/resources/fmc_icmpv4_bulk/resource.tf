resource "fmc_icmpv4_bulk" "example" {
  bulk = [
    {
      name        = "icmpv4_obj"
      icmp_type   = "3"
      code        = 0
      description = "Test Description"
      overridable = true
      parent_id   = "icmpv4_obj_uuid"
      parent_type = "ICMPV4Object"
      target_id   = "domain_uuid"
      target_type = "Domain"
      target_name = "Global \\ EUROPE"
    }
  ]
}
