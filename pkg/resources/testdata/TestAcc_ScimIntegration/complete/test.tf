resource "snowflake_scim_integration" "test" {
  name           = var.name
  enabled        = var.enabled
  scim_client    = var.scim_client
  sync_password  = var.sync_password
  network_policy = var.network_policy_name
  run_as_role    = var.run_as_role
  comment        = var.comment
}
