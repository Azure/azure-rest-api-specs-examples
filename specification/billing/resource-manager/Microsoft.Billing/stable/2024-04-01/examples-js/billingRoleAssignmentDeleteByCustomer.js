const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a role assignment on a customer. The operation is supported for billing accounts with agreement type Microsoft Partner Agreement.
 *
 * @summary Deletes a role assignment on a customer. The operation is supported for billing accounts with agreement type Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentDeleteByCustomer.json
 */
async function billingRoleAssignmentDeleteByCustomer() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30";
  const billingProfileName = "BKM6-54VH-BG7-PGB";
  const customerName = "703ab484-dda2-4402-827b-a74513b61e2d";
  const billingRoleAssignmentName =
    "30000000-aaaa-bbbb-cccc-100000000000_00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingRoleAssignments.deleteByCustomer(
    billingAccountName,
    billingProfileName,
    customerName,
    billingRoleAssignmentName,
  );
  console.log(result);
}
